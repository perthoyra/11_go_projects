package controllers

import (
	"encoding/json"
	"fmt"
	"go-db-gorm/pkg/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var NewItem models.TodoItem

func init() {
	fmt.Println("Configuring context...")
}

func GetAllTodoItems(c *gin.Context) {
	var newItems []models.TodoItem

	newItems, err := models.GetAllTodoItems()

	if err != nil {
		c.IndentedJSON(http.StatusNoContent, "No rows in result")
		return
	}

	if newItems != nil {
		c.IndentedJSON(http.StatusOK, newItems)
	}
}

func GetTodoItemById(c *gin.Context) {
	itemId := c.Param("itemId")
	ID, err := strconv.ParseInt(itemId, 0, 0)

	if err != nil {
		fmt.Println(("Error parsing data."))

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err,
		})
	}

	itemDetails := models.GetTodoItemById(ID)

	c.JSON(http.StatusOK, gin.H{
		"items": itemDetails,
	})
}

func CreateItem(c *gin.Context) {
	itemCreator := &models.TodoItem{}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("Error reading body content: %v", err) // Maybe change this to return a http error code
	}

	newTodoItem := itemCreator.CreateTodoItem()

	jsonErr := json.Unmarshal(jsonData, &newTodoItem)
	if jsonErr != nil {
		log.Fatalf("Error unmarshalling JSON: %v", jsonErr) // Maybe change this to return a http error code
	}

	itemDetails := models.UpdateTodoItem(newTodoItem.Id, newTodoItem)

	c.JSON(http.StatusOK, gin.H{
		"items": itemDetails,
	})
}

func UpdateItem(c *gin.Context) {

	// Read values from the request
	itemId := c.Param("itemId")
	ID, err := strconv.ParseInt(itemId, 0, 0)

	if err != nil {
		fmt.Println(("Error parsing data."))
	}

	// Get the expected json data
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("Error reading body content: %v", err) // Maybe change this to return a http error code
	}

	// find the existing item
	updatedItem := models.GetTodoItemById(ID)

	// Update the item
	jsonErr := json.Unmarshal(jsonData, &updatedItem)
	if jsonErr != nil {
		log.Fatalf("Error unmarshalling JSON: %v", jsonErr) // Maybe change this to return a http error code
	}

	// Persist the changes to the DB
	itemDetails := models.UpdateTodoItem(ID, updatedItem)

	c.JSON(http.StatusOK, gin.H{
		"items": itemDetails,
	})
}

func DeleteItem(c *gin.Context) {

	// Read values from the request
	itemId := c.Param("itemId")
	ID, err := strconv.ParseInt(itemId, 0, 0)

	if err != nil {
		fmt.Println(("Error parsing data."))
	}

	deletedItem := models.DeleteTodoItem(ID)

	// res, _ := json.Marshal(deletedItem)

	c.JSON(http.StatusOK, gin.H{
		"items": deletedItem,
	})
}
