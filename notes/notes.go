package notes

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	NoteName    string    `json:"title"`
	NoteContent string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}

func (n Note) CreateNote() Note {
	scanner := bufio.NewScanner(os.Stdin)

	var title string
	var content string

	fmt.Println("Enter title of your note: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Enter note contents: ")
	scanner.Scan()
	content = scanner.Text()

	finalNote := Note{
		NoteName:    title,
		NoteContent: content,
		CreatedAt:   time.Now(),
	}

	return finalNote
}

func (note Note) SaveToJSON() error {

	noteInJSON, err := json.Marshal(note)

	if err != nil {
		return errors.New("error converting to JSON")
	}
	fmt.Println("Successfully conversion to JSON")

	// noteTitle := strings.ToLower(note.NoteName)
	// noteTitle = strings.ReplaceAll(note.NoteName, " ", "_")

	// file, err1 := os.OpenFile("notes.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// if  err1 != nil {
	// 	return err
	// }

	_ = os.WriteFile("notes.json", noteInJSON, 0644)

	return nil
}

func ViewNotes() (Note, error) {

	file, err := os.ReadFile("notes.json")
	if err != nil {
		return Note{}, errors.New("error reading from file")
	}
	var notes Note
	json.Unmarshal(file, &notes)

	fmt.Println("Title: " + notes.NoteName)
	fmt.Println("Content: " + notes.NoteContent)
	return notes, nil
}

func (n Note) Display() error {
	noteTitle := strings.ReplaceAll(n.NoteName, " ", "_")
	noteTitle = strings.ToLower(noteTitle)
	file, err := os.ReadFile(noteTitle + ".json")
	if err != nil {
		return errors.New("error reading from file")
	}
	var notes Note
	json.Unmarshal(file, &notes)
	fmt.Println("Title: " + n.NoteName)
	fmt.Println("Content: " + n.NoteContent)
	return nil
}
