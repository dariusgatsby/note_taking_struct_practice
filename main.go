package main

import (
	"fmt"

	"example.com/notes/notes"
)

var n notes.Note

func main() {
	fmt.Println("Welcome to your personal note taking app!")
	for {
		fmt.Println("1. Create Note 2. Display Notes")
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {

			finalNote := notes.Note.CreateNote(n)

			err := notes.Note.SaveToJSON(finalNote)

			if err != nil {
				fmt.Println(err)
				return
			}

		}
		if choice == 2 {
			// _, err := notes.ViewNotes()

			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			notes.ViewNotes()
		}
		if choice == 3 {
			fmt.Println("Bye have a wonderful time!")
			break
		}
	}
}

// func getUserData(promptString string) string {
// 	fmt.Print(promptString)
// 	var value string
// 	var title string
// 	fmt.Scanln(&value)
// 	fmt.Scanln(&title)
// 	return value, title
// }
