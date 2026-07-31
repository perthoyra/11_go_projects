package controllers

import "go-fyne-todo/models"

func LoadTodoItemList(todoItemList *[]models.TodoItem) {
	// Call rest api with a GET here!
	models.LoadTodoItemList(todoItemList)
}

func GetTodoItemById(id int64) models.TodoItem {
	var item models.TodoItem
	return item
}

func AddTodoItem() models.TodoItem {
	var item models.TodoItem
	return item
}

func RemoveTodoItem(id int64) bool {
	return false
}

func UpdateTodoItem(update models.TodoItem) models.TodoItem {
	var current models.TodoItem

	models.UpdateTodoItemDetails(&current, &update)

	return current
}
