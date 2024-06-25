package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"log"
		"html/template"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Book One", Author: "John Doe"},
	{ID: "2", Title: "Book Two", Author: "Jane Doe"},
}

func main() {
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: os.Stdout}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.GET("/books", getBooks)
	r.GET("/", getBooks)
	r.GET("/books/:id", getBookByID)
	r.POST("/books", createBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	r.Run(":8080")

}


func getBooks(c *gin.Context) {
	// t, err := template.New("books").ParseFiles("templates/books.html")
	t, err := template.ParseFiles("templates/books.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "template error"})
		return
	}

	log.Printf("Executing template with books: %+v\n", books)
	err = t.Execute(c.Writer, books)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "template execution error"})
		return
	}
}


func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			if err := c.BindJSON(&book); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			books[i] = book
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}


// run:
//  go run A4_gin_website.go
// or build and then run
//  go build -o A4_gin_website A4_gin_website.go
//  ./A4_gin_website


// FOR AIR and auto reloading

// First run 
// go mod init waterproject
// go mod tidy
// go build -o A4_gin_website A4_gin_website.go


// To enable automatic reloading for a Go web server, you typically use an external tool like `air` or `fresh`. The `go run --reload` command is not a built-in feature of Go.

// Here is how you can use `air` for automatic reloading:

// 1. Install `air`:
//     ```sh
//     go install github.com/air-verse/air@latest
//     ```

// 2. Create an `air.toml` configuration file (if needed) in your project directory:
//     ```toml
//     # Config file for air
//     [build]
//       cmd = "go build -o ./tmp/main ."
//       bin = "./tmp/main"
//       full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
//       log = "build.log"
//       includes = ["."]
//       excludes = ["tmp", "vendor", "air"]
//       delay = 1000
//       kill_delay = 500
//       rebuild_on_error = false
//     ```

// 3. Run your application with `air`:
//     ```sh
//     air
//     ```
