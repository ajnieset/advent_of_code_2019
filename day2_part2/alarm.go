package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	const delim = ","
	const final_solution = 19690720

	file := "github.com/ajnieset/advent_of_code/day2_part1/input.txt"

	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	splits := strings.Split(string(input), delim)

	opcodes := create_opcodes(splits)

	var final_code int
	for j:=0; j < 100; j++ {
		opcodes[1] = j
		for k:=0; k < 100; k++ {
			opcodes[2] = k

			final_code = opcode_reader(opcodes)
			if final_code == final_solution {
				break
			} else {
				opcodes = create_opcodes(splits)
			}
		}
		if final_code == final_solution {
			break
		} else {
			opcodes = create_opcodes(splits)
		}
	}

	//final_code := opcode_reader(opcodes)

	if final_code == final_solution {
		fmt.Printf("Opcode at Pos 1: %v\n", opcodes[1])
		fmt.Printf("Opcode at Pos 2: %v\n", opcodes[2])
		fmt.Printf("Opcode at Pos 2: %v\n", 100 * opcodes[1] + opcodes[2])
	} else {
		fmt.Println("Solution not found probably and error")
	}

}

func create_opcodes(splits []string) []int {
	opcodes := make([]int, 0, len(splits))

	for _, i := range splits {
		opcode, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		opcodes = append(opcodes, opcode)
	}
	return opcodes
}

func opcode_reader(opcodes []int) int {

	pos := 0
	for {
		if opcodes[pos] == 99 {
			break
		} else if opcodes[pos] == 1 {
			opcodes[opcodes[pos+3]] = opcodes[opcodes[pos+1]] + opcodes[opcodes[pos+2]]
		} else if opcodes[pos] == 2 {
			opcodes[opcodes[pos+3]] = opcodes[opcodes[pos+1]] * opcodes[opcodes[pos+2]]
		}
		pos += 4
	}

	return opcodes[0]
}
