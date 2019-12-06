package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	const delim = ','
	var opcodes = make([]int, 0)

	file, err := os.Open("github.com/ajnieset/advent_of_code/day2_part1/input.txt")
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

	for {
		input, err := r.ReadString(delim)
		if input[:len(input)-1] != "" {
			input = input[:len(input)-1]
		}
		if err != nil {
			if err == io.EOF {
				opcode, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println(err)
					return
				}
				opcodes = append(opcodes, opcode)
				break
			}
			fmt.Println(err)
			return
		}

		opcode, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		opcodes = append(opcodes, opcode)
	}
	fmt.Printf("Last Number: %d\n", opcodes[len(opcodes)-1])
	fmt.Printf("Second to Last Number: %d\n", opcodes[len(opcodes)-2])
	fmt.Printf("Size of slice: %d\n", len(opcodes))

}

func opcode_reader(opcodes []int) int {

	pos := 0
	for {
		if opcodes[pos] == 1 {
			val1 := opcodes[opcodes[pos+1]]
			val2 := opcodes[opcodes[pos+2]]

			result := val1 + val2

			opcodes[opcodes[pos+3]] = result
		} else if opcodes[pos] == 99 {
			break
		}
	}

	return opcodes[0]
}
