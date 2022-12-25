package gee

import (
	"fmt"
	"net/http"
)

func main(){
	r := New()
	r.Get("/", indexHandler)
	r.Get("/hello", helloHandler)
	r.Run(":9999")
}


func indexHandler(w http.ResponseWriter, req  *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

