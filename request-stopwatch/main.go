package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	c := make(chan statistics, len(os.Args)-1)
	start := time.Now()
	for _, url := range os.Args[1:] {
		go fetch(url, c)
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(<-c)
	}
	end := time.Since(start)
	fmt.Println("total time =", end.Seconds())
}

type statistics struct {
	url       string
	time      float64
	status    int
	readBytes int64
}

func (s statistics) String() string {
	return fmt.Sprintf("URL= %s || Time= %f || Status= %d || Size= %d", s.url, s.time, s.status, s.readBytes)
}

func fetch(url string, c chan statistics) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error during get:", err.Error())
		end := time.Since(start).Seconds()
		c <- statistics{url: url, time: end, status: http.StatusNotFound, readBytes: 0}
	} else {
		written, err := io.Copy(ioutil.Discard, resp.Body)
		err = resp.Body.Close()
		if err != nil {
			fmt.Println("error during reading", err.Error())
		}
		end := time.Since(start).Seconds()
		c <- statistics{url: url, time: end, status: resp.StatusCode, readBytes: written}
	}
}
