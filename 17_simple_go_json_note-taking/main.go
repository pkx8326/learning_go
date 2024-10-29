package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/simple_go_jason_note-taking/note"
)

func main() {
	var title, content string
	var newnote note.Note
	title = getNote("Please input a title :")
	content = getNote("Please input the note text :")
	newnote = note.NewNote(title, content)
	newnote.DisplayNote()
	err := newnote.SaveNote()
	if err != nil {
		fmt.Println("Error! Can't save json file.")
	} else {
		fmt.Println("Note saved successfully.")
	}
}

/////////////

func getNote(prompt string) (note string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		note, _ = reader.ReadString('\n') // This is a single byte character in Go or 'rune'
		note = strings.TrimSpace(note)
		if note == "" {
			fmt.Println("The text cannot be blank. Please try again.")
			continue
		} else {
			return
		}
	}
}
