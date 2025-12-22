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
	// to understand and use arrays
	get_average()
	// to read files data
	read_file()
}

func get_average() {

	// declare arrays with literals.
	item_list := [3]float64{56.43, 34.23, 76.45}
	var sum, average float64
	for _, item := range item_list {
		sum += item
	}
	average = sum / float64(len(item_list))
	fmt.Printf("average value is %0.2f\n", average)
}

func read_file() {
	var average, sum float64
	i := 0
	file, err := os.Open("C:\\Users\\mohan\\Dev1\\GoProject\\chapter5\\file_list")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		value, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += value
		i++
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	average = sum/float64(i)
	fmt.Println(average)
	fmt.Printf("The average value of the values in the file is %0.2f\n", average)
}
