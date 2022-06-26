package main

import (
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Create some test data
var books = []Book{
	{ID: "1", Title: "On Writing Well", Author: "William Zinsser"},
	{ID: "2", Title: "Stein on Writing", Author: "Sol Stein"},
	{ID: "3", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"},
}

func GetBooks(c *gin.Context) {
	c.JSON(200, gin.H{
		"books": books,
	})
}

func GetBook(c *gin.Context) {
	bookId := c.Query("id")
	for _, item := range books {
		// Check to see if a book matches the ID passed in as a parameter
		if item.ID == bookId {
			c.JSON(200, gin.H{
				"id":     item.ID,
				"title":  item.Title,
				"author": item.Author,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "Book not found",
	})
}

func CreateBook(c *gin.Context) {
	bookId := c.Query("id")
	bookTitle := c.Query("title")
	bookAuthor := c.Query("author")

	if len(bookId) == 0 || len(bookAuthor) == 0 {
		c.JSON(400, gin.H{
			"id":     bookId,
			"title":  bookTitle,
			"author": bookAuthor,
		})
		return
	}

	// Check to see if a book with this ID already exists
	for _, item := range books {
		if item.ID == bookId {
			// 409 means there's a conflict
			c.JSON(409, gin.H{
				"message": "Book already exists",
			})
			return
		}
	}

	// Add the book to the list
	books = append(books, Book{
		ID:     bookId,
		Title:  bookTitle,
		Author: bookAuthor,
	})

	// 201 means created
	c.JSON(201, gin.H{
		"id":     bookId,
		"title":  bookTitle,
		"author": bookAuthor,
	})
}

func UpdateBook(c *gin.Context) {
	bookId := c.Query("id")
	bookTitle := c.Query("title")
	bookAuthor := c.Query("author")

	for index, item := range books {
		if item.ID == bookId {
			// Delete the element at index and preserve the order of the books slice.
			// This approach creates two slices from the original, books[:index] and books[i+index:]
			// and then joins them back together into a single slice.
			// The element at index is not included.
			books = append(books[:index], books[index+1:]...)

			// An alternative approach if preserving order is not necessary:
			// First, copy the last element to index.
			// books[index] = books[len(books)-1]
			// Then, remove the last element from the slice by truncating it.
			// books = books[:len(books)-1]

			// Add the book to the list
			books = append(books, Book{
				ID:     bookId,
				Title:  bookTitle,
				Author: bookAuthor,
			})

			c.JSON(200, gin.H{
				"id":     bookId,
				"title":  bookTitle,
				"author": bookAuthor,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Book not found",
	})
}

func DeleteBook(c *gin.Context) {
	bookId := c.Query("id")

	for index, item := range books {
		if item.ID == bookId {
			books = append(books[:index], books[index+1:]...)
			c.JSON(200, gin.H{
				"id": bookId,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error": "Book not found",
	})
}

func main() {
	// Create a Gin router with default middleware
	r := gin.Default()

	// Register routes
	r.GET("/books", GetBooks)
	r.GET("/book", GetBook)
	r.POST("/book", CreateBook)
	r.PUT("/book", UpdateBook)
	r.DELETE("/book", DeleteBook)

	// Listen and serve on localhost:8080
	r.Run()
}
