package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file := "input1.txt"
	f, err := os.Open(file)
	checkError(err)
	defer f.Close()
	s := bufio.NewScanner(f)
	var temp int
	var count int
	for s.Scan() {
		// fmt.Println(s.Text())
		// fmt.Println(s.Text(), s.Bytes())
		num, err := strconv.Atoi(s.Text())
		// fmt.Println(strconv.Atoi(s.Text()))
		checkError(err)
		if num > temp {
			count = count + 1
			temp = num

		}
		temp = num
	}
	fmt.Println(count - 1)
	// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate)
	// resp, err := http.Get("https://adventofcode.com/2021/day/1/input")
	// if err != nil {
	// 	panic("error during Http get")
	// }
	// defer resp.Body.Close()

	// doc, err := html.Parse(resp.Body)
	// if err != nil {
	// 	panic("Something wrong parsin html body")
	// }

	// fmt.Printf("%+v", doc)
	// // fmt.Println(doc.Data)
	// // fmt.Println(doc.Namespace)
	// // fmt.Println(doc.FirstChild.Data)
	// fmt.Printf("%+v", doc.FirstChild.LastChild)
	// fmt.Printf("%+v", doc.FirstChild.LastChild.FirstChild)

}
