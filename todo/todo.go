package todo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	TodoContent string `json:"todo"`
}

func CreateTodo() Todo {
	scanner := bufio.NewScanner(os.Stdin)

	var content string

	fmt.Println("Enter todo contents: ")
	scanner.Scan()
	content = scanner.Text()

	finalTodo := Todo{
		TodoContent: content,
	}

	return finalTodo
}

func (todo Todo) SaveToJSON() error {

	todoInJSON, err := json.Marshal(todo)

	if err != nil {
		return errors.New("error converting to JSON")
	}
	fmt.Println("Successfully conversion to JSON")

	_ = os.WriteFile("todos.json", todoInJSON, 0644)

	return nil
}

func View() (Todo, error) {

	file, err := os.ReadFile("todos.json")
	if err != nil {
		return Todo{}, errors.New("error reading from file")
	}
	var todo Todo
	json.Unmarshal(file, &todo)

	fmt.Println("Todo: " + todo.TodoContent)
	return todo, nil
}
