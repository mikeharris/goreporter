package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my super awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@goreporter.com\">support@goreporter.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1>Oops! We could not find the page you were looking for.</h1>")
	}
}
