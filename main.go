package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("SECRET_TOKEN")
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := "ok"
		b := r.Body
		if !verifySignature(b, token) {
			m = "ng"
		}
		fmt.Fprintf(w, m)
	})
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func verifySignature(payload interface{}, token string) bool {
	return true
}
