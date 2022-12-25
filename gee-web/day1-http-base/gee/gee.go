package gee

import (
	"fmt"
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]handlerFunc
}

func New() *Engine{
	return &Engine{
		make(map[string]handlerFunc),
	}
}

func (engine *Engine) addRoute(method, path string, handler handlerFunc){
	key := method + "-" + path
	engine.router[key] = handler
}

func (engine *Engine) Get(path string, handler handlerFunc) {
	engine.addRoute("GET", path, handler)
}

func (engine *Engine) Post(path string, handler handlerFunc) {
	engine.addRoute("POST", path, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, ok := engine.router[req.Method + "-" + req.URL.Path]; if ok{
		handler(w, req)
	}else {
		w.WriteHeader(404)
		w.Write([]byte("page not found"))
	}
}

func (engine *Engine) Run(addr string) {
	if err := http.ListenAndServe(addr, engine); err != nil {
		panic(fmt.Sprintf("http server run failed:%v/n", err))
	}
}
