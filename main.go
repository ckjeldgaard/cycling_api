package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type rider struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationaliy"`
	Age         int    `json:"age"`
	Team        string `json:"team"`
}

func main() {
	router := gin.Default()
	router.GET("/riders", getRiders)
	router.GET("/riders/:id", getRiderByID)
	router.POST("/riders", addRider)

	router.Run("localhost:8080")
}

// riders slice to seed rider data.
var riders = []rider{
	{ID: "1", Name: "Wout van Aert", Nationality: "Belgium", Age: 27, Team: "Jumbo-Visma"},
	{ID: "2", Name: "Mathieu van der Poel", Nationality: "Netherlands", Age: 27, Team: "Alpecin-Fenix"},
	{ID: "3", Name: "Julian Alaphilippe", Nationality: "France", Age: 29, Team: "Quick-Step Alpha Vinyl Team"},
	{ID: "4", Name: "Kasper Asgreen", Nationality: "Denmark", Age: 27, Team: "Quick-Step Alpha Vinyl Team"},
	{ID: "5", Name: "Tadej Pogačar", Nationality: "Slovenia", Age: 23, Team: "UAE Team Emirates"},
}

// getRiders responds with the list of all riders as JSON.
func getRiders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, riders)
}

// addRider adds a new rider from JSON received in the request body.
func addRider(c *gin.Context) {
	var newRider rider

	// Call BindJSON to bind the received JSON to newRider.
	if err := c.BindJSON(&newRider); err != nil {
		return
	}

	// Add the new rider to the slice.
	riders = append(riders, newRider)
	c.IndentedJSON(http.StatusCreated, newRider)
}

// getRiderByID locates the riders whose ID value matches the id
// parameter sent by the client, then returns that rider as a response.
func getRiderByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of riders, looking for
	// a rider whose ID value matches the parameter.
	for _, a := range riders {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rider not found"})
}
