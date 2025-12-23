package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// learn slice
	sums_slice, err := learn_slice()
	sum := 0.0
	if err != nil {
		log.Fatal(err)
	}
	for _, each_sum := range sums_slice {
		sum += each_sum
	}
	average := sum / float64(len(sums_slice))
	fmt.Printf("The average is %0.2f\n", average)

	// for learning args sent to porogram while running
	fmt.Println(args_example())

}

func learn_slice() ([]float64, error) {
	var numbers_slice []float64
	// or simply numbers_slice1 := make([]float64, 45)
	file, err := os.Open("C:\\Users\\mohan\\Dev1\\GoProject\\chapter5\\file_list")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			return nil, err
		}
		numbers_slice = append(numbers_slice, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers_slice, nil

}

func args_example() float64 {
	numbers := os.Args[1:]
	sum := 0.0
	for _, number := range numbers {
		number, err := strconv.ParseFloat(number, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	return sum / float64(len(numbers))

}
