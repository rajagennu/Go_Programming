package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Description  string
	Status       string
	Created_at   time.Time
	Completed_at time.Time
}

var Todos []Todo

func (t *Todo) createNewTodo() {
	Todos = append(Todos, *t)
}

func (t *Todo) updateTodo(desc string) {
	t.Description = desc
	fmt.Println(t)
}

func (t *Todo) delete() {
	for index, value := range Todos {
		if value.Description == t.Description {
			Todos = append(Todos[0:index], Todos[index+1:]...) // 1,2,3,4,5 => 4 [0:3][4:]
		}
	}

}

func (t *Todo) markAsComplete() {
	// Exercise
}

func main() {

	for {
		GetTodos()
		fmt.Println("Please mention your operation\n1.Create\n2.Delete\n3.Update ")
		scanner := bufio.NewScanner(os.Stdin)
		_ = scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			createTodo()

		case "2":
			updateTodo()

		case "3":
			deleteTodo()

		default:
			fmt.Errorf("Invalid operation ID : %s ", input)
		}
	}

}

func GetTodos() {
	fmt.Println("All the todos \n ")
	fmt.Println("Description \t Status \t Created At \t Completed At")
	for _, todo := range Todos {
		fmt.Printf("%s %s %v %v \n", todo.Description, todo.Status, todo.Created_at, todo.Completed_at)
	}
}

func createTodo() {
	fmt.Println("Please enter todo description ")
	todoInput := bufio.NewScanner(os.Stdin)
	_ = todoInput.Scan()
	todoDescription := todoInput.Text()
	newTodo := Todo{
		Description: todoDescription,
		Status:      "in_progress",
		Created_at:  time.Now(),
	}
	newTodo.createNewTodo()
}

func updateTodo() {
	// Exercise for next class
	// by using bufio scanner read which todo number he wants to update
	// and also ask for new description
	// convert the string:: index into interger format
	// once you get the index, you get the todo from slice Todos
	// call todo.updateTodo()
	// return
}

func deleteTodo() {
	// Exercise for next class
	// same as above but read only index, convert it into int, get the todo from Todos and call delete method
}

func markTodoAsComplete() {
	// Exercise for next class
	// also make sure to update todo completed_at time.
}
