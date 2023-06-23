package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	qmgo "github.com/qiniu/qmgo"
)

var database *qmgo.Database
var collection *qmgo.Collection

func main() {

	// create new Client
	const databaseURI = "mongodb://localhost:27017"
	fmt.Println("Connecting to database", databaseURI)
	ctx := context.Background()
	connection, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: databaseURI})

	database = connection.Database("test")    // creating Database connection
	collection = database.Collection("books") // get the collection
	defer func() {
		if err = connection.Close(ctx); err != nil {
			fmt.Println("Closing Connection to database", databaseURI)
			panic(err)
		}
	}()

	router := gin.Default() // create router using gin

	// register routes
	router.POST("/books", CreateBook)
	router.GET("/books", GetBooks)
	router.GET("/books/:bookId", GetBook)
	router.PATCH("/books/:bookId", UpdateBook)
	router.DELETE("/books/:bookId", DeleteBook)

	fmt.Println("Service is up & running at localhost:8000")
	router.Run(":8000") // register router to port 8000
}
