package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var toManyChar int = 0

func main() {
	f, err := os.Open("test")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var length = len([]rune(scanner.Text()))
		if length >= 100 {
			fmt.Println("over")
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	if toManyChar == 1 {
		fmt.Println("Line exceeds 100 characters")
	}

}