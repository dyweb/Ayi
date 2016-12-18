package web

import (
	"fmt"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gocraft/web"
)

type Server struct {
	Root   string
	Port   int
	Router *web.Router
}

type ayiContext struct {
}

func (c *ayiContext) Index(rw web.ResponseWriter, req *web.Request) {
	// NOTE: http.ServeFile(rw, req, content io.ReadSeeker ) is not supported, since when using append zip
	// Downsides for appending ... does not provide a working Seek method.
	box := rice.MustFindBox("app-web-public")
	indexHTML, err := box.String("index.html")
	if err != nil {
		fmt.Fprintf(rw, "Error: index html not found!")
	} else {
		fmt.Fprint(rw, indexHTML)
	}
}

func NewAyiServer(port int) *Server {
	server := Server{}
	server.Port = port
	server.Router = web.New(ayiContext{})

	server.Router.Middleware(web.LoggerMiddleware)
	box := rice.MustFindBox("app-web-public")
	// NOTE: index file does not work, because isDir return false
	server.Router.Middleware(web.StaticMiddlewareFromDir(box.HTTPBox(), web.StaticOption{}))
	server.Router.Get("/", (*ayiContext).Index)
	return &server
}

type emptyContext struct {
}

func NewStaticServer(root string, port int) *Server {
	server := Server{}
	server.Root = root
	server.Port = port
	server.Router = web.New(emptyContext{})

	// NOTE: the middleware does NOT list folder, it is said to avoid content length problem
	server.Router.Middleware(web.LoggerMiddleware)
	server.Router.Middleware(web.StaticMiddleware(server.Root, web.StaticOption{IndexFile: "index.html"}))
	return &server
}

func (server *Server) Run() {
	log.Info("Start server on " + fmt.Sprintf("localhost:%d", server.Port))
	http.ListenAndServe(fmt.Sprintf("localhost:%d", server.Port), server.Router)
}
