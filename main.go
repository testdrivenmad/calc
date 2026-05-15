package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// Use a compact log prefix so diagnostic output stays distinguishable
	// from the program's stdout result.
	log.SetFlags(0)
	log.SetPrefix("calc: ")

	log.Printf("starting with %d argument(s): %v", len(os.Args)-1, os.Args[1:])

	if len(os.Args) < 4 {
		log.Printf("too few arguments (need operator and at least 2 numbers); printing usage")
		usage()
	}

	op := os.Args[1]
	log.Printf("operator = %q", op)

	log.Printf("parsing %d operand(s) as float64", len(os.Args)-2)
	nums := make([]float64, 0, len(os.Args)-2)
	for _, a := range os.Args[2:] {
		n, err := strconv.ParseFloat(a, 64)
		if err != nil {
			log.Printf("failed to parse operand %q", a)
			fmt.Fprintf(os.Stderr, "invalid number %q: %v\n", a, err)
			os.Exit(2)
		}
		nums = append(nums, n)
	}
	log.Printf("parsed operands: %v", nums)

	result := nums[0]
	switch op {
	case "add", "+":
		log.Printf("performing addition across %d operand(s)", len(nums))
		for _, n := range nums[1:] {
			result += n
		}
	case "sub", "-":
		log.Printf("performing subtraction across %d operand(s)", len(nums))
		for _, n := range nums[1:] {
			result -= n
		}
	case "mod", "%":
		log.Printf("performing modulus across %d operand(s)", len(nums))
		for _, n := range nums[1:] {
			if n == 0 {
				log.Printf("aborting: modulus by zero encountered")
				fmt.Fprintln(os.Stderr, "modulus by zero")
				os.Exit(2)
			}
			result = math.Mod(result, n)
		}
	default:
		log.Printf("unknown operator %q; printing usage", op)
		usage()
	}

	log.Printf("computation complete; writing result to stdout")
	fmt.Println(strconv.FormatFloat(result, 'f', -1, 64))
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <add|sub|mod> <num> <num> [num ...]\n", os.Args[0])
	os.Exit(2)
}
