package main

import (
	"fmt"
	"time"
)

type todo struct {
	description string
	status      string
	createdAt   time.Time
	completedAt time.Time
}

var todos []todo

func main() {

	item1 := todo{description: "Todo List1 ", status: "new"}
	item2 := todo{description: "Todo List2 ", status: "new"}
	item3 := todo{description: "Todo List3 ", status: "new"}
	item4 := todo{description: "Todo List4 ", status: "new"}

	Add(item1)
	Add(item2)
	Add(item3)
	Add(item4)
	Get()
	Delete(1)
	Get()
	completeTodo(1)
	Get()
	Update(3, "Item list 4")
	Get()
}

func Add(newTodo todo) {
	todos = append(todos, newTodo)
}

func Delete(id int) {
	if id > len(todos) {
		fmt.Printf("Todo with ID %d not available \n", id)
		return
	}
	// arr[0:] arr[:10] -> todos[:id-1], todos[id:]
	todos = append(todos[:id-1], todos[id:]...)
}

func Update(id int, todoDesc string) {
	if id > len(todos) {
		fmt.Printf("Todo with ID %d not available \n", id)
		return
	}
	todos[id-1].description = todoDesc

}

func Get() {
	fmt.Println("Printing all todos")
	for _, val := range todos {
		fmt.Printf("%s ||  %s || %v || %v \n", val.description, val.status, val.createdAt, val.completedAt)
	}
}

func completeTodo(id int) {
	if id > len(todos) {
		fmt.Printf("Todo with ID %d not available \n", id)
		return
	}

	todos[id-1].status = "completed"
	todos[id-1].completedAt = time.Now()
}
