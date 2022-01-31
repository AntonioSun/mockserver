package main

import (
	"encoding/json"
	"log"
	"reflect"
	"regexp"

	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = "application/json"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	for _, k := range respMap {
		if path == k.HTTPRequest.Path {
			if e.Verbose >= 2 {
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
