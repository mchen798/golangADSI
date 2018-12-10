package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	if !sort.IsSorted(sort.IntSlice(A)) {
		sort.Sort(sort.IntSlice(A))
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
	if !sort.IsSorted(sort.IntSlice(B)) {
		sort.Sort(sort.IntSlice(B))
	}
	binarysearch(A, counterA, B, counterB)

}

/*
type IntSlice []int
func (c IntSlice) Len() int{
	return len(c)
}
func (c IntSlice) Swap(i,j int) {
	c[i],c[j]=c[j],c[i]
}
func (c IntSlice) Less(i,j int) bool{
	return c[i]<c[j]
}*/
func binarysearch(A []int, cA int, B []int, cB int) {
	flag := 0
	for i := 0; i < cB; i++ {
		if findbybinary(A, 0, cA-1, B[i]) == 1 {
			flag++
		}
	}
	fmt.Println(flag)
}
func findbybinary(arr []int, low, high, tar int) int {
	if low > high {
		return 0
	}
	mid := low + (high-low)/2

	if arr[mid] > tar {
		return findbybinary(arr, low, mid-1, tar)
	} else if arr[mid] < tar {
		return findbybinary(arr, mid+1, high, tar)
	}
	return 1
}
