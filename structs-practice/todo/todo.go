package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo) // convert the note struct to a JSON string

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // permission to write to the owner and read to everyone else
}

func New(content string) (*Todo, error) {
	if content == "" {
		fmt.Println("Please provide all required fields")
		return &Todo{}, errors.New("invalid input")
	}

	return &Todo{
		Text: content,
	}, nil
}
