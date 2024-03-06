// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/sameer-gits/htmx-go/api"
	"github.com/sameer-gits/htmx-go/staticpages"
)

func main() {
	http.HandleFunc("/", api.Handler)
	http.HandleFunc("/page", staticpages.Page)
	http.HandleFunc("/page2", staticpages.PageTwo)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
