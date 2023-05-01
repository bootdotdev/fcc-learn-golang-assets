package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

// Don't touch above this line

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", &m.Recipient)
	fmt.Printf("Message: %v\n", &m.Text)
}

// Don't touch below this line

func test(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	sendMessage(m)
	fmt.Println("=====================================")
}

func main() {
	test("Lane", "Textio is getting better everyday!")
	test("Allan", "This pointer stuff is weird...")
	test("Tiffany", "What time will you be home for dinner?")
}
