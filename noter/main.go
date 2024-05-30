package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dexter0323/go-learn/noter/note"
)

func main() {
	title, content := getNoteDate()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	err = note.Save()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteDate() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
