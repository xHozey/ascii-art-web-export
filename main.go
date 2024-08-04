package main

import (
	"fmt"
	"net/http"
	"os"

	web "ascii_art_web/web_server"
)

// Register the root URL and start HTTP server
func main() {
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", web.NotFoundHandler(fs)))
	http.HandleFunc("/", web.Index)
	http.HandleFunc("/ascii-art", web.Generate)
	http.HandleFunc("/download", web.Download)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
