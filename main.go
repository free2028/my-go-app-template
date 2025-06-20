package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

var version = "dev" // 这个值会在构建时被替换

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	// 路由定义
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/api/info", infoHandler).Methods("GET")

	fmt.Printf("Server starting on port %s...\n", port)
	fmt.Printf("Version: %s\n", version)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
    <html>
    <head><title>My Go App</title></head>
    <body>
        <h1>Welcome to My Go App!</h1>
        <p>Version: ` + version + `</p>
        <p><a href="/health">Health Check</a></p>
        <p><a href="/api/info">API Info</a></p>
    </body>
    </html>`
	w.Write([]byte(html))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{
		Status: "OK",
		Time:   time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message:   "Hello from My Go App!",
		Timestamp: time.Now(),
		Version:   version,
	}
	json.NewEncoder(w).Encode(response)
}
