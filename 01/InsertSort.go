package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var Length int
	fmt.Scanln(&Length)
	arr := make([]int, Length, 100)
	for i := 0; i < Length; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		printbe(arr, i)
	}
}
func printbe(arr []int, ref int) {
		str := convert(arr)
		fmt.Printf("%s",str)
	//if ref < len(arr)-1 {
		fmt.Println()
	//}
}
func convert(arr []int) string {
	s :=make([]string, len(arr))
	for i:=range arr {
		s[i] = strconv.Itoa(arr[i])
	}
	return strings.Join(s," ")
}
