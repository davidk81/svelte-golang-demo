package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	addr     = flag.String("addr", ":8000", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	log.Printf("%s %s\n", ctx.Request.Header.Method(), ctx.Path())
	switch string(ctx.Path()) {
	case "/session":
		Session(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}
