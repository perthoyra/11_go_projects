package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TodoItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date_due    string `json:"date_due"`
	Date_added  string `json:"date_added"`
	IsDone      bool   `j́son:"isdone"`
}

var items []TodoItem

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range items {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)

	var item TodoItem
	_ = json.NewDecoder(r.Body).Decode(&item)

	item.Id = strconv.Itoa(rand.Intn(100000000))
	items = append(items, item)

	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range items {
		if item.Id == params["id"] {
			// pretty shitty way to do this, but we do it just for testing
			// delete the item and then add the update item it back in
			items = append(items[:index], items[index+1:]...)

			var updatedItem TodoItem
			_ = json.NewDecoder(r.Body).Decode(&updatedItem)
			updatedItem.Id = params["id"]
			items = append(items, updatedItem)
			json.NewEncoder(w).Encode(updatedItem)
			break
		}
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range items {
		if item.Id == params["id"] {
			// copy the item(s) before the found index and then add the item(s) after the found index
			// => removing the item at the found index
			items = append(items[:index], items[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(items)
}

func main() {

	router := mux.NewRouter()

	items = append(items, TodoItem{Id: "1", Title: "Default item", Description: "Description goes here.", Date_due: "2026-07-10", Date_added: "2026-07-08", IsDone: false})

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	fmt.Println("Starting up server on :9090")

	log.Fatal(http.ListenAndServe(":9090", router))
}
