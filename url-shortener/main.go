package main

import (
	"fmt"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if u, ok := pathsToUrls[req.URL.RequestURI()]; ok {
			http.Redirect(res, req, u, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(res, req)
		}
	}
}

func main() {
	router := http.DefaultServeMux

	urls := map[string]string{
		"/dogs": "en.wikipedia.org/wiki/Dog",
		"/cats": "en.wikipedia.org/wiki/Cat",
	}

	h := MapHandler(urls, router)
	err := http.ListenAndServe(":8086", h)

	if err != nil {
		fmt.Println(err)
	}
}
