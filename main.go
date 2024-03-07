// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/sameer-gits/htmx-go/routes"
)

func main() {
    http.HandleFunc("/", routes.Index)
    http.HandleFunc("/page", routes.Page)
    http.HandleFunc("/page2", routes.PageTwo)
    http.HandleFunc("/uuid", routes.HandleUUID)
    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}