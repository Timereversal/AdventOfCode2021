package main

import (
	"os"
	"testing"
)

func TestTotalMovement(t *testing.T) {

	filename := "inputDay2P1.txt"
	// filename := "inputDay2P1_test1.txt"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	x, y := totalMovement(f)
	result_x, result_y := 1967, 1031
	if x != result_x || y != result_y {
		t.Errorf("position  should be %d %d  got %d %d", result_x, result_y, x, y)
	}

}
