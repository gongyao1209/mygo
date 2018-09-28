package main

import (
	"mygo/test"
)

func main() { //这里主要是试了一下信道的用法

	max := 1000
	ch := make(chan int, max) //缓存
	test.TestChannel(ch, max)

	return
}