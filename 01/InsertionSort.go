package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"
	"strconv"
)
func main(){
	//str := "/src/main/ADSI/01/in1.txt"
	//str, _ := os.Getwd()
	//path := strings.Join([]string{str, "\\src\\main\\ADSI\\01\\in1.txt"},"")
	//fmt.Println(path)
	//os.Open(path)
	inF, inE := os.Open("./src\\main\\ADSI\\01\\in1.txt")
	if inE != nil {
		fmt.Printf("An error occurred on opening the inputfile \nDoes the file exist?\nHave you got acces to it?\n")
		return // exit the function on error
	}
	defer inF.Close()
	inR := bufio.NewReader(inF)
	var Length int
	var arr []int
	for i:=0;;i++{
		if i==0 {
			inS, rE := inR.ReadString('\n')
			temp := strings.Replace(inS, "\r\n", "", -1)
			//fmt.Printf("The input was : %s", temp)
			i, _ := strconv.Atoi(temp)
			//fmt.Println(i)
			Length=i
			arr =make([]int,Length,100)
			if rE == io.EOF {
				return
			}
		} else {
			inS, rE := inR.ReadString(' ')
			//fmt.Printf("The input was : %s", inS)
			temp := strings.Replace(inS, " ", "", -1)
			num, _ := strconv.Atoi(temp)
			//fmt.Println(num)
			arr[i-1]=num
			if rE == io.EOF {
				break
			}
		}
		//printbe(arr)
		
	}
	//var Length int
	//fmt.Scanln(&Length)
	//arr :=make([]int,Length,100)
	//for i:=0; i<Length; i++ {
	//	fmt.Scanf("%d", &arr[i])
	//}
	for i:=0 ; i< len(arr) ; i++ {	
		key := arr[i]
		//insert A[i] into the sorted sequence A[0,...,n-1]
		j := i-1
		for j >= 0 && arr[j] >key {
			arr[j+1] =arr[j]
			j--
		}
		arr[j+1] = key
		writeToFile(arr)
		//printbe(arr)
	}
	//io writestring to file

}
func writeToFile(arr []int){
	var writeString string
	writeString= convert(arr)
	writeString+="\n"
	//fmt.Println(writeString)	
	var filename = "./src\\main\\ADSI\\01\\out1.txt"
	var f *os.File
	var err1 error
	
	if checkFileIsExist(filename) {
		f, err1 =os.OpenFile(filename, os.O_APPEND, 0666)
		fmt.Println("exist file")
	} else {
		f, err1 =os.Create(filename)
		fmt.Println("no exist, it will create")
	}
	check(err1)
	
	n, err1 := io.WriteString(f, writeString)
	check(err1)
	fmt.Println(n)
	f.Close()
}
//covert int array to string
func convert(arr []int) string {
	s :=make([]string,len(arr))
	for i :=range arr{
		s[i]=strconv.Itoa(arr[i])
	}
	return strings.Join(s, " ")
}
func check(e error){
	if e!=nil {
		panic(e)
	}
}
func checkFileIsExist(filename string) bool {
	var exist =true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist =false
	}
	return exist
}
func printbe(arr []int){
	for i, v := range arr {
		if arr[i] == v {
			fmt.Printf("%d ",v)
		}
	}
	fmt.Println()
}