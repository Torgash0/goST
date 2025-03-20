package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

var (
	task   string
	taskMu sync.Mutex
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskMu.Lock()
	defer taskMu.Unlock()
	fmt.Fprintf(w, "Hello, %s!", task)
}

func SetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var Massage struct {
		massage string `json:"massage"`
	}

	if err := json.NewDecoder(r.Body).Decode(&Massage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskMu.Lock()
	defer taskMu.Unlock()
	task = Massage.massage
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/massage", GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/massage", SetTaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
