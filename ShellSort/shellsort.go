package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var cnt int
var err error
var inputReader *bufio.Reader

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	var temp int
	var input string
	input, err = inputReader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	temp, err = strconv.Atoi(input)
	//fmt.Scanf("%d",&temp)
	A := make([]int, temp, 1000000)
	for i := 0; i < temp; i++ {
		input, err = inputReader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		A[i], err = strconv.Atoi(input)
		//fmt.Scanf("%d",&A[i])
	}
	//t1 := time.Now()
	shellSort(A, temp)
	//elapsed1 := time.Since(t1)
	//fmt.Println("loop: ", elapsed1)
}
func shellSort(A []int, n int) {
	var m int
	G := make([]int, n, 1000000)
	G1 := make([]int, n, 1000000)
	j := 0
	for m = 1; ; m++ {
		temp := (int(math.Pow(float64(3), float64(m))) - 1) / 2
		//fmt.Println(temp)
		if temp > n {
			break
		}
		G1[j] = temp
		j++
	}
	for i := 0; i < j; i++ {
		G[i] = G1[j-i-1]
	}
	m = m - 1
	fmt.Println(m)
	for i := 0; i < m; i++ {
		fmt.Printf("%d ", G[i])
		insertionSort(A, n, G[i])
	}
	fmt.Println()
	fmt.Println(cnt)
	for i := 0; i < n; i++ {
		fmt.Println(A[i])
	}
}
func insertionSort(A []int, n int, g int) {
	for i := g; i < n; i++ {
		v := A[i]
		j := i - g
		for j >= 0 && A[j] > v {
			A[j+g] = A[j]
			j = j - g
			cnt++
		}
		A[j+g] = v
	}
}
