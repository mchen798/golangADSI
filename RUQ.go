package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	var n int
	var out=bufio.NewWriter(os.Stdout)
	var q int
	fmt.Scanf("%d %d",&n,&q)
	n=int(math.Ceil(math.Log2(float64(n))))
	if n==0 {
		n++
	}
	n=1<<uint(n)
	lazy:=make([]int,2*n-1)
	arr:=make([]int,q)
	InitA(n,lazy)
	var a,b,c,d int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d",&a,&b)
		b=b+n-1
		if a==0 {
			fmt.Scanf("%d %d",&c,&d)
			arr[i]=d
			update(b,c+n-1,i,lazy)
		}else {
			out.Flush()
			fmt.Fprintln(out,find(b,n,lazy,arr))
		}
	}
	out.Flush()
}
func InitA(n int,lazy []int){
	for i := 0;i < 2*n-1;i++  {
		lazy[i]=-1
	}

}
func  update(l,r,v int,lazy []int){
	if l+1==r || l==r{
		lazy[l]=v
		lazy[r]=v
	}else{
		if l%2 != 1 {
			lazy[l] = v
			l++
		}
		if r%2 != 0 {
			lazy[r] = v
			r--
		}
		update((l-1)/2,(r-1)/2,v,lazy)
	}
}
func find(b,n int,lazy []int,arr []int) int{
	max:=b
	for ;b>0;{
		b=(b-1)/2
		if lazy[max]<lazy[b]{
			max=b
		}
	}
	var temp int
	if lazy[max] == -1 {
		temp=math.MaxInt32
	}else {
		temp=arr[lazy[max]]
	}
	return temp
}
