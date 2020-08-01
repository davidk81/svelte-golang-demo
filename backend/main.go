package main

//go:generate sqlboiler --wipe psql

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

var (
	addr     = flag.String("addr", ":8000", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	// init db
	patientdb.Init()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	go func() {
		log.Println("server starting on port", *addr)
		if err := fasthttp.ListenAndServe(*addr, h); err != nil {
			log.Fatalf("Error in ListenAndServe: %s", err)
		}
	}()

	// cleanup
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-sigs
		done <- true
	}()
	<-done
	patientdb.Close()
}

func route(ctx *fasthttp.RequestCtx) error {
	// route that dont need session
	switch string(ctx.Path()) {
	case "/api/v1/session":
		return session.HandleSession(ctx)
	case "/api/v1/patients":
		return patient.HandlePatientList(ctx)
	case "/api/v1/patient":
		return patient.HandlePatient(ctx)
	case "/api/v1/patient/note":
		return patient.HandlePatientNote(ctx)
	case "/api/v1/patient/notes":
		return patient.HandlePatientNoteList(ctx)
	}

	// routes that need session
	if !session.VerifySession(ctx, "nurse") {
		ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
	}

	switch string(ctx.Path()) {
	default:
		ctx.NotFound()
		return nil
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	log.Printf("%s %s\n", ctx.Request.Header.Method(), ctx.Path())

	ctx.Response.Header.Set("access-control-allow-credentials", "true")
	ctx.Response.Header.Set("access-control-allow-origin", string(ctx.Request.Header.Peek("Origin")))
	ctx.Response.Header.Set("access-control-expose-headers", "WWW-Authenticate,Server-Authorization")
	ctx.Response.Header.Set("cache-control", "no-cache")
	ctx.Response.Header.Set("Connection", "keep-alive")

	switch string(ctx.Request.Header.Method()) {
	case "OPTIONS":
		handleMethodOptions(ctx)
		return
	}

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
