package views

import (
	"fmt"
	"go-fyne-todo/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowEditWindow(a *fyne.App, w *fyne.Window, todo models.TodoItem) models.TodoItem {
	edit_win := (*a).NewWindow("Edit todo")
	edit_win.Resize((*w).Content().Size())

	var result models.TodoItem
	var isNew bool = true

	if todo.Title != "" {
		isNew = false
		result = todo
	} else {
		isNew = true
		result = models.AddTodoItem("Title 2", "Description 2", false)
	}

	edit_win.SetOnClosed(func() {
		(*w).Show()
	})

	title_entry := widget.NewEntry()
	title_entry.SetPlaceHolder("Enter title")
	save_btn := widget.NewButton("Save", func() {
		if isNew {
			fmt.Println("Save new was clicked!")
		} else {
			fmt.Println("Save update was clicked!")
		}

		edit_win.Close()
	})

	// Uncomment this if the button should be disabled by default
	save_btn.Disable()

	newTodoDesc := widget.NewEntry()
	newTodoDesc.PlaceHolder = "New Todo goes here..."
	newTodoDesc.OnChanged = func(s string) {
		// What happens when the text in the text field changes
		save_btn.Disable()

		if len(s) >= 3 {
			save_btn.Enable()
		}
	}

	buttonRow := container.NewBorder(
		nil, // TOP
		nil, // BOTTOM
		nil, // LEFT
		// RIGHT ↓
		save_btn,
		// take the rest of the space
		nil,
	)

	edit_body := container.NewBorder(
		nil,
		buttonRow,
		nil,
		nil,
		newTodoDesc,
	)

	edit_win.SetContent(edit_body)

	edit_win.Show()
	(*w).Hide()

	return result
}
