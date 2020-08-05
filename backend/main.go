package main

// backend http server for patient service demo
// usage: ./backend --help

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/davidk81/svelte-golang-demo/backend/patient"
	"github.com/davidk81/svelte-golang-demo/backend/patientdb"
	"github.com/davidk81/svelte-golang-demo/backend/session"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
)

// config flags
var (
	dbConn   = flag.String("db", "host=localhost dbname=patientdb user=docker password=docker sslmode=disable", "db connection string")
	addr     = flag.String("addr", "localhost:8000", "tcp listen address & port")
	compress = flag.Bool("compress", false, "response compression [true/false]")
)

func main() {
	flag.Parse()

	// init db connection
	patientdb.Init(*dbConn)

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	// start server
	go func() {
		log.Println("server starting on port", *addr)
		if err := fasthttp.ListenAndServe(*addr, h); err != nil {
			log.Fatalf("error during ListenAndServe: %s", err)
		}
	}()

	// wait for cleanup signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-sigs
		done <- true
	}()
	<-done

	// cleanup
	patientdb.Close()
}

func route(ctx *fasthttp.RequestCtx) error {
	// routes that dont need session
	// TODO: enable session checking
	apiPrefix := "/api/v1/"
	switch string(ctx.Path()) {
	case apiPrefix + "session":
		return session.HandleSession(ctx)
	case apiPrefix + "register":
		return session.HandleRegister(ctx)
	case "/healthz":
		return handleHealth(ctx)
	}

	// check users' session token
	_, err := session.ValidateSession(ctx, "nurse")
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
		return nil
	}

	// routes that need session
	switch string(ctx.Path()) {
	case apiPrefix + "patients":
		return patient.HandlePatientList(ctx)
	case apiPrefix + "patient":
		return patient.HandlePatient(ctx)
	case apiPrefix + "patient/note":
		return patient.HandlePatientNote(ctx)
	case apiPrefix + "patient/notes":
		return patient.HandlePatientNoteList(ctx)
	default:
		ctx.NotFound()
		return nil
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	log.Printf("%s %s\n", ctx.Request.Header.Method(), ctx.URI().RequestURI())

	// enable cors for development
	ctx.Response.Header.Set("access-control-allow-credentials", "true")
	ctx.Response.Header.Set("access-control-allow-origin", string(ctx.Request.Header.Peek("Origin")))
	ctx.Response.Header.Set("access-control-expose-headers", "WWW-Authenticate,Server-Authorization")
	ctx.Response.Header.Set("cache-control", "no-cache")
	ctx.Response.Header.Set("Connection", "keep-alive")

	if ctx.IsOptions() {
		handleMethodOptions(ctx)
		return
	}

	// handle request routes
	err := route(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.SetBody([]byte(err.Error()))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
}

func handleMethodOptions(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("access-control-allow-headers", "Accept,Authorization,Content-Type,If-None-Match")
	ctx.Response.Header.Set("access-control-allow-methods", string(ctx.Request.Header.Peek("Access-Control-Request-Method")))
	ctx.Response.Header.Set("access-control-max-age", "86400")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func handleHealth(ctx *fasthttp.RequestCtx) error {
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
