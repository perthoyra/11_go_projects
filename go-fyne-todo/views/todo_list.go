package views

import (
	"go-fyne-todo/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func showTodoList(data binding.List[any]) *widget.List {
	todo_list := widget.NewListWithData(
		// the binding list with the data
		data,
		// func that returns the component structure of the List Item
		// this is the template all items will follow
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil, nil, nil,
				// left of the border
				widget.NewCheck("", func(b bool) {}),
				// takes the rest of the space
				widget.NewLabel(""),
			)
		},
		// func that is called for each item in the list and allows
		// you to show the content on the previously defined ui structure
		func(di binding.DataItem, o fyne.CanvasObject) {
			ctr, _ := o.(*fyne.Container)
			// ideally we should check `ok` for each one of those casting
			// but we know that they are those types for sure
			l := ctr.Objects[0].(*widget.Label)
			c := ctr.Objects[1].(*widget.Check)

			// diu, _ := di.(binding.Untyped).Get()
			// todo := diu.(models.TodoItem)
			todo := models.NewTodoFromDataItem(di)

			l.SetText(todo.Description)
			c.SetChecked(todo.IsDone)
		},
	)

	return todo_list
}
