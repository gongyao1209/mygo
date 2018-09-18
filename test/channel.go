package test

import (
	"fmt"
)

type T int64
type mymap map[string]string

func sum(j int64) *[]int64 {
	res := make([]int64, 0)

	if j == 0 {
		return &res;
	}

	var i,sum int64
	for i = 0; i <= j; i++ {
		sum += i
		res = append(res, sum)
	}

	return &res
}

func ji(j int64) *[]int64 {
	res := make([]int64, 0)

	if j == 0 {
		return &res
	}

	var i,ji int64
	ji = 1
	for i = 1; i <= j; i++ {
		ji *= i

		res = append(res, ji)
	}

	return &res
}

func GetNum(number int64) int64 {
	//var a int64
	const max = 2

	ch1 := make(chan *[]int64)
	defer close(ch1)

	ch2 := make(chan *[]int64)
	defer close(ch2)

	go func() {
		temp := sum(number)
		ch1 <- temp
	}()

	go func() {
		temp := ji(number)
		ch2 <- temp
	}()

	a :=<-ch1
	fmt.Println(a)

	b :=<-ch2
	fmt.Println(b)  //信道 和 goroutine 的 关系 最好是一一对应

	return 1
	//fmt.Println(a)
	//for i := 0; i < max; i++ {
	//	a := <-ch
	//	fmt.Println(a)
	//}
	//
	//return a
}