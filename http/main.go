package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var httpPort = os.Getenv("PORT")

func main() {
	http.HandleFunc("/", homeHandle)
	http.ListenAndServe(":"+httpPort, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
