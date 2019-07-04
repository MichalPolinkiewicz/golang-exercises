package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

var links []link

type link struct {
	Href string
	Text string
}

func main() {
	pageMock := `<p>Links:</p>
			<ul>
			<li>
			<a href="foo">Foo</a><li>
			<a href="/bar/baz">BarBaz</a>
			</ul>
			<a href="/janusz">aaa</a>`

	doc, err := html.Parse(strings.NewReader(pageMock))
	if err != nil {
		log.Fatal(err)
	}

	getLinks(doc)
	fmt.Println(links)
}

func getLinks(n *html.Node) {
	if n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, link{a.Val, n.FirstChild.Data})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getLinks(c)
	}
}
