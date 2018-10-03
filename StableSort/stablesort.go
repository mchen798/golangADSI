package main

import "fmt"

func main() {
	var counter, temp int
	fmt.Scanf("%d", &temp)
	counter = temp
	arr := make([]int, counter, 10000)
	char := make([]byte, counter, 10000)
	arr1 := make([]int, counter, 10000)
	char1 := make([]byte, counter, 10000)
	for i := 0; i < counter; i++ {
		fmt.Scanf("%c%d", &char[i], &arr[i])
	}
	chartar, arrtar := findstable(arr, char)
	copy(arr1, arr)
	copy(char1, char)
	for i := 0; i < counter; i++ {
		for j := counter - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, char, j, j-1)
			}
		}
	}
	output(arr, char)
	Isstable(arr, char, arrtar, chartar)
	for i := 0; i < counter; i++ {
		mini := i
		j := 0
		for j = i; j < counter; j++ {
			if arr1[j] < arr1[mini] {
				mini = j
			}
		}
		swap(arr1, char1, i, mini)
	}
	output(arr1, char1)
	Isstable(arr1, char1, arrtar, chartar)
}
func output(arr []int, char []byte) {
	i := 0
	for i = 0; i < len(arr)-1; i++ {
		fmt.Printf("%c%d ", char[i], arr[i])
	}
	fmt.Printf("%c%d", char[i], arr[i])
	fmt.Println()
}
func swap(arr []int, char []byte, j int, k int) {
	temp := arr[j]
	tempc := char[j]
	arr[j] = arr[k]
	char[j] = char[k]
	arr[k] = temp
	char[k] = tempc
}
func Isstable(arr []int, char []byte, arrtar []int, chartar []byte) {
	for j := 0; j < len(arrtar); j++ {
		for i := 0; i < len(arr); i++ {
			if arr[i] == arrtar[j] {
				if char[i] != chartar[j] {
					fmt.Println("Not stable")
					return
				}
				break
			}
		}
	}
	fmt.Println("Stable")
}
func findstable(arr []int, char []byte) ([]byte, []int) {
	arrtemp := make([]int, len(arr), 1000)
	chartemp := make([]byte, len(char), 1000)
	copy(arrtemp, arr)
	copy(chartemp, char)
	var dytararr []int
	var dytarbyte []byte
	var temp int = -1
	for i := 0; i < len(arrtemp); i++ {
		if arrtemp[i] == 0 {
			continue
		}
		for j := i + 1; j < len(arrtemp); j++ {
			if temp == arrtemp[j] {
				arrtemp[j] = 0
			}
			if arrtemp[i] != 0 && arrtemp[i] == arrtemp[j] {
				//fmt.Println(char[i],arr[i])
				dytararr = append(dytararr, arrtemp[i])
				dytarbyte = append(dytarbyte, chartemp[i])
				temp = arrtemp[i]
				arrtemp[j] = 0
				//break
			}

		}
	}
	return dytarbyte, dytararr
}
