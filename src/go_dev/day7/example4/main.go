package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("An error occurred with file create\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello World\n"

	for i := 0; i <= 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
