package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world! We are working!")
	})
	http.HandleFunc("/k8s", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Kubernetes!")
	})
	http.HandleFunc("/nginx", func(w http.ResponseWriter, r *http.Request) {
		url := "http://nginx"
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(body)
	})

	http.ListenAndServe(":8080", nil)
}
