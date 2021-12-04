package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_countSum(t *testing.T) {

	file := "input1test.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	totalCount := countSum(3, s)
	fmt.Println(totalCount)
	if totalCount != 1 {
		t.Error("Incorrect value, expected 1  got: ", totalCount)
	}
}
