package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	res := answer(convertInput(sc.Text()))
	fmt.Println(res)
}

func convertInput(s string) []int {
	sp := strings.Split(s, " ")
	var input []int
	for _, v := range sp {
		i, _ := strconv.Atoi(v)
		input = append(input, i)
	}
	return input
}

func answer(input []int) int {
	var maxV int = -1000
	var minV int = input[0]
	for _, v := range input {
		now := v - minV
		if maxV < now {
			maxV = now
		}
		if minV > v {
			minV = v
		}
	}
	return maxV
}
