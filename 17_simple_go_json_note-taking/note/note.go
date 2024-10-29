package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title string, content string) (newnote Note) {
	newnote = Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return
}

func (note Note) SaveNote() (err error) {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(note) //The .Marshal() method works only if fields in the struct are publicly accessible
	if err != nil {
		return
	} else {
		fmt.Println("Saving the note to json...")
		return os.WriteFile(filename, json, 0644)
	}
}

func (note Note) DisplayNote() { //Create a method for the note struct.
	fmt.Printf("Your note titled \"%v\" contains the following message:\n %v", note.Title, note.Content)
} // If we don't expose the fields in the Note struct, we can use this method to show them in the main program.
