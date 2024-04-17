package main

import (
	"api/pkg/tasks"
	"fmt"
	"time"
)

func main() {
	// Criar algumas tarefas de exemplo
	task1 := tasks.Task{
		Title:       "Fazer compras",
		Description: "Comprar leite e pão",
		CreatedAt:   time.Now(),
	}

	task2 := tasks.Task{
		Title:       "Estudar Go",
		Description: "Ler o livro de Go",
		CreatedAt:   time.Now(),
	}

	// Adicionar as tarefas a uma lista de tarefas
	taskList := tasks.TaskList{Tasks: []tasks.Task{task1, task2}}

	// Exibir as tarefas
	fmt.Println("Tarefas:")
	for _, task := range taskList.Tasks {
		fmt.Printf("- %s\n", task.Title)
	}
}
