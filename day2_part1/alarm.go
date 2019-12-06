package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
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

	fmt.Printf("Size of r: %d\n", r.Size())

	for {
		input, err := r.ReadString(delim)
		if input[:len(input)-1] != ""{
			input = input[:len(input)-1]
		}
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Value of last number: %q\n", input)
				break
			}
			fmt.Println(err)
			return
		}

		opcode, err := strconv.Atoi(input)
		opcodes = append(opcodes, opcode)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}