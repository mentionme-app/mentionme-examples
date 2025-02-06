package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Item struct {
	Backend    string    `json:"backend"`
	BackendSub string    `json:"backend_sub"`
	Type       string    `json:"type"`
	IconURL    string    `json:"icon_url"`
	Timestamp  time.Time `json:"timestamp"`

	ItemURL      string `json:"item_url"`
	Author       string `json:"author"`
	ParentAuthor string `json:"parent_author"`

	Text      string `json:"text"`
	Title     string `json:"title"`
	TitleType int    `json:"title_type"`

	Lang   string `json:"lang"`
	Filter string `json:"filter"`
}

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "No body", http.StatusBadRequest)
		return
	}

	var items []Item
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received %d items\n", len(items))
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/syften-webhook", ItemsHandler)
	log.Println("Server started on :5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
