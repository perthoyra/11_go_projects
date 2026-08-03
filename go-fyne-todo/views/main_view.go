package views

import (
	"fmt"
	"go-fyne-todo/controllers"
	"go-fyne-todo/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func CreateMainWindow(app *fyne.App, main_window *fyne.Window, data *[]models.TodoItem) {

	// create the user controls we want in the window
	t := controllers.AddTodoItem("Remember me", "Show this on the window", false)
	(*data) = append((*data), t)

	// create a data binding for the todo list
	todolist := binding.NewUntypedList()
	for _, t := range *data {
		todolist.Append(t)
	}

	fmt.Printf("Size of todolist: %v\n", todolist.Length())

	add_btn := widget.NewButton("Add", func() {
		newItem := controllers.CreateNewItem()
		todolist.Append(ShowEditWindow(app, main_window, (*newItem)))
		fmt.Println("Add was clicked!")
	})
	// Uncomment this if the button should be disabled by default
	// add_btn.Disable()

	// newTodoDesc := widget.NewEntry()
	// newTodoDesc.PlaceHolder = "New Todo goes here..."
	// newTodoDesc.OnChanged = func(s string) {
	// 	// What happens when the text in the text field changes
	// 	add_btn.Disable()

	// 	if len(s) >= 3 {
	// 		add_btn.Enable()
	// 	}
	// }

	buttonRow := container.NewBorder(
		nil, // TOP
		nil, // BOTTOM
		nil, // LEFT
		// RIGHT ↓
		add_btn,
		// take the rest of the space
		nil,
	)

	main_body := container.NewBorder(
		nil,
		buttonRow,
		nil,
		nil,
		showTodoList(todolist),
	)

	(*main_window).SetContent(main_body)
}
