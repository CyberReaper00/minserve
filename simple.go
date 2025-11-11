package main

import (
	"os"
	"log"
	"path"
	"strings"
	"net/http"
)

func Dynamic_Server(root string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {

		var file_path string
		if req.URL.Path == "/" { file_path = "index.html"
		} else { file_path = path.Join(root, path.Clean(req.URL.Path)) }

		contents, file_err := os.ReadFile(file_path)

		if file_err != nil || strings.HasPrefix(string(req.URL.Path), "/.") {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("404 File Not Found"))
			return
		}
		log.Printf("File served: %s\n", file_path)

		writer.Header().Set("Content-Type", http.DetectContentType(contents))
		writer.Write(contents)
	})
}

func main() {
	_, main_err := os.ReadFile("index.html")
	if main_err != nil { log.Fatalln("index.html was not found, exiting...") }

	http.Handle("/", Dynamic_Server("."))

	log.Println("Server started at port: 12323")
	http.ListenAndServe(":12323", nil)
}
