package main

import (
	"fmt"
)

func main() {
	var counter, temp, flag int
	fmt.Scanf("%d", &temp)
	counter = temp
	arr := make([]int, counter, 1000)
	for i := 0; i < counter; i++ {
		fmt.Scanf("%d", &temp)
		arr[i] = temp
	}
	for i := 0; i < len(arr); i++ {
		mini := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[mini] {
				mini = j
			}
		}
		if mini != i {
			temp = arr[i]
			arr[i] = arr[mini]
			arr[mini] = temp
			flag++
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

/*1 for i = 0 to A.length-1
2     mini = i
3     for j = i to A.length-1
4         if A[j] < A[mini]
5             mini = j
6     swap A[i] and A[mini]*/
