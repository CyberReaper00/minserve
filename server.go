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

func Load_Page(filename, target string) {

	http.HandleFunc(target, func(writer http.ResponseWriter, requestor *http.Request) {
		content, read_err := os.ReadFile(filename)

		if read_err != nil {
			http.Error(writer, "Could not read file: ", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		writer.Write(content)
	})
}

func main() {
	horz_top := str.Repeat("━", 47)
	horz_mid := str.Repeat("━", 70)
	horz_bot := str.Repeat("━", 81)

	flag.Usage = func() {
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
		return
	}
	
	flag.Parse()

	if len(os.Args) < 2 { fmt.Println("No port was provided..."); return; }

	homepage := ""
	if len(os.Args) < 3 { homepage = "index.html"
	} else { homepage = os.Args[2] }

	http.HandleFunc("/", func(writer http.ResponseWriter, requestor *http.Request) {
		homepage_content, hread_err := os.ReadFile(homepage)
		if hread_err != nil { log.Println("Homepage content could not be read..."); return; }

		if requestor.URL.Path != "/" {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(writer, "404 Not Found")
			return
		}
		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		writer.Write(homepage_content)
	})
	
	page_list := os.Args[2:]
	
	if len(page_list) > 0 {
		for _, name := range page_list {
			if !str.Contains(name, "html") { continue }
			
			basename := str.TrimSuffix(name, ".html")
			Load_Page(name, "/" + basename)
		}
		
	} else {
		page_names := []string{}

		files, read_err := os.ReadDir(".")
		if read_err != nil { panic(read_err) }

		for _, file := range files {
			page_names = append(page_names, file.Name())
		}
		
		if !hu.StrSliceContains(page_names[:], "index", "fuzzy") {
			log.Println("index.html was not found, exiting...")
			return
		}
		
		pages := []string{}
		for _, name := range page_names {
			if !str.Contains(name, "html") { continue }
			pages = append(pages, name)
		
			basename := str.TrimSuffix(name, ".html")
			Load_Page(name, "/" + basename)
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
