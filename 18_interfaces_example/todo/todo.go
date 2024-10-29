package todo

import (
	"encoding/json" //for saving json files (used with json.Marshal())
	"fmt"
	"os"
)

type Todo struct {
	TodoContent string `json:"todo_content"`
}

func NewToDo(todotext string) (newtodo Todo) {
	newtodo = Todo{
		TodoContent: todotext,
	}
	return
}

func (todo Todo) SaveTodo() (err error) {
	filename := "todo.json"
	json, err := json.Marshal(todo) //The .Marshal() method works only if fields in the struct are publicly accessible
	if err != nil {
		return
	} else {
		fmt.Println("Saving the todo text to json...")
		return os.WriteFile(filename, json, 0644)
	}
}

func (todo Todo) DisplayTodo() { //Create a method for the todo struct.
	fmt.Printf("Todo:\n %v \n", todo.TodoContent)
} // If we don't expose the fields in the Note struct, we can use this method to show them in the main program.
