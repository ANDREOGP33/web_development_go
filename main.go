package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "page not found", http.StatusNotFound)

	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, `<h1>Contact Page</h1>
	<p>To get in touch, email me at <a href=\\"mailto:andre.gasparoni@outlook.com\\">andre.gasparoni@outlook.com</a></p>`)

}

//2.8 The http.Handler Type
