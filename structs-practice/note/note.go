package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", note.Title, note.Content, note.CreatedAt)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note) // convert the note struct to a JSON string

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // permission to write to the owner and read to everyone else
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		fmt.Println("Please provide all required fields")
		return &Note{}, errors.New("invalid input")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
