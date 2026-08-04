package main

import (
	"fmt"
	"go-fyne-todo/controllers"
	"go-fyne-todo/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var TodoList []models.TodoItem

func main() {
	app := app.New()
	main_window := app.NewWindow("TODO App")
	main_window.Resize(fyne.NewSize(300, 400))
	main_window.SetMaster()

	TodoList = make([]models.TodoItem, 0)

	controllers.LoadTodoItemList(&TodoList)

	fmt.Printf(TodoList[0].Description)

	// views.CreateMainWindow(&app, &main_window, &TodoList)

	main_window.Show()

	app.Run()
}
