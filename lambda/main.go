package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/morelj/lambada"
)

func main() {
	lambada.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		baseUrl := os.Getenv("TARGET_BASE_URL")

		if len(baseUrl) == 0 {
			w.Write(([]byte)("<html><body><h1>TARGET_BASE_URL environment variable is not set</h1></body></html>"))
			return
		}

		target, err := url.Parse(baseUrl)
		if err != nil {
			w.Write(([]byte)(fmt.Sprintf("<html><body><h1>Error parsing base url: %s</h1></body></html>", err.Error())))
			return
		}
		fmt.Println(r.URL.Path)
		fmt.Println(r.URL.RawQuery)
		proxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				path := req.URL.Path
				if !strings.HasPrefix(path, "/") && len(path) > 0 {
					path = "/" + path
				}
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host
				req.Host = target.Host
				req.URL.Path = path
				req.Header.Set("X-Forwarded-Host", req.Host)
			},
		}
		proxy.ServeHTTP(w, r)
	}))
}
