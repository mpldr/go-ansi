package main

import (
	"fmt"
	"time"

	"git.sr.ht/~poldi1405/go-ansi"
)

func main() {
	goto link
	fmt.Println("Styles:", "normal", ansi.Bold("bold"), ansi.Faint("faint"), ansi.Italic("italic"), ansi.Underscore("underscore"), ansi.DoubleUnderscore("double underscore"), ansi.Blink("blinking"), ansi.FastBlink("blinking fast"), ansi.ReverseVideo("FG and BG exchanged"), "["+ansi.Conceal("concealed")+"]", ansi.Strikethrough("strikethrough"))
	fmt.Println("Colors:", "default", ansi.Black("black"), ansi.Red("red"), ansi.Green("green"), ansi.Yellow("yellow"), ansi.Blue("blue"), ansi.Magenta("magenta"), ansi.Cyan("cyan"), ansi.White("white"), ansi.Color256(25, "256color"), ansi.ColorTrue(172, 138, 39, "TrueColor"))
	fmt.Println("Colors (BG):", "default", ansi.BlackBG("black"), ansi.RedBG("red"), ansi.GreenBG("green"), ansi.YellowBG("yellow"), ansi.BlueBG("blue"), ansi.MagentaBG("magenta"), ansi.CyanBG("cyan"), ansi.WhiteBG("white"), ansi.Color256BG(25, "256color"), ansi.ColorTrueBG(172, 138, 39, "TrueColor"))
	fmt.Print("Fonts: ")
	for i := 0; i < 10; i++ {
		fmt.Print(" " + ansi.Font(i, "Font #", i) + "")
	}
	fmt.Println(" " + ansi.Fraktur("Fraktur"))

	fmt.Println("█████████████████████████████████")
	fmt.Println("█████████████████████████████████")
	fmt.Println("█████████████████████████████████")
	fmt.Println("█████████████████████████████████")
	fmt.Println("█████████████████████████████████")
	fmt.Println("█████████████████████████████████")
	fmt.Print("█████████████████████████████████")
	fmt.Print(ansi.Left(), "9")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.LeftX(18), "8")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.UpX(3), ansi.Left(), "5")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.RightX(16), ansi.UpX(3), "3")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\r", ansi.DownX(6), "7")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.RightX(14), ansi.UpX(6), "2")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.RightX(16), ansi.DownX(3), "6")
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\r", ansi.UpX(3), "1")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.Left(), ansi.DownX(3), "4")
	time.Sleep(500 * time.Millisecond)

	fmt.Print(ansi.DownX(16), ansi.RightX(8), "saving this position ->", ansi.SavePos())
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.MoveCursor(5, 12), ansi.ReverseVideo("now i am here"))
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.MoveCursor(18, 8), ansi.ReverseVideo("and now i'm here"))
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.MoveCursor(3, 54), ansi.ReverseVideo("hi there!"))
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.MoveCursor(9, 80), ansi.ReverseVideo("what did you expect?"))
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.MoveCursor(25, 60), ansi.ReverseVideo("so… where were we?"))
	time.Sleep(500 * time.Millisecond)
	fmt.Print(ansi.RestorePos(), ansi.ReverseVideo("We were here"))
link:
	fmt.Println(ansi.LinkString("http://example.com", "an example"))
	fmt.Println(ansi.LinkFile("test.txt", "a file example"))
}
