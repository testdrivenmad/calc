package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		usage()
	}

	op := os.Args[1]
	nums := make([]float64, 0, len(os.Args)-2)
	for _, a := range os.Args[2:] {
		n, err := strconv.ParseFloat(a, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid number %q: %v\n", a, err)
			os.Exit(2)
		}
		nums = append(nums, n)
	}

	result := nums[0]
	switch op {
	case "add", "+":
		for _, n := range nums[1:] {
			result += n
		}
	case "sub", "-":
		for _, n := range nums[1:] {
			result -= n
		}
	default:
		usage()
	}

	fmt.Println(strconv.FormatFloat(result, 'f', -1, 64))
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <add|sub> <num> <num> [num ...]\n", os.Args[0])
	os.Exit(2)
}
