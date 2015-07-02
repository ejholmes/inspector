package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Request struct {
	Method     string
	URL        *url.URL
	Header     http.Header
	Host       string
	RemoteAddr string
	RequestURI string
}

func newRequest(r *http.Request) *Request {
	return &Request{
		Method:     r.Method,
		URL:        r.URL,
		Header:     r.Header,
		Host:       r.Host,
		RemoteAddr: r.RemoteAddr,
		RequestURI: r.RequestURI,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	raw, err := json.MarshalIndent(newRequest(r), "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	io.Copy(w, bytes.NewReader(raw))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
