package controllers

import (
	"encoding/json"
	"fmt"
	"go-db-gorm/pkg/models"
	"go-db-gorm/pkg/utils"
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

		// This or that ^, not both since you will respond with everything twince if you do
		// c.JSON(200, gin.H{
		// 	"items": newItems,
		// })

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
	}

	itemDetails := models.GetTodoItemById(ID)
	// res, _ := json.Marshal(itemDetails) // No need to Marshal the result, Gin handles this.

	c.JSON(200, gin.H{
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

	// res, _ := json.Marshal(newTodoItem)
	c.JSON(200, gin.H{
		"items": newTodoItem,
	})
}

func UpdateItem(c *gin.Context) {
	var updatedItem = &models.TodoItem{}

	itemId := c.Param("itemId")
	ID, err := strconv.ParseInt(itemId, 0, 0)

	if err != nil {
		fmt.Println(("Error parsing data."))
	}

	utils.ParseBody(c.Request, &updatedItem)
	itemDetails := models.UpdateTodoItem(ID, updatedItem)

	res, _ := json.Marshal(itemDetails)
	c.JSON(200, gin.H{
		"items": res,
	})
}

func DeleteItem(c *gin.Context) {
	// itemId := c.Param("itemId")
	// vars := mux.Vars(r)
	// itemId := vars["Id"]
	// ID, err := strconv.ParseInt(itemId, 0, 0)

	// if err != nil {
	// 	fmt.Println(("Error parsing data."))
	// }

	// deletedItem := models.DeleteItem(ID)

	// res, _ := json.Marshal(deletedItem)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader((http.StatusOK))
	// w.Write(res)
}
