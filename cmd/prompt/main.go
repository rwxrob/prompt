package main

import (
	"fmt"
	"os"

	"github.com/rwxrob/prompt-go"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "plain" {
		val := prompt.Plain("Plain: ")
		fmt.Printf("You entered: %v\n", val)
		return
	}
	switch os.Args[1] {

	case "until":
		val := prompt.Until("Required: ")
		fmt.Printf("You entered: %v\n", val)
		return

	case "strictsecret":
		val := prompt.StrictSecret("Strict secret number: ", "^[0-9]$")
		fmt.Printf("You entered: %v\n", val)
		return

	case "strict":
		val := prompt.Strict("Single Single Digit: ", "^[0-9]$")
		if val == "" {
			fmt.Println("didn't match anything")
			return
		}
		fmt.Printf("You entered: %v\n", val)
		return

	case "secret":
		val := prompt.Secret("Secret: ")
		fmt.Printf("You entered: %v\n", val)
		return

	case "untilstrict":
		val := prompt.UntilStrict("Single Single Digit: ", "^[0-9]$")
		fmt.Printf("You entered: %v\n", val)
		return

	case "untilsecret":
		val := prompt.UntilSecret("Anything:")
		fmt.Printf("You entered: %v\n", val)
		return

	case "untilstrictsecret":
		val := prompt.UntilStrictSecret("Some digits: ", "^[0-9]+$")
		fmt.Printf("You entered: %v\n", val)
		return

	default:
		fmt.Println("usage: prompt [plain|secret|strict|until|untilstrict|untilstrictsecret]")
	}
}
