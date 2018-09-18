package test

import (
	"fmt"
	"time"
)

type T int64

func sum(j int64) int64 {
	if j == 0 {
		return 0;
	}

	var i,sum int64
	for i = 0; i <= j; i++ {
		sum += i
	}

	return sum
}

func ji(j int64) int64 {
	if j == 0 {
		return 0;
	}

	var i,ji int64
	ji = 1
	for i = 1; i <= j; i++ {
		ji *= i
	}

	return ji
}

func GetNum(number int64) int64 {
	var a int64
	const max = 2
	ch := make(chan int64, max)

	go func() {
		temp := sum(number)
		//a += temp
		ch <- temp
	}()

	go func() {
		temp := ji(number)
		//a += temp
		ch <- temp
	}()

	for i := 0; i < max; i++ {
		a += <-ch
	}

	return a

	timeout := time.After(3 * time.Second)

	//for {
		select {
		case x := <-ch:
			fmt.Println(x)
		case <-timeout:
			fmt.Println("task timeout.")
			close(ch)
			break
		}
	//}

	//for x := range ch {
	//	fmt.Println(x)
	//}
	//
	//close(ch)

	return a
	//var a int64
	//defer func() {
	//	select {
	//	case num :=<-ch:
	//		fmt.Println(num)
	//		a += num
	//	default:
	//		fmt.Println(1111)
	//		break
	//	}
	//
	//	close(ch)
	//}()
	//for {

	//}



	return a


	//ji(number)
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	ch <- sum(number)
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	ch <- ji(number)
	//}()
	//
	//wg.Wait()
	//fmt.Println(ch)
	//
	//var a int64
	//for num := range ch {
	//	a += num;
	//}
	//
	//
	//
	//close(ch)
	//
	//return a
}