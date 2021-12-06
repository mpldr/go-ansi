# go-ansi

![Version](https://img.shields.io/badge/dynamic/json?color=green&label=Version&query=%24.latest_version&url=https%3A%2F%2Fmpldr.codes%2Fansi%2Fapi%2Fversions%2Flatest&style=flat-square&logo=git&color=F05032)
![License: MIT](https://img.shields.io/static/v1?color=green&label=License&message=MIT&style=flat-square&logo=open-source-initiative&color=3DA639)

```
import "mpldr.codes/ansi"
```

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

## License

This library is licensed under the MIT License which can be found
[here](https://git.sr.ht/~poldi1405/go-ansi/tree/master/item/LICENSE).
