package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the universe handler for all requests
type Engine struct {

}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprint(w, "404 not found!")
	}
}

func main(){
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
