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

	c.JSON(http.StatusOK, gin.H{
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

	// Source - https://stackoverflow.com/a/61920252
	// Posted by chash, modified by community. See post 'Timeline' for change history
	// Retrieved 2026-07-14, License - CC BY-SA 4.0

	// jsonData, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// Handle error
	// }

	// json.Unmarshal(jsonData, &updatedItem)

	itemDetails := models.UpdateTodoItem(ID, updatedItem)

	// res, _ := json.Marshal(itemDetails)
	c.JSON(http.StatusOK, gin.H{
		"items": itemDetails,
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
