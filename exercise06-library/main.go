package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	books := make(map[int]string)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 3) 

		command := parts[0]

		if command == "ADD" {
			isbn := parts[1]
			title := parts[2]

			var isbnInt int
			fmt.Sscanf(isbn, "%d", &isbnInt)
			books[isbnInt] = title
		} else if command == "REMOVE" {
			var isbn int
			fmt.Sscanf(parts[1], "%d", &isbn)
			delete(books, isbn)
		}
	}

	var bookList []struct {
		ISBN  int
		Title string
	}

	for isbn, title := range books {
		bookList = append(bookList, struct {
			ISBN  int
			Title string
		}{ISBN: isbn, Title: title})
	}

	sort.Slice(bookList, func(i, j int) bool {
		if bookList[i].Title == bookList[j].Title {
			return bookList[i].ISBN < bookList[j].ISBN
		}
		return bookList[i].Title < bookList[j].Title
	})

	for _, book := range bookList {
		fmt.Println(book.ISBN)
	}
}