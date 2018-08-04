package main

import (
	"fmt"
)
func main(){
	var a,b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	if a<b {
		temp := a
		a=b
		b=temp
	}
	i := gcd(a,b)
	fmt.Println(i)
}
func gcd(a int,b int) int {
	if b==0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
