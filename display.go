package ansi

import "fmt"

// ClearRight clears all characters right of the cursor from the display
func ClearRight() string {
	return escape + fmt.Sprint(clearLine, 0)
}

// ClearRight clears all characters left of the cursor from the display
func ClearLeft() string {
	return escape + fmt.Sprint(clearLine, 1)
}

// ClearRight clears all characters on the line
func ClearLine() string {
	return escape + fmt.Sprint(clearLine, 2)
}

// ClearScreen clears the entire screen and moves the cursor to (0;0)
func ClearScreen() string {
	return escape + fmt.Sprint(clearScreen, 2)
}

// ClearScreen clears the entire screen and moves the cursor to (0;0)
func ClearBegin() string {
	return escape + fmt.Sprint(clearScreen, 1)
}

// ClearScreen clears the entire screen and moves the cursor to (0;0)
func ClearEnd() string {
	return escape + fmt.Sprint(clearScreen, 0)
}

// ClearScreen clears the entire screen and moves the cursor to (0;0)
func ClearScreenBuffer() string {
	return escape + fmt.Sprint(clearScreen, 3)
}
