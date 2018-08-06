package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int
	fmt.Scanf("%d", &counter)
	arr := make([]int, counter, 10000)
	for i := 0; i < counter; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	IsPrime(arr)
}
func IsPrime(arr []int) {
	//var temp int
	n := 0
	var flag int
	for i, v := range arr {
		//temp = arr[i]
		flag = 1
		//if v!=2 && v!=3 {
		rt_v := int(math.Sqrt(float64(arr[i])))
		//fmt.Println(rt_v)
		for j := 2; rt_v >= j; j++ {
			if j > 2 && j%2 == 0 {
				continue
			}
			if v%j == 0 {
				flag = 0
				break
			}
			//fmt.Println(arr[i])
		}
		//fmt.Println("There is out of the loop")
		//}else{
		//	fmt.Println(arr[i])
		//}
		if flag == 1 {
			n++
			//fmt.Println(temp)
		}

	}
	fmt.Println(n)
}
