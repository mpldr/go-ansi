package main

import (
	"fmt"
	"time"

	"gitlab.com/poldi1405/go-ansi"
)

func main() {
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
	fmt.Print(ansi.UpX(3), "5")
	fmt.Print(ansi.MoveCursor(0, 0))
}
