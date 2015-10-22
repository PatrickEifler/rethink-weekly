package main

import (
	"bufio"
	"fmt"
	"os"
)

type newsletter struct{}

// Send out newsletter
func (n *newsletter) send() {
}

// Finding outstandng issues and ask us chooese one to send out
func (n *newsletter) menu() {
	fmt.Println("Start to send out news letter")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select an issue:")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func runNewsLetter() {
}
