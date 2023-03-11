package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		fmt.Println("loading env failed")
// 	}

// 	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
// 		Scheme: "http",
// 		Host:   os.Getenv("HOST"),
// 		Path:   "*/*",
// 	})
// 	http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), proxy)
// }

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("loading env failed")
	}

	// initialize a reverse proxy and pass the actual backend server url here
	// proxy, err := NewProxy(os.Getenv("HOST"))
	// if err != nil {
	// 	panic(err)
	// }

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   os.Getenv("HOST"),
	})

	// handle all requests to your server using the proxy
	http.HandleFunc("/", ProxyRequestHandler(proxy))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil))
}
