package ansi

import "os"

var nocolorIsSet bool

func init() {
	_, nocolorIsSet = os.LookupEnv("NO_COLOR")
}
