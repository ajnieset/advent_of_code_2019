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
	var opcodes = make([]int64, 173)

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

	for i := 0; i < 173; i++ {
		opcode, err := r.ReadString(delim)
		if opcode[:len(opcode)-1] != ""{
			opcode = opcode[:len(opcode)-1]
		}
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Value of I: %d\n", i)
				fmt.Printf("Value of last number: %q\n", opcode)
				break
			}
			fmt.Println(err)
			return
		}

		opcodes[i], err = strconv.ParseInt(opcode,10,0)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}