package ansi

import (
	"fmt"
	"regexp"
	"testing"
)

func BenchmarkStripping(b *testing.B) {
	DetectionPattern := regexp.MustCompile(`(?m)(((\x1b\[[;\d]*[A-Za-z])|\x1b]8;;|\x1b\\.*?\x1b|]8;;\x1b\\)*)`)
	str := fmt.Sprint(Black("black text"), "some kept text", ReverseVideo("inverted text"), Green("green text"), Bold("bold text"))

	b.Run("regex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DetectionPattern.ReplaceAllString(str, "")
		}
	})

	b.Run("split", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StripString(str)
		}
	})
}

func TestStrip(t *testing.T) {
	str := fmt.Sprint(Black("black text"), "some kept text", ReverseVideo("inverted text"), Green("green text"), Bold("bold text"))
	if StripString(str) != "black textsome kept textinverted textgreen textbold text" {
		t.Errorf("strip not successful!\ngot:      %s\n%v\nexpected: black textsome kept textinverted textgreen textbold text\n%v", StripString(str), []byte(StripString(str)), []byte("black textsome kept textinverted textgreen textbold text"))
	}

	str = fmt.Sprintf("Link to my homepage: %s", LinkString("https://moritz.sh", "Linktext"))
	if StripString(str) != "Link to my homepage: https://moritz.sh" {
		t.Errorf("strip not successful!\ngot:      %s\n%v\nexpected: Link to my homepage: https://moritz.sh\n%v", StripString(str), []byte(StripString(str)), []byte("Link to my homepage: https://moritz.sh"))
	}
}
