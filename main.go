package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	openapiclient "github.com/jom6768/go-oda-tmf632/api"
)

// listIndividual
func listIndividual(c *gin.Context) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	resp, _, err := apiClient.IndividualAPI.ListIndividual(context.Background()).Execute()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func main() {
	r := gin.Default()
	r.GET("/tmf632/customer", listIndividual)
	r.Run(":8082")
}
