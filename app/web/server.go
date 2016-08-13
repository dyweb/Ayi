package web

import (
	"fmt"
	"net/http"

	"github.com/dyweb/Ayi/util"
	"github.com/gocraft/web"
)

var log = util.Logger

type Server struct {
	Root   string
	Port   int
	Router *web.Router
}

type emptyContext struct {
}

func NewStaticServer(root string, port int) *Server {
	// TODO: may get the config from viper
	server := Server{}
	server.Root = root
	server.Port = port
	server.Router = web.New(emptyContext{})

	// TODO: does the middleware list folder? NO
	server.Router.Middleware(web.LoggerMiddleware)
	server.Router.Middleware(web.StaticMiddleware(server.Root, web.StaticOption{IndexFile: "index.html"}))
	return &server
}

func (server *Server) Run() {
	log.Info("Start server on " + fmt.Sprintf("localhost:%d", server.Port))
	http.ListenAndServe(fmt.Sprintf("localhost:%d", server.Port), server.Router)
}
