package main

import (
	"fmt"

	"example.com/notes/notes"
	"example.com/notes/todo"
)

type saver interface {
	SaveToJSON() error
}

func DataSaver(data saver) error {
	err := data.SaveToJSON()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Welcome to your personal note taking app!")
	for {
		fmt.Println("1. Create Note 2. Display Notes")
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {

			finalNote := notes.New()

			err := DataSaver(finalNote)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
		if choice == 2 {
			notes.View()
		}
		if choice == 3 {
			todo := todo.CreateTodo()
			err := saver.SaveToJSON(todo)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		if choice == 4 {
			todo.View()
		}
		if choice == 5 {
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
