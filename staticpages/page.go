package staticpages

import (
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	//   http.ServeFile(w, r, "static/page.html")
	http.ServeFile(w, r, "static/page.html")
}

func PageTwo(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "static/page2.html")
}