package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	openapiclient "github.com/jom6768/go-oda-tmf632/api"
)

// createIndividual
func createIndividual(c *gin.Context) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	resp, _, err := apiClient.IndividualAPI.CreateIndividual(context.Background()).Execute()
	log.Println("resp:", resp)
	log.Println("err:", err)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to insert customer"})
		return
	}

	c.JSON(http.StatusCreated, "")
}

// listIndividual
func listIndividual(c *gin.Context) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	resp, _, err := apiClient.IndividualAPI.ListIndividual(context.Background()).Execute()
	log.Println("resp:", resp)
	log.Println("err:", err)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func main() {
	r := gin.Default()
	r.POST("/tmf632/customer", createIndividual)
	r.GET("/tmf632/customer", listIndividual)
	r.Run(":8082")
}
