package main

import (
	"encoding/json"
	"log"
	"net/http"

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

var items []Item

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	found := false
	for _, item := range items {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			found = true
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not found")
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = params["id"]
	items = append(items, item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(items)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(items)
}

func main() {
	items = append(items, Item{ID: "1", Name: "Milk", Project: &Project{Name: "Home", Notes: "Important"}})
	items = append(items, Item{ID: "2", Name: "Cofee", Project: &Project{Name: "Work", Notes: "Do not forget!"}})
	router := mux.NewRouter()

	router.HandleFunc("/items", Index).Methods("GET")
	router.HandleFunc("/items/{id}", Show).Methods("GET")
	router.HandleFunc("/items/{id}", Create).Methods("POST")
	router.HandleFunc("/items/{id}", Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
