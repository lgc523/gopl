package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	UrlPrefixHttp  = "http://"
	UrlPrefixHttps = "https://"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchWithChan(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchWithChan(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	var fileName = start.String() + ".txt"
	out, err := os.Create(fileName)
	defer out.Close()
	io.Copy(out, resp.Body)
	//nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	//defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t\t%s", secs, url)
}

func simpleFetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, UrlPrefixHttp) || !strings.HasPrefix(url, UrlPrefixHttps) {
			url = UrlPrefixHttp + url
		}
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("RESP:%s\t%s\t%s\n", resp.Status, resp.Proto, resp.Header.Get("Location"))

		//b, err := ioutil.ReadAll(resp.Body)
		out, err := os.Create("fetch.txt")
		defer out.Close()
		io.Copy(out, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
