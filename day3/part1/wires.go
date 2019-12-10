package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	const delim = ","
	const wire_path_file1 = "/home/ajnieset/go/src/github.com/ajnieset/advent_of_code/day3/path1.txt"
	const wire_path_file2 = "/home/ajnieset/go/src/github.com/ajnieset/advent_of_code/day3/path2.txt"

	input, err := ioutil.ReadFile(wire_path_file1)
	if err != nil {
		fmt.Println(err)
		return
	}
	wire_path1 := strings.Split(string(input), delim)

	input, err = ioutil.ReadFile(wire_path_file2)
	if err != nil {
		fmt.Println(err)
		return
	}
	wire_path2 := strings.Split(string(input), delim)
}