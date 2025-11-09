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
	"fmt"
	"log"
	"time"
	"flag"
ctx	"context"
sc	"syscall"
scv	"strconv"
str	"strings"
	"net/http"
	"os/signal"

hu	"local/main/humain"
)

var horz_top = str.Repeat("━", 47)
var horz_mid = str.Repeat("━", 70)
var horz_bot = str.Repeat("━", 81)

func usage() {
		fmt.Println("" +
		"┏━━━━━ usage: <port> [<filenames>] " + horz_top + "┓" +
		"\n┃\t\t\t\t\t\t\t\t\t\t  ┃" +
		"\n┣━\033[1;7m COMMANDS \033[0m" + horz_mid + "┫" +
		"\n┃  port\t\t\t\t\t\t\t\t\t\t  ┃" +
		"\n┃\tthe port of the target\t\t\t\t\t\t\t  ┃" +
		"\n┃  filenames\t\t\t\t\t\t\t\t\t  ┃" +
		"\n┃\tif no name(s)/path(s) are given for files to be loaded then MinServe\t  ┃" +
		"\n┃\twill look for 'index.html' in the current directory for the homepage,\t  ┃" +
		"\n┃\tif not found, MinServe will give an error - if found then all html files  ┃" +
		"\n┃\tpresent in the current directory will be parsed and a corresponding path  ┃" +
		"\n┃\twith the same name will be appended to the target port without the file\t  ┃" +
		"\n┃\textension\t\t\t\t\t\t\t\t  ┃" +
		"\n┃\t\t\t\t\t\t\t\t\t\t  ┃" +
		"\n┃\tif any name(s)/path(s) are given then only those will be parsed and\t  ┃" +
		"\n┃\ta corresponding path with the same name will be appended to the target\t  ┃" +
		"\n┃\tport without the file extension\t\t\t\t\t\t  ┃" +
		"\n┗" + horz_bot + "┛")

		flag.PrintDefaults()
	}

func Load_Page(filename, target string) {

	http.HandleFunc(target, func(writer http.ResponseWriter, requestor *http.Request) {
		content, read_err := os.ReadFile(filename)

		if read_err != nil {
			http.Error(writer, "Could not read file: ", http.StatusInternalServerError); return }

		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		writer.Write(content)
	})
}

var static_ext_list = []string{
	".js", ".css", ".php", ".jpg", ".jpeg", ".png",
	".gif", ".ico", ".svg", ".json", ".xml", ".txt",
}

func is_static(filename string) bool {
	name := str.ToLower(filename)

	for _, ext := range static_ext_list {
		if str.HasSuffix(name, ext) { return true }
	}
	return false
}

func main() {

	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 { fmt.Println("No port was provided..."); return }

	var homepage string
	if len(os.Args) < 3 { homepage = "index.html"
	} else { homepage = os.Args[2] }

	http.HandleFunc("/", func(writer http.ResponseWriter, requestor *http.Request) {
		homepage_content, hread_err := os.ReadFile(homepage)
		if hread_err != nil { log.Println("Homepage content could not be read..."); return }

		if requestor.URL.Path != "/" {
			_404_page, _ := os.ReadFile("page_not_found.html")
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(writer, string(_404_page))
			return
		}

		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		writer.Write(homepage_content)
	})

	page_list := os.Args[2:]

	page_names := []string{}
	file_list, read_err := os.ReadDir(".")
	if read_err != nil { panic("Error occured reading dir files") }

	for _, file := range file_list { page_names = append(page_names, file.Name()) }
	
	if len(page_list) > 0 {
		for _, name := range page_list {
			if !str.Contains(name, "html") || !str.Contains(name, ".js") { continue }

			if str.Contains(name, "html") {
				basename := str.TrimSuffix(name, ".html")
				Load_Page(name, "/" + basename)
			}
		}

	} else {
		if !hu.StrSliceContains(page_names[:], "index") {
			log.Println("index.html was not found, exiting..."); return }
		
		pages := []string{}
		for _, name := range page_names {
			if !str.Contains(name, "html") || !str.Contains(name, ".js") { continue }
			if str.Contains(name, "page_not_found") { continue }

			pages = append(pages, name)
		}

		hu.Print_List("The following pages have been activated:", pages)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, sc.SIGTERM)

	port, arg_err := scv.Atoi(os.Args[1])
	if arg_err != nil { log.Println("Unknown value entered, port must be int..."); return; }

	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go func() {
		if len(page_list) > 0 {
			hu.Print_List("The following pages have been activated:", page_list)
		}
		fmt.Printf("Server running on: http://localhost:%d\n\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port(%v): %v\n", port, err)
		}
	}()

	<-stop

	log.Println("Shutting down server...\n")

	ctx, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Shutdown forced...\n")
	}

	log.Println("Shutdown complete")
}
