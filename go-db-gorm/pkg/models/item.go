package models

import (
	"fmt"
	"go-db-gorm/pkg/config"
	"log"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type TodoItem struct {
	// gorm.Model              // Do i need this? Look it up!!
	// gorm.Model contains the following:
	// ID        uint           `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"
	// So we can use gorm.Model to get these by default in each entity

	Id          int64     `gorm:"primarykey json:id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueAt       string    `json:"date_due"`
	CreatedAt   time.Time `json:"date_added"`
	UpdatedAt   time.Time `json:"date_updated"`
	IsDone      bool      `j́son:"isdone"`
}

func init() {
	fmt.Println("Initializing entities...")

	config.CreateDBConnection()
	db = config.GetDB()

	err := db.AutoMigrate(&TodoItem{})
	if err != nil {
		log.Fatalf("Failed automigrate(): %s", err)
	}
}

// Using non generic version here to keep it simple

func (item *TodoItem) CreateTodoItem() *TodoItem {
	fmt.Println("Trying CreateTodoItem()")
	db.Create(&item)
	return item
}

func GetAllTodoItems() ([]TodoItem, error) {
	fmt.Println("Trying GetAllTodoItems()")

	var listItems []TodoItem

	if db == nil {
		fmt.Printf("INvalid database reference.")
	}

	res := db.Find(&listItems)

	if res.Error != nil {
		fmt.Printf("DB Error: %s", res.Error)
		return nil, res.Error
	}
	return listItems, nil
}

func GetTodoItemById(Id int64) *TodoItem {
	var item TodoItem
	db.First(&item, Id)
	return &item
}

func DeleteTodoItem(Id int64) TodoItem {
	var item TodoItem
	db.Where("ID = ?", Id).Delete(&item)
	return item
}

func UpdateTodoItem(Id int64, updatedItem *TodoItem) *TodoItem {
	// // Get the item from the DB
	itemDetails := GetTodoItemById(Id)

	if itemDetails != nil {
		UpdateTodoItemDetails(itemDetails, updatedItem)

		db.Save(&itemDetails)

		return itemDetails
	}

	return nil
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
