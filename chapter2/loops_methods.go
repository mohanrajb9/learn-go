package main

import (
	"fmt"
	"time"
	"strings"
	"reflect"
	"bufio"
	"os"
	"strconv"
	"math/rand"
)

func main(){
	// time package method Now() returns the current local time which is of time.Time type
	var currentTime time.Time = time.Now()
	// the value currentTime has a method Year() which returns the year as an int
	var currentYear int = currentTime.Year()

	fmt.Println("Current Year is:", currentYear)

	var city string = "Chennai"
	replacerObj := strings.NewReplacer("C", "c")
	fmt.Println(reflect.TypeOf(replacerObj))
	replacerObj.Replace(city)
	fmt.Println("After replacing C with c:", city)


	fmt.Print("Enter grade:")
	reader := bufio.NewReader(os.Stdin)
	user_input, err := reader.ReadString('\n')

	grade, _ := strconv.ParseFloat(strings.TrimSpace(user_input), 64)
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	fmt.Println("user input is", user_input)
	fmt.Println("user input number is", grade)
	if grade >= 90 {
		fmt.Println("You got A grade")
	} else if grade >= 80 {
		fmt.Println("You got B grade")
	} else {
		fmt.Println("You need to work hard")
	}

	GuessNum()

}


func GuessNum() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	rand_num := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)

	for x := 0; x < 10; x ++ {
		fmt.Println("current this number of guesses left:", 10-x)
		fmt.Print("Guess  the number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error in reading string")
			return
		}
		input = strings.TrimSpace(input)
		userGuess, err := strconv.Atoi(input)
		if rand_num < userGuess {
			fmt.Println("your guess is HIGH")
			continue
		} else if rand_num > userGuess {
			fmt.Println("your guess is LOW")
			continue
		} else {
			fmt.Println("you guessed right", rand_num, userGuess)
			break
		}
	}
}