package main

import "fmt"

func qweight1p(r int) int {
	if r == 1 {
		qs = 2
	} else if r == 2 {
		qs = 4
	} else if r == 3 {
		qs = 6
	} else if r == 4 {
		qs = 8
	} else if r == 5 {
		qs = 10
	}
	return qs
}

func qweight1n(r int) int {
	if r == 1 {
		qs = 10
	} else if r == 2 {
		qs = 8
	} else if r == 3 {
		qs = 6
	} else if r == 4 {
		qs = 4
	} else if r == 5 {
		qs = 2
	}
	return qs
}
func qweight2p(r int) int {
	if r == 1 {
		qs = 4
	} else if r == 2 {
		qs = 8
	} else if r == 3 {
		qs = 12
	} else if r == 4 {
		qs = 16
	} else if r == 5 {
		qs = 20
	}
	return qs
} /*
func qweight2n(r int) int {	//unused
	if r == 1 {
		qs = 20
	} else if r == 2 {
		qs = 16
	} else if r == 3 {
		qs = 12
	} else if r == 4 {
		qs = 8
	} else if r == 5 {
		qs = 4
	}
	return qs
}*/
func qweight3p(r int) int {
	if r == 1 {
		qs = 6
	} else if r == 2 {
		qs = 12
	} else if r == 3 {
		qs = 18
	} else if r == 4 {
		qs = 24
	} else if r == 5 {
		qs = 30
	}
	return qs
}
func qweight3n(r int) int {
	if r == 1 {
		qs = 30
	} else if r == 2 {
		qs = 24
	} else if r == 3 {
		qs = 18
	} else if r == 4 {
		qs = 12
	} else if r == 5 {
		qs = 6
	}
	return qs
}
func printresponse(total int) {
	if total >= 75 {
		fmt.Println("We should go on a date sometime.")
	} else if total >= 50 {
		fmt.Println("I think we can be good friends.")
	} else {
		fmt.Println("I don't think we should talk anymore.")
	}

}

var response1 int
var response2 int
var response3 int
var response4 int
var response5 int

func main() {
	fmt.Println("Pizza is one of the best types of food.") //w3p
	fmt.Scanln(&response1)

	fmt.Println("Video games are fun to play.") //w1p
	fmt.Scanln(&response2)

	fmt.Println("Baseball is a fun sport.") //w1n
	fmt.Scanln(&response3)

	fmt.Println("Chinese food taste great.") //w3n
	fmt.Scanln(&response4)

	fmt.Println("Basketball is a fun sport.") //w2p
	fmt.Scanln(&response5)

	var score1 int = qweight3p(response1)
	var score2 int = qweight1p(response2)
	var score3 int = qweight1n(response3)
	var score4 int = qweight3n(response4)
	var score5 int = qweight2p(response5)

	var total int = score1 + score2 + score3 + score4 + score5

	printresponse(total)

}
