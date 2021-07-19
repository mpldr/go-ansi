# go-ansi

Write pretty terminal output the easy way.

Want your text green? `ansi.Green("I am Shrek")`

Want it black on a red ground? `ansi.RedBG(ansi.Black("You can just
combine them"))`

Moving around is as simple as Specifying direction `ansi.Up()`

## ANSI Support

- movement
- line operations (clearing lines)
- fore- and background colour
- formatting (bold, italic, strikethrough, …)
- stripping ANSI-Codes
- hyperlinks (WIP)

## User Choice

This  library  implements   the  (informal)
[`NO_COLOR`](https://no-color.org/) standard.   This  disables  all
colour-related   ANSI-Codes  but  does  not restrict  any other
function (like  cursor  movement  or  formatting).  If this behavior
is unwanted please stay on version 1.3.0 or below.
