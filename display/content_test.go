package display

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestLineDiff(t *testing.T) {
	s1 := "This string is equal to the second"
	s2 := "This string is equal to the second"
	diff := calculateDiff(s1, s2)
	if diff != "" {
		t.Fail()
	}

	s1 = "This string is *not* equal to the second"
	diff = calculateDiff(s1, s2)
	if strings.Contains(diff, "string") {
		t.Fail()
	}

	diff = calculateDiff(s2, s1)
	if strings.Contains(diff, "string") {
		t.Fail()
	}

	s2 = "This string is a bit equal to the second"
	diff = calculateDiff(s1, s2)
	if strings.Contains(diff, "string") {
		t.Fail()
	}

	s1 = "💓"
	s2 = "💘"
	diff = calculateDiff(s1, s2)
	if strings.Contains(diff, "string") {
		t.Fail()
	}

	s1 = "abacaba 💓 what's my text"
	s2 = "dead💘beef what's this  text"
	diff = calculateDiff(s1, s2)
	diff2 := calculateDiff(s2, s1)
	_ = diff2
	fmt.Println(s2)
	fmt.Println(diff)
	fmt.Print(s2 + "\r" + diff)
	if utf8.ValidString(diff) {
		t.Fail()
	}
}
