package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var webhooks []interface{}

//All Webhooks
func getWebhooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(webhooks)
}

//Recieve New Webhook
func createWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var webhook map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&webhook)
	webhooks = append(webhooks, webhook)
	json.NewEncoder(w).Encode(webhooks)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	// port := os.Getenv("PORT")

	//Route Handlers / Endpoints
	r.HandleFunc("/api/webhooks", getWebhooks).Methods("GET")
	r.HandleFunc("/api/webhooks", createWebhook).Methods("POST")

	// log.Fatal(http.ListenAndServe(":"+port, r))

	log.Fatal(http.ListenAndServe(":8080", r))
}
