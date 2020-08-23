package ansi

import "fmt"

func Bold(content ...interface{}) string {
	return SetBold() + fmt.Sprint(content...) + UnsetBold()
}

// SetBold makes the text following text become bold
func SetBold() string {
	return escape + bold + set
}

// UnsetBold removes the bold property
func UnsetBold() string {
	return escape + boldOff + set
}

func Faint(content ...interface{}) string {
	return SetFaint() + fmt.Sprint(content...) + UnsetFaint()
}

// SetBold makes the text following text become bold
func SetFaint() string {
	return escape + faint + set
}

// UnsetBold removes the bold property
func UnsetFaint() string {
	return escape + boldOff + set
}

func Italic(content ...interface{}) string {
	return SetItalic() + fmt.Sprint(content...) + UnsetItalic()
}

// SetBold makes the text following text become bold
func SetItalic() string {
	return escape + italic + set
}

// UnsetBold removes the bold property
func UnsetItalic() string {
	return escape + italicOff + set
}

func Underscore(content ...interface{}) string {
	return SetUnderscore() + fmt.Sprint(content...) + UnsetUnderscore()
}

func SetUnderscore() string {
	return escape + underscore + set
}

func UnsetUnderscore() string {
	return escape + underscoreOff + set
}

func DoubleUnderscore(content ...interface{}) string {
	return SetDoubleUnderscore() + fmt.Sprint(content...) + UnsetDoubleUnderscore()
}

func SetDoubleUnderscore() string {
	return escape + doubleUnderscore + set
}

func UnsetDoubleUnderscore() string {
	return escape + underscoreOff + set
}
func Blink(content ...interface{}) string {
	return SetBlink() + fmt.Sprint(content...) + UnsetBlink()
}

// Blink makes the text blink
func SetBlink() string {
	return escape + blink + set
}

// UnBlink stops the text from blinking
func UnsetBlink() string {
	return escape + blinkOff + set
}

func FastBlink(content ...interface{}) string {
	return SetFastBlink() + fmt.Sprint(content...) + UnsetFastBlink()
}

// Blink makes the text blink
func SetFastBlink() string {
	return escape + fastblink + set
}

// UnBlink stops the text from blinking
func UnsetFastBlink() string {
	return escape + fastblinkOff + set
}

func ReverseVideo(content ...interface{}) string {
	return SetReverseVideo() + fmt.Sprint(content...) + UnsetReverseVideo()
}

func SetReverseVideo() string {
	return escape + reverseVideo + set
}

// UnReverseVideo disables line wrapping
func UnsetReverseVideo() string {
	return escape + reverseVideoOff + set
}

func Conceal(content ...interface{}) string {
	return SetConceal() + fmt.Sprint(content...) + UnsetConceal()
}

// Conceal conceals text
func SetConceal() string {
	return escape + conceal + set
}

// UnConceal removes the concealed property
func UnsetConceal() string {
	return escape + concealOff + set
}

func Strikethrough(content ...interface{}) string {
	return SetStrikethrough() + fmt.Sprint(content...) + UnsetStrikethrough()
}

// Conceal conceals text
func SetStrikethrough() string {
	return escape + strikethrough + set
}

// UnConceal removes the concealed property
func UnsetStrikethrough() string {
	return escape + strikethroughOff + set
}

// Reset clears all colors and styles (bold, underscore, blink, concealed)
func Reset() string {
	return escape + reset + set
}
