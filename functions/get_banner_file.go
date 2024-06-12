package functions

import (
	"os"
	"strings"
)

//Function assigns the banner file name that is to be used in the program depending on the input(number of cli arguments)
//from the user
func FileName() string {
	if len(os.Args) == 2 || len(os.Args) == 3 {
		return "standard.txt"
	}

	if len(os.Args) == 4 {
		switch strings.ToLower(os.Args[3]) {
		case "standard":
			return "standard.txt"
		case "thinkertoy":
			return "thinkertoy.txt"
		case "shadow":
			return "shadow.txt"
		default:
			return "Invalid bannerfile name"
		}
	}
	return ""
}
