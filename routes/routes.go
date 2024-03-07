// routes.go

package routes

import (
	"net/http"

	"github.com/google/uuid"
)

func Index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
}

func Page(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/page.html")
}

func PageTwo(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/page2.html")
}

func HandleUUID(w http.ResponseWriter, r *http.Request) {
    uuid := GenerateUuid()
    _, err := w.Write([]byte(uuid))
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func GenerateUuid() string {
    id := uuid.New()
    uuidString := id.String()
    return uuidString
}