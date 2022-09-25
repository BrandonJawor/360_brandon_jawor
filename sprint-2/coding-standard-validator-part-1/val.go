package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var numOfComments int = 0
var toManyChar int = 0
var nothingWrittenLicense int = 0
var nothingWrittenReadMe int = 0
var validChecks int = 0
var countLineLicense int = 0
var countLineReadMe int = 0

func main() {
	f, err := os.Open("dummy.go")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var length = len([]rune(scanner.Text()))
		if length >= 100 {
			toManyChar = 1 //used to confirm line exceeds 100 characters
		}
		if strings.Contains(scanner.Text(), "//") {
			numOfComments += 1 //counts each comment
		}
		if strings.Contains(scanner.Text(), "/*") {
			numOfComments += 1
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	l, err := os.Open("LICENSE")

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	scanner = bufio.NewScanner(l)
	for scanner.Scan() {
		var length = len([]rune(scanner.Text()))
		if length != 0 {
			countLineLicense += 1 // counts lines that contain something written on it
		}
		if countLineLicense == 0 {
			nothingWrittenLicense = 1 // used to confirm if that nothing is written in file
		}
	}

	r, err := os.Open("README.md")

	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	scanner = bufio.NewScanner(l)
	for scanner.Scan() {
		var length = len([]rune(scanner.Text()))
		if length != 0 {
			countLineReadMe += 1 // counts lines that contain something written on it
		}
		if countLineReadMe == 0 {
			nothingWrittenReadMe = 1 // used to confirm if that nothing is written in file
		}
	}

	if toManyChar == 1 {
		fmt.Println("File contains line that exceeds 100 characters.")
	} else {
		fmt.Println("File does not exceed 100 characters per line.")
		validChecks += 1
	}

	if numOfComments == 0 {
		fmt.Println("File does not contain any comments.")
	} else {
		fmt.Printf("File contains %d comments.\n", numOfComments)
		validChecks += 1
	}

	if nothingWrittenLicense == 1 {
		fmt.Println("License file does not contain any contents.")
	} else {
		fmt.Println("License file is valid")
		validChecks += 1
	}
	if nothingWrittenReadMe == 1 {
		fmt.Println("ReadMe file does not contain any contents.")
	} else {
		fmt.Println("ReadMe file is valid")
		validChecks += 1
	}
	fmt.Printf("These files meet %d out of the 4 validation standards.\n", validChecks)

}
