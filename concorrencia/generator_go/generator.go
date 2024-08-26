package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura

func titulo(urls ...string) <-chan string {
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, _ := http.Get(url)
			html, _ := io.ReadAll(res.Body)

			r, _ := regexp.Compile(`<title>(.*?)<\\/title>`)
			ch <- r.FindStringSubmatch(string(html))[1]
			fmt.Println(r)
		}(url)
	}

	return ch
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com.br", "https://www.youtube.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
