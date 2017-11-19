package main

import (
	"encoding/json"
	"log"
	"net/http"

	// TODO use net/http/#ServeMux
	"github.com/gorilla/mux"
)

type Item struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Project *Project `json:"project,omitempty"`
}
type Project struct {
	Name  string `json:"name,omitempty"`
	Notes string `json:"notes,omitempty"`
}

// TODO pass by reference
var items []Item

func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	found := false
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			found = true
		}
	}
	if !found {
		// TODO use http status
		json.NewEncoder(w).Encode("Not found")
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = params["id"]
	items = append(items, item)
	json.NewEncoder(w).Encode(items)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	// TODO handle JSON errors
	json.NewEncoder(w).Encode(items)
}

func main() {
	items = append(items, Item{ID: "1", Name: "Milk", Project: &Project{Name: "Home", Notes: "Important"}})
	items = append(items, Item{ID: "2", Name: "Cofee", Project: &Project{Name: "Work", Notes: "Do not forget!"}})
	router := mux.NewRouter()

	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", GetItem).Methods("GET")
	router.HandleFunc("/items/{id}", CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
