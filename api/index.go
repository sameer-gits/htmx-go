package api

import (
	"net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "index.html")
}