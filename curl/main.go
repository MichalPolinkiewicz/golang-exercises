package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, r := range os.Args[1:] {
		if !strings.HasPrefix(r, "http://") && !strings.HasPrefix(r, "https://") {
			r = "http://" + r
		}
		resp, err := http.Get(r)
		if err != nil {
			log.Fatal("error during get:", err.Error())
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal("error during copy", err.Error())
		}
		fmt.Println("\nSTATUS CODE:", resp.StatusCode)
	}
}
