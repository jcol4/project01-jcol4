package main

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func Extract(bodyText []byte) ([]string, []string) {
	text, err := html.Parse(bytes.NewReader(bodyText))
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	words := []string{}
	hrefs := []string{}
	for n := range text.Descendants() {
		if n.Type == html.TextNode {
			for _, word := range strings.Fields(n.Data) {
				words = append(words, word)
			}
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					hrefs = append(hrefs, attr.Val)
				}
			}
		}
	}
	return words, hrefs
}
