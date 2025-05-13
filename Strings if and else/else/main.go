package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

func main() {
	var books []Book
	var n int

	fmt.Print("How many books do you want to enter? ")
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		var book Book
		fmt.Printf("\nEntering details for Book #%d\n", i+1)

		fmt.Print("Enter Book ID: ")
		fmt.Scan(&book.ID)

		fmt.Print("Enter Book Title: ")
		book.Title, _ = reader.ReadString('\n')
		book.Title = strings.TrimSpace(book.Title)

		fmt.Print("Enter Author Name: ")
		book.Author, _ = reader.ReadString('\n')
		book.Author = strings.TrimSpace(book.Author)

		fmt.Print("Enter Genre: ")
		book.Genre, _ = reader.ReadString('\n')
		book.Genre = strings.TrimSpace(book.Genre)

		books = append(books, book)
	}

	fmt.Println("Library Books")
	for _, b := range books {
		fmt.Printf("Book ID: %d\n", b.ID)
		fmt.Printf("Title: %s\n", b.Title)
		fmt.Printf("Author: %s\n", b.Author)
		fmt.Printf("Genre: %s\n", b.Genre)
	}

	// Search books by genre
	fmt.Print("\nEnter a genre to search for books: ")
	searchGenre, _ := reader.ReadString('\n')
	searchGenre = strings.TrimSpace(strings.ToLower(searchGenre))

	fmt.Printf("\nBooks in genre '%s':\n", searchGenre)
	found := false
	for _, b := range books {
		if strings.ToLower(b.Genre) == searchGenre {
			fmt.Printf("- %s by %s (ID: %d)\n", b.Title, b.Author, b.ID)
			found = true
		}
	}

	if !found {
		fmt.Println("No books found for this genre.")
	}
}
