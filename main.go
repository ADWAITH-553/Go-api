package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"errors"
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
func bookById( c *gin.Context){
	id:=c.Param("id")
	book,err:=getBookbyId(id)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not Found"})
		return 
	}
	c.IndentedJSON(http.StatusOK,book)
}
func getBookbyId(id string)(*book,error){
	for i,b:=range books{
		if b.ID==id{
			return &books[i],nil
		}
	}
	return nil, errors.New("books not found")
}

func checkoutBook(c *gin.Context){
	id,ok:=c.GetQuery("id")
	if ok==false{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"NO BOOK WITH ID PRESENT"})
		return
	}
	book,err:=getBookbyId(id)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not Found"})
	}
	if book.Quantity<=0{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Book not available"})
		return
	}
	book.Quantity-=1
	c.IndentedJSON(http.StatusOK,book)

}


func createBook(c *gin.Context) {
	var newBook book
	if err:= c.BindJSON(&newBook); err!=nil{
		return
	}
	books=append(books,newBook)
	c.IndentedJSON(http.StatusCreated,newBook)
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id",bookById)
	router.POST("/books",createBook)
	router.PATCH("/checkout",checkoutBook)
	router.Run("localhost:9768")
}
