package server

import (
	"strconv"
	"net/http"
	"log"
	"os"
	"path"
)

// http://www.alexedwards.net/blog/golang-response-snippets
// http://www.alexedwards.net/blog/serving-static-sites-with-go
// http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang

func Run(folder string, port int) {
	//	log.Fatal(http.ListenAndServe("localhost:" + strconv.Itoa(port), http.FileServer(http.Dir(folder))))
	http.HandleFunc("/", serveFileWithCORS)
	http.ListenAndServe("localhost:" + strconv.Itoa(port), nil)
}

func serveFileWithCORS(w http.ResponseWriter, r *http.Request) {
	// Apiache!
	w.Header().Set("Server", "Apiache")
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fp := path.Join("front/public_html", r.URL.Path)
	log.Print(fp)

	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		// TODO: other error here.
		w.WriteHeader(400)
		return
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		fp = path.Join(fp, "index.html")
		log.Print(fp)
		_, err = os.Stat(fp)
		if err == nil {
			http.ServeFile(w, r, fp)
		}else {
			if os.IsNotExist(err) {
				http.NotFound(w, r)
				return
			}
		}
		w.WriteHeader(400)
		return
	}

	http.ServeFile(w, r, fp)
}