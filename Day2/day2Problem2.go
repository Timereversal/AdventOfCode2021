package main

type submarine2 struct {
	horizontal, depth, aim int
}

func (s *submarine2) down(downSteps int) {
	s.aim = s.aim + downSteps
}
func (s *submarine2) up(upSteps int) {
	s.aim = s.aim - upSteps
}
func (s *submarine2) forward(forwardSteps int) {
	s.horizontal = s.horizontal + forwardSteps
	s.depth = s.depth + s.aim*forwardSteps
}

// func main() {
// 	var sub submarine2
// 	inputFile := "inputDay2P1_test1.txt"
// 	f, err := os.Open(inputFile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	s := bufio.NewScanner(f)
// 	for s.Scan() {
// 		direction := strings.Split(s.Text(), " ")
// 		steps, err := strconv.Atoi(direction[1])
// 		if err != nil {
// 			fmt.Printf("parsing direction number generate error %v", err)
// 		}
// 		switch direction[0] {
// 		case "down":
// 			sub.down(steps)
// 		case "up":
// 			sub.up(steps)
// 		case "forward":
// 			sub.forward(steps)
// 		}
// 	}
// 	fmt.Printf("Submarine Position horizontal %d - depth %d : Final Result %d", sub.horizontal, sub.depth, sub.horizontal*sub.depth)
// }
