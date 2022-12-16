package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "10.0.0.40:8081",
	})
	http.ListenAndServe(":8081", proxy)
}
