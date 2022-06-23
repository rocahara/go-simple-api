package main

import "net/http"

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(`{"status":"ok","method":"GET"}`))
		default:
			http.Error(w, `{"status":"Error"}`, http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8888", nil)
}
