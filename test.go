package main

import (
	"context"
	"fmt"

	"time"

	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/utahta/go-linenotify"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ShirtColor struct {
	ID    string `json:"id"`
	Day   string `json:"day"`
	Color string `json:"color"`
}

type Message struct {
	Message string `json:"message"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

var luckycolor = []ShirtColor{
	{ID: "1", Day: "Sunday", Color: "Red"},
	{ID: "2", Day: "Monday", Color: "Yellow"},
	{ID: "3", Day: "Tuesday", Color: "Pink"},
	{ID: "4", Day: "Wednesday", Color: "Green"},
	{ID: "5", Day: "Thursday", Color: "Orange"},
	{ID: "6", Day: "Friday", Color: "Blue"},
	{ID: "7", Day: "Saturday", Color: "Purple"},
}

var linetoken = "rSGlM4NJsgjoOe9gp4iWQ2ytyrLNilQjkvXOObzinAm"

func getbookHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		books,
	)
}

func postbookHandler(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
	}
	books = append(books, book)
	c.JSON(
		http.StatusCreated,
		book,
	)
}

func deletebookHandler(c *gin.Context) {
	id := c.Param("id")
	for i, a := range books {
		if a.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	c.Status(http.StatusNoContent)
}

func getcolor(c *gin.Context) {
	dt := time.Now().Weekday()
	color := "White"
	switch dt.String() {
	case luckycolor[0].Day:
		color = luckycolor[0].Color
		fmt.Println(color)
	case luckycolor[1].Day:
		color = luckycolor[1].Color
		fmt.Println(color)
	case luckycolor[2].Day:
		color = luckycolor[2].Color
		fmt.Println(color)
	case luckycolor[3].Day:
		color = luckycolor[3].Color
		fmt.Println(color)
	case luckycolor[4].Day:
		color = luckycolor[4].Color
		fmt.Println(color)
	case luckycolor[5].Day:
		color = luckycolor[5].Color
		fmt.Println(color)
	case luckycolor[6].Day:
		color = luckycolor[6].Color
		fmt.Println(color)
	}

	c.JSON(
		http.StatusOK,
		dt.String()+" : "+color,
	)
}

func noti(c *gin.Context) {
	dt := time.Now().Weekday()
	color := "White"
	switch dt.String() {
	case luckycolor[0].Day:
		color = luckycolor[0].Color
		fmt.Println(color)
	case luckycolor[1].Day:
		color = luckycolor[1].Color
		fmt.Println(color)
	case luckycolor[2].Day:
		color = luckycolor[2].Color
		fmt.Println(color)
	case luckycolor[3].Day:
		color = luckycolor[3].Color
		fmt.Println(color)
	case luckycolor[4].Day:
		color = luckycolor[4].Color
		fmt.Println(color)
	case luckycolor[5].Day:
		color = luckycolor[5].Color
		fmt.Println(color)
	case luckycolor[6].Day:
		color = luckycolor[6].Color
		fmt.Println(color)
	}

	cli := linenotify.NewClient()
	cli.Notify(context.Background(), linetoken, dt.String()+" : "+color, "", "", nil)
}

func main() {

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{"message": "Hello World!"},
		)
	})

	r.GET("/books", getbookHandler)

	r.POST("/books", postbookHandler)

	r.DELETE("/books/:id", deletebookHandler)

	r.GET("/color", getcolor)

	r.POST("/noti", noti)

	r.Run()
}
