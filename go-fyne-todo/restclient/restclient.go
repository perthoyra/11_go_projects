package restclient

import (
	"fmt"
	"go-fyne-todo/config"
	"io"
	"log"
	"net/http"
)

// '../v1/item/{id}'

func Get(config *config.AppConfig, action string) []byte {
	url := fmt.Sprintf("%s/%s/%s", config.Base_Url, config.Version_number, action)
	fmt.Println(url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		// os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	return responseData
}

func GetItemById() {

}
func Post() {

}
func Put() {

}
func Delete() {

}
