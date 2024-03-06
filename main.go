// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/sameer-gits/htmx-go/handler" // Import the index.go file from the htmx-go repository
)

func main() {
    http.HandleFunc("/", handler.Handler)
    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}
