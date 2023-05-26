package main

import (
	"log"
	"net/http"

	"01-Login/platform/authenticator"
	"01-Login/platform/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type NameList struct {
	NameF string  `json:"fName"`
	NameL string  `json:"lName"`
	Age   float32 `json:"Age"`
}

var names = []NameList{
	{NameF: "Quandale", NameL: "Dingle the third", Age: 59},
	{NameF: "James", NameL: "Howlet", Age: 134},
	{NameF: "Bihan", NameL: "Tundra", Age: 38},
}

func getNames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, names)
}

func handlerFunc(c *gin.Context) {
	c.String(200, "My record of Kira's kills")
}

func deleteNames(c *gin.Context) {
	var deleteItem NameList
	if err := c.ShouldBindJSON(&deleteItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var remove bool
	for i, d := range names {
		if d.NameF == deleteItem.NameF && d.NameL == deleteItem.NameL && d.Age == deleteItem.Age {
			names = append(names[:i], names[i+1:]...)
			remove = true
			break
		}
	}

	if remove {
		c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
	}
}

func putNames(c *gin.Context) {
	var update NameList
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	names = append(names, update)
	c.IndentedJSON(http.StatusCreated, update)
}

func postNames(c *gin.Context) {
	var newName NameList
	if err := c.BindJSON(&newName); err != nil {
		return
	}

	names = append(names, newName)
	c.IndentedJSON(http.StatusCreated, newName)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	r := gin.Default()

	// Routes
	r.GET("/", handlerFunc)
	r.GET("/names", getNames)
	r.PUT("/names", putNames)
	r.POST("/names", postNames)
	r.DELETE("/names", deleteNames)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", r); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
