package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetLogFile(c *gin.Context) {
	file, err := os.Open("logfile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	jsonStirng := []map[string]interface{}{}

	for scanner.Scan() {
		var data map[string]interface{}

		// Unmarshal the JSON string into the data variable
		err := json.Unmarshal([]byte(scanner.Text()), &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
		jsonStirng = append(jsonStirng, data)
	}

	// Convert the joined string to a []byte

	c.JSON(http.StatusOK, gin.H{
		"data": jsonStirng,
	})

}
