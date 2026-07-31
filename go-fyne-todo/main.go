package main

import (
	"fmt"
	"go-fyne-todo/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var TodoList []models.TodoItem

func ShowEditWindow(a fyne.App, w fyne.Window) {
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
	main_window.Resize(fyne.NewSize(300, 400))

	// controllers.LoadTodoItemList(&TodoList)

	t := models.AddTodoItem("Remember me", "Show this on the window", false)

	add_btn := widget.NewButton("Add", func() {
		fmt.Println("Add was clicked!")
	})
	add_btn.Disable()

	newTodoDesc := widget.NewEntry()
	newTodoDesc.PlaceHolder = "New Todo goes here..."
	newTodoDesc.OnChanged = func(s string) {
		add_btn.Disable()

		if len(s) >= 3 {
			add_btn.Enable()
		}
	}

	entry := container.NewBorder(
		nil, // TOP
		nil, // BOTTOM
		nil, // Left
		// RIGHT ↓
		add_btn,
		// take the rest of the space
		newTodoDesc,
	)

	main_body := container.NewBorder(
		nil,
		entry,
		nil,
		nil,
		container.NewCenter(
			widget.NewLabel(t.String()),
		),
	)

	main_window.SetContent(main_body)

	main_window.Show()

	a.Run()
}
