package main

import (
	"fmt"
	"time"
)

func displayTodaysBestSeller(f func(string) string) {
	fmt.Println(f("Best selling car on " + time.Now().Format("01/02/2006")))
}

func addTwo() func() int {
	sum := 0
	return func() int {
		sum += 2
		return sum
	}
}


func main() {
	make := "Lexus"
	closure := func(label string) string {
		return label + ": " + make
	}
	fmt.Printf("closure type: %T\n", closure)
	displayTodaysBestSeller(closure)


	twoMore := addTwo()
	fmt.Println(twoMore())
	fmt.Println(twoMore())
}
