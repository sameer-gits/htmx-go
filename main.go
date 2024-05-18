// main.go

package main

import (
	"fmt"
	"net/http"
    "os"
	"github.com/sameer-gits/htmx-go/routes"
)

func main() {
    // added favicon made public folder available
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
    http.HandleFunc("/", routes.Index)
    http.HandleFunc("/page", routes.Page)
    http.HandleFunc("/page2", routes.PageTwo)
    http.HandleFunc("/uuid", routes.HandleUUID)
    fmt.Println("Server is running on port http://localhost:8080")
    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }

    http.ListenAndServe("0.0.0.0:" + port, nil)
}
