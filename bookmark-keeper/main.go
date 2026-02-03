package main

import "fmt"

func main() {
	bookmarks := map[string]string{}
programLoop:
	for {
		displayMenu()
		choice := getResult()

		switch choice {
		case 1:
			viewBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break programLoop
		}
	}

}

func displayMenu() {
	fmt.Println("- 1. View Bookmarks")
	fmt.Println("- 2. Add Bookmark")
	fmt.Println("- 3. Delete Bookmark")
	fmt.Println("- 4. Exit")
	fmt.Print("Enter choice: ")
}

func getResult() int {
	var menuChoice int
	fmt.Scan(&menuChoice)
	return menuChoice
}

func viewBookmarks(bookmarks map[string]string) {
	for k, v := range bookmarks {
		fmt.Println(k, ":", v)
	}
}

func addBookmark(bookmarks map[string]string) {
	var name, address string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter address: ")
	fmt.Scan(&address)

	bookmarks[name] = address
}

func deleteBookmark(bookmarks map[string]string) {
	var name string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	delete(bookmarks, name)
}
