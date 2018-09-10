package test

import (
	"fmt"
	"time"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
}

func Sort(data Interface) {
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := data.Len()
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	}

type Person struct {
	Name string
	Age int
}

func (p Person)GetName() string {
	return p.Name
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

//-----------
var number int

func init()  {
	number = 1
}

func Get(param int) int {
	if param == 0 {
		return number
	} else {
		return param + 1
	}
}

func GetSum(arr []int) int {
	var sum int
	for _, value := range arr {
		sum += value
		time.Sleep(time.Second)
	}

	fmt.Printf("GetSum: %d\n", sum)
	return sum
}

func GetSum2(arr []int) int {
	var sum int
	for _, value := range arr {
		if sum == 0 {
			sum = value
		} else {
			sum *= value
		}

		time.Sleep(time.Second)
	}

	fmt.Printf("%v\n", arr)
	fmt.Printf("GetSum2: %d\n", sum)
	return sum
}