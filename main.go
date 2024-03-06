// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/sameer-gits/htmx-go/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
