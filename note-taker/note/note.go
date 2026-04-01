package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Note struct {
	Id        int       `json:"ìd"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(title, content string) *Note {

	if title == "" || content == "" {
		return &Note{}
	}

	return &Note{
		Id:        rand.Int(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (n Note) PrintNote() {
	fmt.Printf("== Nota %v == \n", n.Id)
	fmt.Println("Title: " + n.Title + " \nContent: " + n.Content)
}

func (n Note) Save() error {
	fileName := "note-" + strconv.Itoa(n.Id) + ".json"

	jsonContent, err := json.Marshal(n)

	if err != nil {
		return err
	}

	errWriting := os.WriteFile(fileName, jsonContent, 0644)

	if errWriting != nil {
		return errors.New("An error ocurred while writing file")
	}
	return nil
}
