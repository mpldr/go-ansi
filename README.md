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

Moving around is as simple as Specifying a direction `ansi.Up()`

## ANSI Support¹

- movement
- line operations (clearing lines)
- fore- and background colour
- formatting (bold, italic, strikethrough, …)
- stripping ANSI-Codes
- links (even to files)
- notifications

¹) support depends on the terminal emulator used and configuration

## User Choice

This library implements the (informal) [`NO_COLOR`](https://no-color.org/)
standard. This disables all colour-related ANSI-Codes but does not restrict any
other function (like cursor movement or formatting). If this behavior is
unwanted please stay on version 1.3.0 or below.

## Bugs Reports & Feature Requests

Bug Reports and Feature Requests can be sent to the [issue
tracker](https://todo.sr.ht/~poldi1405/issues?title=[go-ansi]%20) or via email
to [~poldi1405/issues@todo.sr.ht](mailto:~poldi1405/issues@todo.sr.ht?subject=[go-ansi]%20)
(possible without creating an account on Sourcehut).

### Patches

Want to improve it yourself? Awesome! Please send your patches to
[~poldi1405/patches@lists.sr.ht](mailto:~poldi1405/patches@lists.sr.ht) (see
submitted patches [here](https://lists.sr.ht/~poldi1405/patches)).

For an easy quickstart, just run these commands:

```shell
$ git config format.subjectPrefix "PATCH go-ansi"
$ git config sendemail.to "~poldi1405/patches@lists.sr.ht"
```

Please sign-off your Commits to confirm the [DCO](https://developercertificate.org/).

## License

This library is licensed under the MIT License which can be found
[here](https://git.sr.ht/~poldi1405/go-ansi/tree/master/item/LICENSE).
