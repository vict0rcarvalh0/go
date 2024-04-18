package main

import (
	"fmt"
	"os"
	"time"
	"todo/pkg/tasks"
)

func main() {
	taskList := tasks.TaskList{}

	// Carregar tarefas do arquivo (se existir)
	err := taskList.CarregarFromFile("tasks.json")
	if err != nil {
		fmt.Println("Erro ao carregar tarefas:", err)
	}

	defer func() {
		// Salvar tarefas no arquivo ao sair
		err := taskList.SalvarToFile("tasks.json")
		if err != nil {
			fmt.Println("Erro ao salvar tarefas:", err)
		}
	}()

	// Exibir opções
	fmt.Println("O que você deseja fazer?")
	fmt.Println("1. Adicionar uma tarefa")
	fmt.Println("2. Atualizar uma tarefa")
	fmt.Println("3. Excluir uma tarefa")
	fmt.Println("4. Listar todas as tarefas")
	fmt.Println("5. Sair")

	var opcao int
	fmt.Scanln(&opcao)

	switch opcao {
	case 1:
		var title, description string
		fmt.Println("Digite o título da tarefa:")
		fmt.Scanln(&title)
		fmt.Println("Digite a descrição da tarefa:")
		fmt.Scanln(&description)
		taskList.AdicionarTask(title, description)
		fmt.Println("Tarefa adicionada com sucesso!")
	case 2:
		var title, newTitle, newDescription string
		fmt.Println("Digite o título da tarefa que deseja atualizar:")
		fmt.Scanln(&title)
		fmt.Println("Digite o novo título da tarefa:")
		fmt.Scanln(&newTitle)
		fmt.Println("Digite a nova descrição da tarefa:")
		fmt.Scanln(&newDescription)
		err := taskList.AtualizarTask(title, newTitle, newDescription)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Tarefa atualizada com sucesso!")
		}
	case 3:
		var title string
		fmt.Println("Digite o título da tarefa que deseja excluir:")
		fmt.Scanln(&title)
		err := taskList.ExcluirTask(title)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Tarefa excluída com sucesso!")
		}
	case 4:
		fmt.Println("Tarefas:")
		for _, task := range taskList.Tasks {
			fmt.Printf("- %s\n", task.Title)
			fmt.Printf("  Descrição: %s\n", task.Description)
			fmt.Printf("  Criada em: %s\n", task.CreatedAt.Format(time.RFC3339))
			if !task.CompletedAt.IsZero() {
				fmt.Printf("  Concluída em: %s\n", task.CompletedAt.Format(time.RFC3339))
			}
		}
	case 5:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida")
	}
}
