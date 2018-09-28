package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var inputReader *bufio.Reader
var input int
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	var input string
	var temp int
	var counter int
	input, err = inputReader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	temp, err = strconv.Atoi(input)
	counter = temp
	var seller, buyer, pri, diff int
	seller = 0
	diff = math.MinInt64
	buyer = math.MaxInt64
	t1 := time.Now() // get current time
	for i := 0; i < counter; i++ {
		input, err = inputReader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		temp, err = strconv.Atoi(input)
		if i != 0 && diff < temp-pri {
			diff = temp - pri
		}
		if temp < buyer && i < counter-1 {
			buyer = temp
			seller = 0
			pri = temp
			continue
		}
		if temp > seller {
			seller = temp
		}
		if seller != 0 && diff < seller-buyer {
			diff = seller - buyer
		}
		pri = temp
	}
	elapsed1 := time.Since(t1)
	fmt.Println("loop: ", elapsed1)
	fmt.Println(diff)
}
