package controllers

import (
	"encoding/json"
	"fmt"
	"go-db-gorm/pkg/models"
	"go-db-gorm/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewItem models.TodoItem

func GetItems(w http.ResponseWriter, r *http.Request) {
	newItems := models.GetAllItems()

	res, _ := json.Marshal(newItems)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(res)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["Id"]
	ID, err := strconv.ParseInt(itemId, 0, 0)

	if err != nil {
		fmt.Println(("Error parsing data."))
	}

	itemDetails, _ := models.GetItemById(ID)
	res, _ := json.Marshal(itemDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(res)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	ItemCreator := &models.TodoItem{}
	utils.ParseBody(r, ItemCreator)

	newItem := ItemCreator.CreateItem()
	res, _ := json.Marshal(newItem)

	w.WriteHeader((http.StatusOK))
	w.Write(res)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["Id"]
	ID, err := strconv.ParseInt(itemId, 0, 0)

	if err != nil {
		fmt.Println(("Error parsing data."))
	}

	deletedItem := models.DeleteItem(ID)

	res, _ := json.Marshal(deletedItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(res)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updatedItem = &models.TodoItem{}
	utils.ParseBody(r, updatedItem)

	vars := mux.Vars(r)
	itemId := vars["Id"]
	ID, err := strconv.ParseInt(itemId, 0, 0)

	if err != nil {
		fmt.Println(("Error parsing data."))
	}

	// Get the item from the DB
	itemDetails, db := models.GetItemById(ID)

	// Update the info
	if updatedItem.Title != itemDetails.Title {
		itemDetails.Title = updatedItem.Title
	}

	if updatedItem.Description != itemDetails.Description {
		itemDetails.Description = updatedItem.Description
	}

	if updatedItem.Date_due != itemDetails.Date_due {
		itemDetails.Date_due = updatedItem.Date_due
	}

	if updatedItem.Date_added != itemDetails.Date_added {
		itemDetails.Date_added = updatedItem.Date_added
	}

	if updatedItem.IsDone != itemDetails.IsDone {
		itemDetails.IsDone = updatedItem.IsDone
	}

	db.Save(&itemDetails)

	res, _ := json.Marshal(itemDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	w.Write(res)
}
