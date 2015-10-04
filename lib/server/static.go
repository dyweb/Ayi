package server

import (
	"strconv"
	"net/http"
	"log"
)

func Run(folder string, port int) {
	log.Fatal(http.ListenAndServe("localhost:" + strconv.Itoa(port), http.FileServer(http.Dir(folder))))
}