package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// q := r.URL.Query().Get("test")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"foo":"bar"}`))
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Membaca body request
		byteBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "invalid body request", http.StatusBadRequest)
		}

		// Konversi JSON body ke map
		var m map[string]any
		json.Unmarshal(byteBody, &m)
		fmt.Println("Nama", m["nama"])

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteBody)
	})

	fmt.Println("server is running")
	http.ListenAndServe(":4444", mux)
}
