package main

import (
	"go-fyne-todo/controllers"
	"go-fyne-todo/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var TodoList []models.TodoItem

func ShowEditWindow(a fyne.App, w fyne.Window) fyne.Window {
	edit_win := a.NewWindow("Edit todo")

	title_entry := widget.NewEntry()
	title_entry.SetPlaceHolder("Enter title")

	submitButton := widget.NewButton("Submit", func() {
		name := title_entry.Text
		dialog.ShowInformation("Hello", "Hello, "+name+"!", edit_win)
		edit_win.Hide()
		w.Show()
	})

	edit_win.SetContent(container.NewVBox(
		title_entry,
		submitButton,
	))

	edit_win.Show()
	w.Hide()
}

func main() {
	a := app.New()
	main_window := a.NewWindow("TODO App")

	controllers.LoadTodoItemList(&TodoList)

	main_window.SetContent(widget.NewLabel("TODOs will go here"))
	main_window.Show()

	// ShowEditWindow(a, main_window)

	a.Run()
}
