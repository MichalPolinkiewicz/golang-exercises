package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
		if u, ok := pathsToUrls[req.RequestURI]; ok {
			http.Redirect(res, req, u, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(res, req)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YAMLHandler(f []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var container []AddressFromYaml
	err := decodeYaml(f, &container)
	if err != nil {
		log.Println("error during decoding yaml", err)
	}
	addresses := createMap(&container)

	return func(res http.ResponseWriter, req *http.Request) {
		if u, ok := addresses[req.RequestURI]; ok {
			http.Redirect(res, req, u, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(res, req)
		}
	}, err
}

func decodeYaml(b []byte, out *[]AddressFromYaml) error {
	err := yaml.Unmarshal(b, out)
	return err
}

func createMap(addresses *[]AddressFromYaml) map[string]string {
	out := make(map[string]string)
	for _, r := range *addresses {
		out[r.Path] = r.Url
	}
	return out
}

type AddressFromYaml struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func main() {
	//default router and some url's
	router := http.DefaultServeMux
	urls := map[string]string{
		"/dogs": "xxx",
		"/cats": "zzz",
	}

	//reading url's from yaml
	f, err := ioutil.ReadFile("urls.yaml")
	if err != nil {
		log.Println("error during opening yaml", err)
	}

	//decorating default handler with yaml handler
	h, err := YAMLHandler(f, router)
	if err != nil {
		log.Println("error during opening yaml", err)
	}

	//decorating with map handler
	h = MapHandler(urls, h)
	err = http.ListenAndServe(":8086", h)
	if err != nil {
		log.Println(err)
	}
}
