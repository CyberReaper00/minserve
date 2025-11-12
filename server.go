/*
MinServe is a simple but powerful web server that can be used for the
hosting of static pages in a simple and straight-forward manner

Copyright (C) 2025 Nahyan ul haq Siddiqui

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

Contact at <simple.nahyan@gmail.com>
*/

package main

import (
	"os"
	"log"
	"path"
	"strings"
	_ "reflect"
	"net/http"
)

// TODO: figure out a way to get rid of dir in the parameters for Dynamic_Server
// so the user only has to enter one thing and all other functionality can be infered
// from that
func Dynamic_Server(root string, files []os.DirEntry, not_found_contents []byte, port string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {

		// TODO: create a loop that gets all the files from the current directory and then
		// checks request path to see if path matches any filenames in the directory list which
		// will allow the user to create and remove files and directories without restarting
		// the server - this will cause each request to be slower since the directory is being
		// scanned on each request
		// or
		// remove the root parameter and only keep the dir parameter so the directory is only
		// scanned once and never again - this will make each request much faster but will cause
		// an annoyance of having to restart the server everytime some changes are made to the
		// directory
		var file_path	string
		var contents	[]byte
		var file_err	error

		if req.URL.Path == "/" { file_path = "index"
		} else { file_path = path.Join(root, path.Clean(req.URL.Path)) }

		// TODO: add function to only serve html files without the extension and every other
		// file should be served as is
		basename := strings.TrimPrefix(file_path, "/")
		for _, file := range files {
			if !strings.Contains(file.Name(), basename) { continue }

			_, after, _ := strings.Cut(file.Name(), ".")
			if after == "" { continue }

			switch after {
				case "html": contents, file_err = os.ReadFile(file_path + ".html")
				default:
					contents, file_err = os.ReadFile(file_path)
			}
		}

		if file_err != nil {
			// TODO: replace this manual implementation with either http.NotFound or
			// http.NotFoundHandler to make it more efficient
			log.Printf("Request could not be intercepted for: localhost:%s%v\n", port, req.URL.Path)
			writer.WriteHeader(http.StatusNotFound)
			writer.Write(not_found_contents)
			return
		}

		log.Printf("Request intercepted for: localhost:%s%v\n", port, req.URL.Path)

		writer.Header().Set("Content-Type", http.DetectContentType(contents))
		writer.Write(contents)
	})
}

func main() {
	if len(os.Args) < 2 { log.Fatalln("No port was provided, exiting...") }
	port := os.Args[1]

	// TODO: create an implementation that allows the user to define the name of the homepage
	// file instead of using index.html and keep index.html as the default homepage file in case no
	// name is provided
	_, main_err := os.ReadFile("index.html")
	if main_err != nil { log.Fatalln("index.html was not found, exiting...") }

	files, _ := os.ReadDir(".")
	content, _ := os.ReadFile("page_not_found")
	dynamic_handler := Dynamic_Server(".", files, content, port)

	http.Handle("/", dynamic_handler)

	log.Println("Server started at port: " + port)
	http.ListenAndServe(":" + port, nil)
}
