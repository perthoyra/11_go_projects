package main

import (
	"go-fyne-todo/models"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var TodoList []models.TodoItem

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")

	models.LoadTodoItemList(&TodoList)

	w.SetContent(widget.NewLabel("TODOs will go here"))
	w.Show()
	a.Run()
}
