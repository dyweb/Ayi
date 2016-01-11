package util

import (
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/spf13/viper"
)

// http://www.alexedwards.net/blog/golang-response-snippets
// http://www.alexedwards.net/blog/serving-static-sites-with-go
// http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang

// ServeStatic start a static server use folder as public directory on port
func ServeStatic() {
	port := viper.GetInt("port")
	log.Print("start on localhost:" + strconv.Itoa(port))
	http.HandleFunc("/", serveFileWithCORS)
	http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
}

// all the response include CORS header
func serveFileWithCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Apiache")
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fp := path.Join(viper.GetString("base"), r.URL.Path)
	// TODO: log the real path
	// TODO: use http request log library
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

	// Return index.html for a folder
	if info.IsDir() {
		fp = path.Join(fp, "index.html")
		log.Print(fp)
		_, err = os.Stat(fp)
		if err == nil {
			http.ServeFile(w, r, fp)
			return
		}

		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		w.WriteHeader(400)
		return
	}

	http.ServeFile(w, r, fp)
}
