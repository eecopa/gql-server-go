package main

import (
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {

	h := handler.New(&handler.Config{
		Schema:   &AcquisitionSchema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)
}
