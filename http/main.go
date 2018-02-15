package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	var httpPort = os.Getenv("HTTP_PORT")
	http.HandleFunc("/", homeHandle)
	http.ListenAndServe(":"+httpPort, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
