package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Todo struct {
	Description  string
	Status       string
	Created_at   time.Time
	Completed_at time.Time
}

var Todos []Todo

func (t Todo) createNewTodo() {
	Todos = append(Todos, t)
}

func main() {

	var todo1 Todo = Todo{Description: "Task1", Status: "in_progress"}
	todo1.createNewTodo()

	todo2 := Todo{Description: "Task2", Status: "completed"}
	todo2.createNewTodo()

	GetTodos()
	StoreIntoJSON(Todos)

}

func GetTodos() {
	fmt.Println("All the todos \n ")
	fmt.Println("Description \t Status \t Created At \t Completed At")
	for _, todo := range Todos {
		fmt.Printf("%s %s %v %v \n", todo.Description, todo.Status, todo.Created_at, todo.Completed_at)
	}
}

func StoreIntoJSON(todos []Todo) {
	// marshal
	result, err := json.Marshal(todos)
	if err != nil {
		fmt.Printf("Received error ")
		return

	}
	fmt.Println(string(result))
	// unmarshal
	var newTodos []Todo
	err = json.Unmarshal(result, &newTodos)
	if err != nil {
		fmt.Println("Error while unmarshalling")
	}
	fmt.Println(newTodos)

}
