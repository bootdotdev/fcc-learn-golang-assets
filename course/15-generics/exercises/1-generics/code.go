package main

import "fmt"

func getLast[]() {

}

// don't edit below this line

type email struct {
	message        string
	senderEmail    string
	recipientEmail string
}

type payment struct {
	amount         int
	senderEmail    string
	recipientEmail string
}

func main() {
	test([]email{}, "email")
	test([]email{
		{
			"Hi Margo",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"Hey Margo I really wanna chat",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"ANSWER ME",
			"janet@example.com",
			"margo@example.com",
		},
	}, "email")
	test([]payment{
		{
			5,
			"jane@example.com",
			"sally@example.com",
		},
		{
			25,
			"jane@example.com",
			"mark@example.com",
		},
		{
			1,
			"jane@example.com",
			"sally@example.com",
		},
		{
			16,
			"jane@example.com",
			"margo@example.com",
		},
	}, "payment")
}

func test[T any](s []T, desc string) {
	last := getLast(s)
	fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))
	for i, v := range s {
		fmt.Printf("Item #%v: %v\n", i+1, v)
	}
	fmt.Printf("Last item in list: %v\n", last)
	fmt.Println(" --- ")
}
