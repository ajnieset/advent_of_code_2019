package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const delim = ","
	//opcodes := []int{1,1,1,4,99,5,6,0,99}

	file := "github.com/ajnieset/advent_of_code/day2_part1/input.txt"

	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	splits := strings.Split(string(input), delim)

	opcodes := make([]int, 0, len(splits))

	for _, i := range splits{
		opcode, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		opcodes = append(opcodes, opcode)
	}

	opcodes[1] = 12
	opcodes[2] = 2
	
	final_code := opcode_reader(opcodes)

	fmt.Printf("Opcode at Pos 0: %v\n", final_code)

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
