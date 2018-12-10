package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var err error
var inputReader *bufio.Reader

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	var counterA, counterB int
	var input string
	input, err = inputReader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	counterA, err = strconv.Atoi(input)
	//fmt.Scanf("%d",&temp)
	A := make([]int, counterA, 100000)
	for i := 0; i < counterA; i++ {
		if i == counterA-1 {
			input, err = inputReader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)
			A[i], err = strconv.Atoi(input)
		} else {
			input, err = inputReader.ReadString(' ')
			input = strings.Replace(input, " ", "", -1)
			A[i], err = strconv.Atoi(input)
			//fmt.Scanf("%d",&A[i])
		}
	}
	input, err = inputReader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	counterB, err = strconv.Atoi(input)
	//fmt.Scanf("%d",&temp)
	B := make([]int, counterB, 100000)
	for i := 0; i < counterB; i++ {
		if i == counterB-1 {
			input, err = inputReader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)
			B[i], err = strconv.Atoi(input)
		} else {
			input, err = inputReader.ReadString(' ')
			input = strings.Replace(input, " ", "", -1)
			B[i], err = strconv.Atoi(input)
			//fmt.Scanf("%d",&A[i])
		}
	}
	findLinear(A, counterA, B, counterB)

}
func findLinear(A []int, cA int, B []int, cB int) {
	flag := 0
	for i := 0; i < cB; i++ {
		for j := 0; j < cA; j++ {
			if A[j] == B[i] {
				flag++
				break
			}
		}
	}
	fmt.Println(flag)
}
