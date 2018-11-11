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
	var input string
	A := make([]int, 1000, 1000000)
	input, err = inputReader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	//fmt.Printf("%s\n",input)
	var data = strings.Split(input, " ")
	//fmt.Println(len(data))
	i := 0
	for j := 0; j < len(data); j++ {
		/*input, err = inputReader.ReadString(' ')
		fmt.Printf("%s\n",input)
		input = strings.Replace(input, " ", "",-1)*/
		//func Split(s, sep string) []string
		//fmt.Printf("%s\n",data[j])
		if data[j] == "+" {
			A[i] = A[i-2] + A[i-1]
			i = i - 2
			A[i] = A[i+2]
			i++
		} else if data[j] == "-" {
			A[i] = A[i-2] - A[i-1]
			i = i - 2
			A[i] = A[i+2]
			i++
		} else if data[j] == "*" {
			A[i] = A[i-2] * A[i-1]
			i = i - 2
			A[i] = A[i+2]
			i++
		} else {
			//buf := bytes.NewBuffer(data[j]) // b is []byte
			//myfirstint, err := binary.ReadVarint(buf)
			A[i], err = strconv.Atoi(data[j])
			//fmt.Printf("%d\n",A[i])
			i++
		}
		//fmt.Println(i)
	}
	fmt.Printf("%d\n", A[0])

}
