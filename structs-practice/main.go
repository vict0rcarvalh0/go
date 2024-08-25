package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"structs-practice/note"
	"structs-practice/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

// embedding interfaces
type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo content:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo) // can pass todo as param because it implements the Save method

	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)  // standard input
	text, err := reader.ReadString('\n') // single quotes are used for single characters

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // for Windows

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving data failed")
		return err
	}

	fmt.Println("Data saved successfully")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
