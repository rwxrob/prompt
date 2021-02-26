package prompt

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"syscall"

	"golang.org/x/term"
)

// ScanLine uses the bufio.NewScanner (which defaults to limited line
// scanning) to safely scan a reasonably long line of standard input.
// Note that other methods could introduce dangerous input reads that
// could consume all available memory. Note this does not check for
// scanner error conditions returning an empty string instead in such
// cases.
func ScanLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Plain simply prompts the user to enter text interactively.
func Plain(s string) string {
	fmt.Printf("%v", s)
	return ScanLine()
}

// Secret prompts the user to enter text interactively but does
// not echo what they are typing to the screen. Warning: only works
// where the syscall.Stdin is supported.
func Secret(s string) string {
	fmt.Printf("%v", s)
	input, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	return string(input)
}

// Strict prompts the user to enter text interactively that must
// strictly match the regular expression provided.
func Strict(s, pattern string) string {
	fmt.Printf("%v", s)
	input := ScanLine()
	if matched, _ := regexp.MatchString(pattern, input); !matched {
		return ""
	}
	return input
}

// StrictSecret prompts the user to enter text interactively that must
// strictly match the regular expression provided but without echoing
// the input to the terminal.
func StrictSecret(s, pattern string) string {
	fmt.Printf("%v", s)
	input, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return ""
	}
	if matched, _ := regexp.Match(pattern, input); !matched {
		return ""
	}
	return string(input)
}

// Until is identical to Plain but loops until at least one
// non-whitespace character is entered, which is simply a convenient way
// of writing UntilStrict("some","\\S").
func Until(s string) string { return UntilStrict(s, "\\S") }

// UntilSecret is identical to Secret but loops until at least one
// non-whitespace character is entered, which is simply a convenient way
// of writing UntilStrictSecret("some","\\S").
func UntilSecret(s string) string { return UntilStrictSecret(s, "\\S") }

// UntilStrictSecret is identical to StrictSecret but prompts until the
// input matches the given pattern.
func UntilStrictSecret(s, pattern string) string {
	for {
		val := StrictSecret(s, pattern)
		if val != "" {
			return val
		}
	}
}

// UntilStrict prompts continuously until the user response matches.
// Note that if the pattern cannot be matched that this function will
// never return.
func UntilStrict(s, pattern string) string {
	for {
		val := Strict(s, pattern)
		if val != "" {
			return val
		}
	}
}
