package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	var totFuel float64

	const delim = '\n'

	file, err := os.Open("advent_of_code/day1_part1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	r := bufio.NewReader(file)

	totFuel = 0
	for i := 0; i < 100; i++ {
		line, err := r.ReadString(delim)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		line = line[:len(line)-1]

		mass, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		totFuel += calculate_fuel_recurse(mass)
	}

	fmt.Printf("Total Fuel needed for part 2: %f\n", totFuel)
}

func calculate_fuel_recurse(mass float64) float64 {

	recurse_mass := math.Floor(mass/3) - 2
	if recurse_mass > 0 {
		recurse_mass += calculate_fuel_recurse(recurse_mass)
		return recurse_mass
	} else {
		return 0
	}
}
