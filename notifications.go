package ansi

import (
	"fmt"
)

// Link creates a terminal Hyperlink that can be clicked, if the terminal
// emulator supports it.
func SendNotification(title, body string) string {
	return fmt.Sprintf(notification, title, body)
}
