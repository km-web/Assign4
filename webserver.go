package main

import (
	"fmt"
	"log"
	"net/http"
)

func WebServer(k http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(k, "KATHIR, %s!\n", r.URL.Path[1:])
}
func main() {

	http.HandleFunc("/", WebServer)
	fmt.Println("webServer kith Handling function")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
