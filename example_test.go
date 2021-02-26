package prompt_test

import (
	"fmt"

	"github.com/rwxrob/prompt"
)

func ExamplePlain() {
	it := prompt.Plain("Enter something: ")
	fmt.Println("Got:", it)
}

func ExampleSecret() {
	it, err := prompt.Secret("Enter secret (you won't see): ")
	fmt.Println("Got:", it)
	fmt.Println("Error:", err)
}

func ExampleStrict() {
	it, err := prompt.Strict("Enter a number: ", "[0-9]")
	fmt.Println("Got:", it)
	fmt.Println("Error:", err)
}
