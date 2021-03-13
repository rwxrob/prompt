package prompt_test

import (
	"fmt"

	"github.com/rwxrob/prompt-go"
)

func ExamplePlain() {
	it := prompt.Plain("Enter something: ")
	fmt.Println("Got:", it)
}

func ExampleSecret() {
	it := prompt.Secret("Enter secret (you won't see): ")
	fmt.Println("Got:", it)
}

func ExampleStrict() {
	it := prompt.Strict("Enter a number: ", "[0-9]")
	fmt.Println("Got:", it)
}
