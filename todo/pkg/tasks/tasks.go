package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Task struct {
	Title       string
	Description string
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TaskList struct {
	Tasks []Task
}

// AdicionarTask adiciona uma nova tarefa à lista de tarefas
func (tl *TaskList) AdicionarTask(title, description string) {
	task := Task{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
	tl.Tasks = append(tl.Tasks, task)
}

// AtualizarTask atualiza uma tarefa existente na lista de tarefas
func (tl *TaskList) AtualizarTask(title, newTitle, newDescription string) error {
	for i, task := range tl.Tasks {
		if task.Title == title {
			tl.Tasks[i].Title = newTitle
			tl.Tasks[i].Description = newDescription
			return nil
		}
	}
	return fmt.Errorf("tarefa '%s' não encontrada", title)
}

// ExcluirTask exclui uma tarefa da lista de tarefas
func (tl *TaskList) ExcluirTask(title string) error {
	for i, task := range tl.Tasks {
		if task.Title == title {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("tarefa '%s' não encontrada", title)
}

// SalvarToFile salva a lista de tarefas em um arquivo JSON
func (tl *TaskList) SalvarToFile(filename string) error {
	data, err := json.Marshal(tl)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// CarregarFromFile carrega a lista de tarefas de um arquivo JSON
func (tl *TaskList) CarregarFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, tl)
}
