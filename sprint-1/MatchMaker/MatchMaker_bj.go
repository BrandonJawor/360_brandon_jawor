package main

import (
	"fmt"
	"strconv"
)

/*
Cheat Sheet:
Question 1:	5
Question 2:	5
Question 3:	1
Question 4:	1
Question 5:	5
*/

const weight1 int = 1
const weight2 int = 2
const weight3 int = 3
const desire1 int = 1 //negative
const desire5 int = 5 //positive
const compatgood int = 85
const compatneutral int = 71
const compatbad int = 70

func validate(r string, q string) int {
	i, err := strconv.Atoi(r)
	if err != nil {
		fmt.Println("*ERROR*")
		fmt.Println("Please enter an integer")
		fmt.Println(q)
		fmt.Scanln(&r)
		i = validate(r, q)
	}
	if i < 1 || i > 5 {
		fmt.Println("*ERROR*")
		fmt.Println("Please enter a number between 1 and 5.")
		fmt.Println(q)
		fmt.Scanln(&r)
		i = validate(r, q)
	}
	return i
}

func printresponse(total int) {
	if total >= compatgood {
		fmt.Println("We should go on a date sometime.")
	} else if total >= compatneutral {
		fmt.Println("I think we can be good friends.")
	} else if total <= compatbad {
		fmt.Println("I don't believe we will get along.")
	}
}

func calcscore(rv int, weight int, desire int) int {
	var score int
	var newrv int
	if desire == 5 {
		score = rv * (weight * 2)
	} else if desire == 1 {
		if rv == 1 {
			newrv = 5
			score = newrv * (weight * 2)
		} else if rv == 2 {
			newrv = 4
			score = newrv * (weight * 2)
		} else if rv == 3 {
			newrv = 3
			score = newrv * (weight * 2)
		} else if rv == 4 {
			newrv = 2
			score = newrv * (weight * 2)
		} else if rv == 5 {
			newrv = 1
			score = newrv * (weight * 2)
		}
	}
	return score
}

func main() {
	fmt.Println("************************     MATCH MAKER     ************************")
	fmt.Println("Instructions:")
	fmt.Println("Answer questions according to if you agree or disagree.")
	fmt.Println("Five being highly agree and one being highly disagree.")
	fmt.Println("Results for your compatibility score will be calculated at the end.")
	fmt.Println("")
	var question1 string = "Pizza is the best food of all time." //weight 3 positive
	fmt.Println(question1)
	var response1 string
	var response1val int
	fmt.Scanln(&response1)
	response1val = validate(response1, question1)

	fmt.Println("")
	var question2 string = "Video games are fun to play." //weight 1 positive
	fmt.Println(question2)
	var response2 string
	var response2val int
	fmt.Scanln(&response2)
	response2val = validate(response2, question2)

	fmt.Println("")
	var question3 string = "Baseball is a fun sport." //weight 1 negative
	fmt.Println(question3)
	var response3 string
	var response3val int
	fmt.Scanln(&response3)
	response3val = validate(response3, question3)

	fmt.Println("")
	var question4 string = "Chinese food taste great." //weight 3 negative
	fmt.Println(question4)
	var response4 string
	var response4val int
	fmt.Scanln(&response4)
	response4val = validate(response4, question4)

	fmt.Println("")
	var question5 string = "Basketball is a fun sport." //weight 2 positive
	fmt.Println(question5)
	var response5 string
	var response5val int
	fmt.Scanln(&response5)
	response5val = validate(response5, question5)

	var score1 int = calcscore(response1val, weight3, desire5) //weight 3 positive
	var score2 int = calcscore(response2val, weight1, desire5) //weight 1 positive
	var score3 int = calcscore(response3val, weight1, desire1) //weight 1 negative
	var score4 int = calcscore(response4val, weight3, desire1) //weight 3 negative
	var score5 int = calcscore(response5val, weight2, desire5) //weight 2 positive
	var total int = score1 + score2 + score3 + score4 + score5

	fmt.Println("")
	fmt.Printf("Your compatibility score is %d.\n", total)
	printresponse(total)
}
