package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
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
		s := http.StatusOK
		b, e := ioutil.ReadAll(r.Body)
		if e != nil {
			s = http.StatusBadRequest
			m = "ng"
		} else {
			sig := r.Header.Get("X-Hub-Signature")
			if !verifySignature(b, token, sig) {
				s = http.StatusBadRequest
				m = "ng"
			}
		}
		w.WriteHeader(s)
		fmt.Fprintf(w, m)
	})
	e := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if e != nil {
		log.Fatalf("%+v", e)
	}
}

func verifySignature(payload []byte, token, signature string) bool {
	mac := hmac.New(sha1.New, []byte(token))
	mac.Write(payload)
	bs := mac.Sum(nil)
	e := hex.EncodeToString(bs)
	return signature == fmt.Sprintf("sha1=%s", e)
}
