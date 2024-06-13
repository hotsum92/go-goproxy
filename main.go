package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/elazarl/goproxy.v1"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	//proxy.Verbose = true

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			body, _ := io.ReadAll(r.Body)
			defer r.Body.Close()

			println(fmt.Sprintf("Request %s %s %s", r.Method, r.URL.String(), string(body)))
			return r, nil
		})

	proxy.OnResponse().DoFunc(
		func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			//body, _ := io.ReadAll(r.Body)
			//defer r.Body.Close()

			//println(fmt.Sprintf("Response %s", string(body)))
			return r
		})

	log.Fatal(http.ListenAndServe(":8080", proxy))
}
