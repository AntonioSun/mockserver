// mock server to replace the Java/NPM counterpart mockserver
package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/valyala/fasthttp"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions
var (
	progname = "mockserver"
	version  = "1.0.1"
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

	fmt.Println("Starting server on", e.Addr)

	err = fasthttp.ListenAndServe(e.Addr, h)
	if err != nil {
		panic(err)
	}
}
