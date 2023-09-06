package belajar_golang_web

import (
	"fmt"
	"net/http"
)

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "https://www.programmerzamannow.com/", http.StatusTemporaryRedirect)
}

func TestRedirect() {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
