package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("SECRET_TOKEN")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m := "ok"
		b := r.Body
		if !verifySignature(b, token) {
			m = "ng"
		}
		fmt.Fprintf(w, m)
	})
	http.ListenAndServe(":8080", nil)
}

func verifySignature(payload interface{}, token string) bool {
	return true
}
