package main

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	for _, k := range respMap {
		if path == k.HTTPRequest.Path {
			log.Println(path)
			ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

			statusCode := k.HTTPResponse.StatusCode
			if statusCode == 0 {
				statusCode = 200
			}
			ctx.Response.SetStatusCode(statusCode)

			if err := json.NewEncoder(ctx).Encode(k.HTTPResponse.Body); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}

			return
		}
	}
	ctx.Error("Path not found", fasthttp.StatusInternalServerError)
}

func getObject(i interface{}, fieldName string) interface{} {
	return reflect.ValueOf(i).MapIndex(reflect.ValueOf(fieldName)).Interface()
}
