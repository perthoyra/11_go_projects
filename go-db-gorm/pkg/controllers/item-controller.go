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

func GetAllTodoItems(c *gin.Context) {
	newItems := models.GetAllTodoItems()
	res, _ := json.Marshal(newItems)

	c.IndentedJSON(http.StatusOK, res)
	c.JSON(200, gin.H{
		"items": res,
	})
}

func GetTodoItemById(c *gin.Context) {
	itemId := c.Query("id")
	ID, err := strconv.ParseInt(itemId, 0, 0)

	if err != nil {
		fmt.Println(("Error parsing data."))
	}

	itemDetails := models.GetTodoItemById(ID)
	res, _ := json.Marshal(itemDetails)

	c.JSON(200, gin.H{
		"items": res,
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

	res, _ := json.Marshal(newTodoItem)
	c.JSON(200, gin.H{
		"items": res,
	})
}

func UpdateItem(c *gin.Context) {
	var updatedItem = &models.TodoItem{}

	itemId := c.Query("id")
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
