package main

import (
	"fmt"
	"log"
)

func main() {
	// about pointers
	pointers()

	// instead pass by value, pass by pointers and update the original value
    num := 2
	double(&num)
	fmt.Println(num)

	// functions to log errror
	var total_paint float64
	paint, err := calculatePaint(24.0, 53.0)
	if err != nil {
		log.Fatal(err)
	}
	total_paint += paint
	paint, err = calculatePaint(24.0, 51.0)
	if err != nil {
		log.Fatal(err)
	}
	total_paint += paint
	fmt.Printf("Total litres of paint required is %5.2f\n", total_paint)
	paint, err = calculatePaint(-24.0, 51.0)
	if err != nil {
		log.Fatal(err)
	}
}

func pointers() {
	var a int = 58
	aPointer := &a
	fmt.Println("original value", a)
	fmt.Println("pointer value", aPointer)
	fmt.Println("extract value from the pointer var", *aPointer)

}


func double(num *int) {
	*num *= 2
}


func calculatePaint(height float64, width float64) (float64, error) {
	if height < 0 || width < 0 {
		return 0, fmt.Errorf("The values cannot be less than 0, height: %f, width: %f", height, width)
	}
	return (height*width)/10, nil
}