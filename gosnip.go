// This is a library of snipets
package gosnip

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Spinner spins for the specified duration, printing a spinner character to the terminal every 50 milliseconds.
func Spinner(duration time.Duration) {
	endTime := time.Now().Add(duration)
	for time.Now().Before(endTime) {
		for _, r := range `|/-\` {
			fmt.Printf("\r%c", r)
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Printf("\r ")
}

// ClearScreen clears the terminal screen.
func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// BPrintln prints a beautified message with optional emojis.
// It prints the emojis (if any), followed by the message, with a newline before and after, and the text in bright blue.
// The emojis are passed as a variable number of rune arguments, and they are printed before the message.
// A space is added between the emojis and the text for better readability.
func BPrintln(msg string, emojis ...rune) {
	fmt.Print("\n\033[1;34m") // start a new line and set the text color to bright blue
	for _, emoji := range emojis {
		fmt.Print(string(emoji) + " ") // print the emoji followed by a space
	}
	fmt.Println(msg, "\033[0m") // print the message and reset the text color to the default
}
