package models

import (
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int64     `json:id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueAt       string    `json:"date_due"`
	CreatedAt   time.Time `json:"date_added"`
	UpdatedAt   time.Time `json:"date_updated"`
	IsDone      bool      `j́son:"isdone"`
}

func (todo *TodoItem) CreateTodoItem() *TodoItem {
	var newItem TodoItem

	newItem.CreatedAt = time.Now()
	newItem.Description = ""
	newItem.IsDone = false
	newItem.Title = ""
	newItem.UpdatedAt = time.Now()

	todo = &newItem
	return todo
}

func LoadTodoItemList(todoItemList *[]TodoItem) {
	// Call rest api with a GET here!
}

func GetTodoItemById(id int64) TodoItem {
	var item TodoItem
	return item
}

func AddTodoItem(title string, description string, isdone bool) TodoItem {
	item := TodoItem{Title: title, Description: description, IsDone: isdone, CreatedAt: time.Now()}
	return item
}

func RemoveTodoItem(id int64) bool {
	return false
}

func UpdateTodoItem(update TodoItem) TodoItem {
	var current TodoItem

	UpdateTodoItemDetails(&current, &update)

	return current
}

func UpdateTodoItemDetails(existingItem *TodoItem, updatedItem *TodoItem) {
	// // Update the info
	if updatedItem.Title != existingItem.Title {
		existingItem.Title = updatedItem.Title
	}

	if updatedItem.Description != existingItem.Description {
		existingItem.Description = updatedItem.Description
	}

	if updatedItem.DueAt != existingItem.DueAt {
		existingItem.DueAt = updatedItem.DueAt
	}

	if updatedItem.IsDone != existingItem.IsDone {
		existingItem.IsDone = updatedItem.IsDone
	}

	existingItem.UpdatedAt = time.Now()
}

func (t TodoItem) String() string {
	return fmt.Sprintf("%s  - %t", t.Description, t.IsDone)
}
