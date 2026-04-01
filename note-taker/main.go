package main

import (
	"bufio"
	"errors"
	"fmt"
	"note-taker/note"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting note taker app")

	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	note := note.New(title, content)

	note.PrintNote()

	note.Save()
}

func getNoteData() (string, string, error) {
	noteTitle, errTitle := getUserInput("Insert title: ")
	noteContent, errContent := getUserInput("Insert content: ")

	if errTitle != nil {
		return "", "", errTitle
	}

	if errContent != nil {
		return "", "", errContent
	}

	return noteTitle, noteContent, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("An erro ocurred while reading")
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil
}
