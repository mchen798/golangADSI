package main

import (
	"fmt"
)

func main() {
	flag := 0
	var temp int
	var counter int
	fmt.Scanf("%d", &temp)
	counter = temp
	arr := make([]int, counter, 1000)
	for i := 0; i < counter; i++ {
		fmt.Scanf("%d", &temp)
		arr[i] = temp
	}
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
				flag++
			}
		}
	}
	putout(arr)
	fmt.Println(flag)
}
func putout(arr []int) {
	i := 0
	for i = 0; i < len(arr)-1; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Print(arr[i])
	fmt.Println()
}
