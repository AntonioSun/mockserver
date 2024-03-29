package main

import (
	"fmt"
	"log"
	"net"
	"reflect"
	"regexp"

	"github.com/valyala/fasthttp"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type ipResponse struct {
	Origin string `json:"origin"`
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = "application/json"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	for _, k := range respMap {
		if path == k.HTTPRequest.Path {
			if e.Verbose >= 3 {
				h, _, _ := net.SplitHostPort(ctx.RemoteAddr().String())
				log.Printf("%s (%s)", path, h)
			} else if e.Verbose >= 2 {
				log.Println(path)
			}
			contentType := k.HTTPRequest.Body.ContentType
			if contentType == "" {
				contentType = strApplicationJSON
			} else if regexp.MustCompile(`(?i)application/x-www-form-urlencoded;`).MatchString(contentType) {
				contentType = "text/html; charset=UTF-8"
			}
			ctx.Response.Header.SetCanonical(strContentType, []byte(contentType))

			statusCode := k.HTTPResponse.StatusCode
			if statusCode == 0 {
				statusCode = 200
			}
			ctx.Response.SetStatusCode(statusCode)

			// directly write to body
			fmt.Fprintf(ctx, k.HTTPResponse.Body)

			return
		}
	}
	ctx.Error("Path not found", fasthttp.StatusInternalServerError)
}

func getObject(i interface{}, fieldName string) interface{} {
	return reflect.ValueOf(i).MapIndex(reflect.ValueOf(fieldName)).Interface()
}
