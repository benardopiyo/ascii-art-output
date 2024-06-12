package functions

import (
	"bufio"
	"fmt"
	"os"
)

// Creates a map of ascii art from a chosen banner file
func BannerArt(s string) map[int][]string {
	var myMap = make(map[int][]string)
	var slice []string
	CheckStatus(s)
	file, er := os.Open(s)
	if er != nil {
		fmt.Println("Error opening file. Check file")
		os.Exit(1)
	}
	defer file.Close()
	var count int
	var key int = 32
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count++
		if scanner.Text() != "" {
			slice = append(slice, scanner.Text())
		}
		if count == 9 {
			myMap[key] = slice
			key++
			count = 0
			slice = nil
		}
	}

	return myMap
}
