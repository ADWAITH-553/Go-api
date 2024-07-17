package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Merchant of Venice", Author: "William Shakespeare", Quantity: 4},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:9768")
}
