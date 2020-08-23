package ansi

import (
	"fmt"
	"strconv"
)

// Black wraps the content in ANSI codes to make its foreground color black
func Fraktur(content ...interface{}) string {
	return SetFraktur() + fmt.Sprint(content...) + UnsetFraktur()
}

// SetBlack sets the foreground color to black
func SetFraktur() string {
	return escape + fraktur + set
}

// UnsetBlack resets the foreground color from black to default.
func UnsetFraktur() string {
	return escape + frakturOff + set
}

// Black wraps the content in ANSI codes to make its foreground color black
func Font(font int, content ...interface{}) string {
	return SetFont(font) + fmt.Sprint(content...) + UnsetFont()
}

// SetBlack sets the foreground color to black
func SetFont(font int) string {
	return escape + strconv.Itoa(font+10) + set
}

// UnsetBlack resets the foreground color from black to default.
func UnsetFont() string {
	return escape + frakturOff + set
}
