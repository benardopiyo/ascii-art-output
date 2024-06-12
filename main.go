package main

import (
	o "ascii-art-output/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	var outputFile string
	var input string
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	if len(os.Args) == 2 {
		if !strings.HasPrefix(os.Args[1], "-") {
			input = os.Args[1]
			outputFile = "banner.txt"
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
	}
	if len(os.Args) == 4 {
		if !(strings.ToLower(os.Args[3]) == "standard" || strings.ToLower(os.Args[3]) == "shadow" || strings.ToLower(os.Args[3]) == "thinkertoy") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
		if !strings.HasPrefix(os.Args[1], "--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
		outputFile = os.Args[1][9:]
		if !strings.HasSuffix(outputFile, ".txt") {
			fmt.Println("The output file should be a '.txt' file. Please confirm and rerun.\n\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(1)
		}
		input = os.Args[2]
	}
	if len(os.Args) == 3 {
		if !strings.HasPrefix(os.Args[1], "--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
		input = os.Args[2]
		outputFile = os.Args[1][9:]
		if !strings.HasSuffix(outputFile, ".txt") {
			fmt.Println("The output file should be a '.txt' file. Please confirm and rerun.\n\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		}
	}

	banner := "banners/" + o.FileName()
	art := o.BannerArt(banner)
	oupt := o.Art(input, art)
	o.CreateFile(outputFile, oupt)

}
