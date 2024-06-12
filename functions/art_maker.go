package functions

import (
	"fmt"
	"os"
	"strings"
)

//Fuction takes a string and map of ascii patterns
//Checks if it is a printable ascii character
//If yes, the it generates the pattern in a continous string format(while also respecting "\n")
//Otherwise, it notifies the user the existence of a unprintable character in the input
func Art(s string, mYmap map[int][]string) string {
	var str string
	s = Tab(s)
	s = Backspace(s)
	s = strings.ReplaceAll(s, "\n", "\\n")
	input := strings.Split(s, "\\n")
	// generating ascii-art for a string
	for j, word := range input {
		if word == "" && j != len(input)-1 {
			str += fmt.Sprintln()
			continue
		} else if word != "" {
			for i := 0; i < 8; i++ {
				for _, cha := range word {
					// check if character is a printable chararcter
					if ok := (cha >= ' ' && cha <= rune(126)) || (cha == '\n'); ok {
						str += mYmap[int(cha)][i]
					} else if !ok {
						fmt.Println("Unprintable character within input")
						os.Exit(1)
					}
				}
				str += fmt.Sprintln()
			}
		}
	}
	return str
}
