package main

import (
	"go-fyne-todo/models"
	"go-fyne-todo/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var TodoList []models.TodoItem

func main() {
	app := app.New()
	main_window := app.NewWindow("TODO App")
	main_window.Resize(fyne.NewSize(300, 400))

	TodoList = make([]models.TodoItem, 0)

	// controllers.LoadTodoItemList(&TodoList)
	views.CreateMainWindow(&app, &main_window, &TodoList)

	main_window.Show()

	app.Run()
}
