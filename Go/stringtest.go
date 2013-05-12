package main

import (
	"strings"
	"time"
	"fmt"
	"bytes"
)

const (
	timeFmt  = " tooks %10v ; "
	testFmt  =   "Test " + timeFmt + "lengths = %v\n"
	stageFmt = "\nStage" + timeFmt + "lengths = %v, coeffs = %v\n"
	packFmt  =   "Pack " + timeFmt + "repeat = %d, count = %d\n"
	totalFmt =   "Total" + timeFmt + "\n\n"
)

func line(pat string) string {
	return strings.Repeat(pat, 64) + "\n\n"
}

var singleLine = line("-")
var doubleLine = line("=")

func main() {
	var total time.Duration
	fmt.Printf(doubleLine)
	total += RunTestPack(100000000, 1)
	total += RunTestPack(1000000, 100)
	total += RunTestPack(10000, 10000)
	total += RunTestPack(100, 1000000)
	fmt.Printf(totalFmt, total)
}

// --------

func RunTestPack(repeat, count int) (result time.Duration) {
	result += RunTestStage(repeat, count, []int{1, 2, 5, 10, 20, 50, 100}, []int{1})
	result += RunTestStage(repeat, count, []int{1, 2, 5, 10, 20, 50},      []int{1, 2})
	result += RunTestStage(repeat, count, []int{1, 2, 5, 10},              []int{1, 10, 2, 5})
	fmt.Printf(packFmt, result, repeat, count)
	fmt.Printf(doubleLine)
	return
}

func RunTestStage(repeat, count int, lengths, coeffs []int) (result time.Duration) {
	for _, length := range lengths {
		lens := makeLengths(length, coeffs)
		result += RunTest(repeat, count, lens)
	}
	fmt.Printf(stageFmt, result, lengths, coeffs)
	fmt.Printf(singleLine)
	return
}

func RunTest(repeat, count int, lengths []int) (result time.Duration) {
	strs  := makeStrings(lengths)
	start := time.Now()
	for rep := 0; rep < repeat; rep++ {
		var buf bytes.Buffer
		for num := 0; num < count; num++ {
			for _, str := range strs {
				buf.WriteString(str)
			}
		}
		buf.String()
	}
	result = time.Since(start)
	fmt.Printf(testFmt, result, lengths)
	return 
}

// --------

func makeStrings(lengths []int) []string {
	result := make([]string, len(lengths))
	for i, length := range lengths {
		result[i] = strings.Repeat("a", length)
	}
	return result
}

func makeLengths(length int, coeffs []int) []int {
	result := make([]int, len(coeffs))
	for i, coeff := range coeffs {
		result[i] = length * coeff
	}
	return result
}