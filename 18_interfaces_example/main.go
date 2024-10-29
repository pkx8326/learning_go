package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/simple_go_jason_note-taking/note"
	"example.com/simple_go_jason_note-taking/todo"
)

// type saver interface {
// 	save()
// }

func main() {
	//////Variable declarations
	var title, content, todotext string
	var newnote note.Note
	var newtodo todo.Todo

	//////NOTE
	title = getNote("Please input a title :")
	content = getNote("Please input the note text :")
	newnote = note.NewNote(title, content)
	newnote.DisplayNote()
	err := newnote.SaveNote()
	if err != nil {
		fmt.Println("Error! Can't save note to json file.")
	} else {
		fmt.Println("Note saved successfully.")
	}

	//////TODO
	todotext = getTodo("Please insert a todo text :")
	newtodo = todo.NewToDo(todotext)
	newtodo.DisplayTodo()
	err = newtodo.SaveTodo()
	if err != nil {
		fmt.Println("Error! Can't save todo to json file.")
	} else {
		fmt.Println("Todo saved successfully.")
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

func getTodo(prompt string) (todo string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		todo, _ = reader.ReadString('\n') // This is a single byte character in Go or 'rune'
		todo = strings.TrimSpace(todo)
		if todo == "" {
			fmt.Println("The text cannot be blank. Please try again.")
			continue
		} else {
			return
		}
	}
}
