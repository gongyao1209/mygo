package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"mygo/db"
	"mygo/model"
)

func main() {

	Db := db.Get()
	defer Db.Close()

	push := model.GetPush(Db,32001)

	fmt.Printf("%v", push)

	Re := db.GetRedis()
	defer Re.Close()
	_, err := Re.Do("SET", "mykey", "superWang")
	if err != nil {

	}

	value, err := redis.String(Re.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", value)
	}
	//fmt.Printf("value %v", value)
}