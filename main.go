package main

import (
	"fmt"
	//io over bufio because i wanted to read the body text into the byte in one line/operation?
	"io"
	"net/http"
)

func main() {
	testUrl := "https://usf-cs272-f25.github.io/test-data/project01/"
	//testUrlTwo := "https://usf-cs272-f25.github.io/test-data/project01/href.html"
	website, err := http.Get(testUrl)
	if err != nil {
		fmt.Println("Error accessing website")
		return
	}
	defer website.Body.Close()

	host := website.Request.URL.String()

	fmt.Println("Host URL: ", host)
	body, bodyErr := io.ReadAll(website.Body)
	if bodyErr != nil {
		fmt.Println("Error accessing body of website")
		return
	}
	words, hrefs := Extract(body)

	fmt.Printf("Words: %s\n", words)
	fmt.Printf("HREFS: %s\n", hrefs)

	//hrefsTest := []string{"/", "/help/", "/syllabus/", "https://gobyexample.com/"}

	cleanUrls := Clean(&host, &hrefs)
	fmt.Println(cleanUrls)
}
