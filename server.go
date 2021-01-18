package main

import (
	"net/http"
	"os"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return ":" + value
}

func main() {
	port := getenv("PORT", ":8080")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(port, nil); err != nil {
	  panic(err)
	}
}