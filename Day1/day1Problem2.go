package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func countSum(input []int, index int, s *bufio.Scanner) {
// var count int

func checkError2(e error) {
	if e != nil {
		panic(e)
	}
}

func countSum(elements int, s *bufio.Scanner) int {
	var input []int
	var count, length, sum int
	for s.Scan() {
		length = length + 1
		num, err := strconv.Atoi(s.Text())
		checkError2(err)
		sum = sum + num
		input = append(input, num)
		if length == elements {
			for s.Scan() {
				num, err := strconv.Atoi(s.Text())
				checkError2(err)
				newSum := sum + num - input[len(input)-elements]
				if newSum > sum {
					count = count + 1
				}
				sum = newSum
				input = append(input, num)
			}
			return count
		}
	}
	return 0
}

func main() {

	file := "input1.txt"
	// file := "input1test.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	totalCount := countSum(3, s)
	fmt.Println("total value", totalCount)

	// intStream := make(chan int, 3)
	// for _, v := range []int{4, 6, 7, 8, 9} {
	// 	fmt.Println(v)
	// 	intStream <- v
	// 	switch <-intStream {
	// 	default:
	// 		fmt.Println("c")
	// 	}
}

// }
