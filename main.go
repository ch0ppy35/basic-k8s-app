package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type PodNameResponse struct {
	PodName string `json:"podName"`
}

type HealthCheckResponse struct {
	Status bool `json:"status"`
}

func createResponse(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	message := MessageResponse{Message: "Hello World!"}
	createResponse(w, message)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	message := MessageResponse{Message: "Yay, K8s!"}
	createResponse(w, message)
}

func podNameHandler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("MY_POD_NAME")
	if name == "" {
		name = "Not running in K8s!"
	}
	podName := PodNameResponse{PodName: name}
	createResponse(w, podName)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status := HealthCheckResponse{Status: true}
	createResponse(w, status)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/pod_name", podNameHandler)
	http.HandleFunc("/healthz", healthCheckHandler)

	log.Println("Server listening at http://127.0.0.1:8080/")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
