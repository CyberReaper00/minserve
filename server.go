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
	"fmt"
	"path"
	_ "time"
	"strings"
	_ "reflect"
	"net/http"

	_ "local/main/humain"
)

func Reload_Server(root string, files []string, not_found_contents []byte, port string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {

		// TODO
		// create a goroutine that allows the server to walk through the current
		// directory and updated the list of files to search through when the user
		// presses ctrl-r
		var file_path	string
		var contents	[]byte
		var file_err	error

		if req.URL.Path == "/" { file_path = "index"
		} else { file_path = path.Join(root, path.Clean(req.URL.Path)) }

		basename := strings.TrimPrefix(file_path, "/")
		for _, item := range files {
			if !strings.Contains(item, basename) { continue }

			_, after, _ := strings.Cut(item, ".")
			if after == "" { continue }

			switch after {
				case "html":
					contents, file_err = os.ReadFile(file_path + ".html")
				default:
					contents, file_err = os.ReadFile(file_path)
			}
		}

		if file_err != nil {
			fmt.Printf("%s was not found for: %s\n", file_path, req.URL.Path)
			writer.WriteHeader(http.StatusNotFound)
			writer.Write(not_found_contents)
			return
		}

		fmt.Printf("%s was found for: %s\n", file_path, req.URL.Path)

		writer.Header().Set("Content-Type", http.DetectContentType(contents))
		writer.Write(contents)
	})
}

func Read_All_Dirs(list []os.DirEntry, files []string, dir string) []string {
	for _, item := range list {
		
		if strings.HasPrefix(item.Name(), ".") { continue }

		if item.IsDir() {
			list, _ = os.ReadDir(item.Name())
			files = Read_All_Dirs(list, files, item.Name())

		} else {
			if dir == "" { files = append(files, item.Name())
			} else { files = append(files, fmt.Sprintf("%s/%s", dir, item.Name())) }
		}
	}

	return files
}

func main() {
	if len(os.Args) < 2 { log.Fatalln("No port was provided...") }
	port := os.Args[1]

	_, main_err := os.ReadFile("index.html")
	if main_err != nil { log.Fatalln("index.html was not found, exiting...") }

	var files []string
	root_files, dir_err := os.ReadDir(".")
	if dir_err != nil { log.Fatalln("ERROR: ", dir_err) }

	all_files := Read_All_Dirs(root_files, files, "")

	content, _ := os.ReadFile("page_not_found")
	reload_handler := Reload_Server(".", all_files, content, port)

	http.Handle("/", reload_handler)

	log.Println("Server started at port: " + port)
	// TODO: create an implementation that allows the user to host their site on a
	// registered domain
	http.ListenAndServe(":" + port, nil)
}
