package main

import (
	"fmt"
	"mygo/db"
	"mygo/model"
)

func main() {

	Db := db.Get()

	push := model.GetPush(Db,32001)

	fmt.Printf("%v", push)
}