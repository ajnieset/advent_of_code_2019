package main

import (
	"fmt"
	"math"
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var mass float64
	var totFuel float64

	const delim = '\n'

	var file, err = os.Open("advent_of_code/day_1/input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
        if err := file.Close(); err != nil {
            panic(err)
        }
    }()

	r := bufio.NewReader(file)

	// fmt.Println("Enter a mass: ")
	
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	totFuel = 0
	for i:= 0; i < 100; i++{
		line, err := r.ReadString(delim)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		line = line[:len(line)-1]

		mass, err = strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		totFuel = totFuel + calculate_fuel(mass)
	}

	fmt.Printf("Total Fuel needed: %g\n", totFuel)
}

func calculate_fuel(mass float64) float64 {

	return math.Floor(mass / 3) - 2
}