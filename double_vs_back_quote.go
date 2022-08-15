package main

import "fmt"

func main() {
	// double vs back quotes
	doubleQuotes := "Hello\n Go!"
	backQuotes := `Hello\n Go!`

	fmt.Printf("doubleQuotes: %s, type: %T\n", doubleQuotes, doubleQuotes)
	fmt.Printf("backQuotes: %s, type: %T\n", backQuotes, backQuotes)

	/* output
	doubleQuotes: Hello
	Go!, type: string
	backQuotes: Hello\n Go!, type: string
	*/

	// error: more than one character in rune literal
	// singleQuotes := 'Hello\n Go!'
}
