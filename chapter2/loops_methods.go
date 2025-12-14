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
}