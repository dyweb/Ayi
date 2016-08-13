package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gocraft/web"
)

type Server struct {
	Root   string
	Port   int
	Router *web.Router
}

type emptyContext struct {
}

func NewStaticServer() *Server {
	// TODO: get root from config
	cwd, _ := os.Getwd()

	server := Server{}
	server.Root = cwd
	server.Port = 8000
	server.Router = web.New(emptyContext{})

	// TODO: does the middleware list folder?
	server.Router.Middleware(web.LoggerMiddleware)
	server.Router.Middleware(web.StaticMiddleware(server.Root, web.StaticOption{IndexFile: "index.html"}))
	return &server
}

func (server *Server) Run() {
	http.ListenAndServe(fmt.Sprintf("localhost:%d", server.Port), server.Router)
}
