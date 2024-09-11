package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Usage: First arg - input filename, Second - symbol, third - output filename.\nExample ./trimmser first.sh \\r second.sh\n")
	inputFilename := os.Args[1:2]
	symbol := os.Args[2:3]
	outputFilename := os.Args[3:4]
	data, _ := trimStrings(inputFilename[0], symbol[0])
	writeFile(data, outputFilename[0])
}

func writeFile(somedata []string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, line := range somedata {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Done")
}

func trimStrings(filename string, symbol string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strings.Trim(line, symbol)
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
