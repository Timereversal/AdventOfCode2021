package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMovements(t *testing.T) {

	var sub submarine2
	filename := "inputDay2P2.txt"
	// filename := "inputDay2P1_test1.txt"
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		direction := strings.Split(s.Text(), " ")
		steps, err := strconv.Atoi(direction[1])
		if err != nil {
			t.Errorf("parsing direction number generate error %v", err)
		}
		switch direction[0] {
		case "down":
			sub.down(steps)
		case "up":
			sub.up(steps)
		case "forward":
			sub.forward(steps)
		}
	}
	if sub.horizontal != 15 || sub.depth != 60 {
		t.Errorf("expected values horizontal %d, depth %d got %d, %d and prodcut h*depth %d", 15, 60, sub.horizontal, sub.depth, sub.horizontal*sub.depth)

	}

}
