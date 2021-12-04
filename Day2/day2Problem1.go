package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type movement struct {
	down, up, forward int
}

// func move(p *Position, direction string) *Position {
// 	directionSteps := strings.Split(direction, " ")
// 	switch direction {
// 	case directionSteps[0] == "down":
// 		p.y + dir
// 	}
// }
type submarine struct {
}

func totalMovement(txtFile *os.File) (int, int) {

	// directionSteps:= string.Split(diection)
	var dir map[string]int = map[string]int{"down": 1, "up": -1, "forward": 1}
	var move = map[string]int{}
	s := bufio.NewScanner(txtFile)
	for s.Scan() {
		direction := strings.Split(s.Text(), " ")
		num, err := strconv.Atoi(direction[1])
		if err != nil {
			fmt.Printf("parsing direction number generate error %v", err)

		}
		fmt.Println(direction[0])
		move[direction[0]] = move[direction[0]] + num*dir[direction[0]]
	}
	return move["forward"], move["down"] + move["up"]

}
func newPosition(initialPosition *Position, textInput *os.File) *Position {
	moveX, moveY := totalMovement(textInput)
	initialPosition.x = initialPosition.x + moveX
	initialPosition.y = initialPosition.y + moveY
	return initialPosition

}

func main() {
	// f, err := os.Open("inputDay2P1_test1.txt")
	f, err := os.Open("inputDay2P1.txt")
	if err != nil {
		panic(err)
	}
	x, y := totalMovement(f)
	fmt.Println(x, y)
	fmt.Println(x * y)
}
