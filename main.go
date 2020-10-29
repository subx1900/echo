package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/honeybadger-io/honeybadger-go"
)

type Response struct {
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
	Path    string      `json:"url"`
	Method  string      `json:"method"`
}

func echo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)

	response := &Response{
		Headers: req.Header,
		Body:    body,
		Path:    req.URL.Path,
		Method:  req.Method,
	}

	json, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	http.HandleFunc("/", echo)

	print("Listening on 3000")
	var handler http.Handler
	log.Fatal(http.ListenAndServe(":3000", honeybadger.Handler(handler)))
}
