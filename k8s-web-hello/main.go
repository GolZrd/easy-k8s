package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/k8s", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Kubernetes!")
	})

	http.ListenAndServe(":8080", nil)
}
