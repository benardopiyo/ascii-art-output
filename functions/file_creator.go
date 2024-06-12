package functions

import (
	"fmt"
	"os"
)

//This function takes a filepath and  variable holding a string literal as arguments
//It then creates the file using os package(with respect to the filepath) and writes the string into the file
//Upon encountering any problems in execution of th two operations, it returns an error
func CreateFile(FilePath string, content string) error {
	file, err := os.Create(FilePath)
	if err != nil {
		return fmt.Errorf("error in file creation: %w", err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error in writing content: %w", err)
	}

	return nil
}
