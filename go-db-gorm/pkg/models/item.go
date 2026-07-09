package models

import (
	"go-db-gorm/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type TodoItem struct {
	gorm.Model
	Id          int64  `gorm:""json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date_due    string `json:"date_due"`
	Date_added  string `json:"date_added"`
	IsDone      bool   `j́son:"isdone"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&TodoItem{})
}

func (item *TodoItem) CreateItem() *TodoItem {
	db.NewRecord(item)
	db.Create(&item)
	return item
}

func GetAllItems() []TodoItem {
	var Items []TodoItem
	db.Find(&Items)
	return Items
}

func GetItemById(Id int64) (*TodoItem, *gorm.DB) {
	var getItem TodoItem
	db := db.Where("ID=?", Id).Find(&getItem)
	return &getItem, db
}

func DeleteItem(Id int64) TodoItem {
	var item TodoItem
	db.Where("ID=?", Id).Delete(item)
	return item
}
