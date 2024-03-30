package main

import (
	"fmt"
	"os"

	"aungpyaephyo.com/go/io/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadTextFile(filePath)
	if err == nil {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", c, c, c)
		// newContent := "Original: " + c + " \n Double the Original: " + c + c
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("ERROR PANIC!! %v", err)
	}

}
