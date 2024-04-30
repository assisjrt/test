package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	timezone, _ := time.LoadLocation("America/Sao_Paulo")
	t := time.Now().In(timezone).Format(time.RFC3339)
	hostname, _ := os.Hostname()

	data := map[string]any{
		"timestamp": t,
		"message":   hostname,
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintln(w, string(dataJSON)); err != nil {
		log.Fatalln(err.Error())
	}

	// json.NewEncoder(os.Stdout).Encode(data)
	// json.NewEncoder(w).Encode(data)
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	documentRoot := os.Getenv("APP_DOCUMENT_ROOT")
	if documentRoot == "" {
		log.Fatalln("Use: APP_DOCUMENT_ROOT='xxxx' " + os.Args[0])
	}

	handler := http.StripPrefix("/", http.FileServer(http.Dir(documentRoot)))
	http.Handle("/", handler)
	http.HandleFunc("/healthcheck", healthcheckHandler)
	log.Printf("Starting server on %s port...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
