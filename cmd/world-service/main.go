package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/hello", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("Hello World"))
	})

	fmt.Println("Server Started at http://localhost:3000")

	http.ListenAndServe(":3000", router)
}
