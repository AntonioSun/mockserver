// mock server to replace the Java/NPM counterpart mockserver
package main

//go:generate sh mockserver_cli.sh

import (
	"fmt"
	"github.com/valyala/fasthttp/prefork"
	"log"

	"github.com/caarlos0/env"
	"github.com/valyala/fasthttp"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions
var (
	progname = "mockserver"
	version  = "1.0.3"
	date     = "2022-01-30"

	e envConfig
)

func main() {
	// == Config handling
	err := env.Parse(&e)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("] %+v\n", e)

	err = ParseMockJson(e.File)
	if err != nil {
		log.Fatal(err)
	}
	err = VerifyMockJson()
	if err != nil {
		log.Fatal(err)
	}
	printPaths()

	h := requestHandler
	if e.Compress {
		h = fasthttp.CompressHandler(h)
	}

	server := &fasthttp.Server{
		Handler: h,
	}

	if e.Prefork {
		// Wraps the server with prefork
		fmt.Println("Starting preforked server on", e.Addr)
		preforkServer := prefork.New(server)
		if err = preforkServer.ListenAndServe(e.Addr); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Starting server on", e.Addr)
		if err = server.ListenAndServe(e.Addr); err != nil {
			panic(err)
		}
	}
}
