package models

import "time"

type TodoItem struct {
	Id          int64     `json:id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueAt       string    `json:"date_due"`
	CreatedAt   time.Time `json:"date_added"`
	UpdatedAt   time.Time `json:"date_updated"`
	IsDone      bool      `j́son:"isdone"`
}

func LoadTodoItemList(todoItemList *[]TodoItem) {
	// Call rest api with a GET here!
}

func GetTodoItemById(id int64) TodoItem {
	var item TodoItem
	return item
}

func AddTodoItem() TodoItem {
	var item TodoItem
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
