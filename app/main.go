package main

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	response := map[string]string{
		"timestamp": time.Now().Format(time.RFC3339),
		"ip":        ip,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
