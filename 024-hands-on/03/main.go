// Serve the files in the "starting-files" folder
// Use "http.FileServer"

package main

import (
	"log"
	"net/http"
)

func main() {

	// serve up everything in this current directory

	//my way
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// http.ListenAndServe(":8080", nil)

	// todd's way
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))

}
