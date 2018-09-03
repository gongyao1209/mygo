package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func func1(c *gin.Context)  {
	name := c.Query("name")
	pwd := c.Query("pwd")

	c.String(http.StatusOK, "参数：name:%s, pwd:%s\n", name, pwd)
}

func func2(c *gin.Context)  {
	arr, err := c.GetRawData()
	if err != nil {

	}
	c.String(http.StatusOK, "name: %v\n", arr)
}

type User struct {
	Name string `form:"name" json:"name" binding:"required"`
	Pwd string `form:"pwd" json:"pwd" binding:"required"`
}


func func3(c *gin.Context)  {
	var json User

	if c.BindJSON(&json) == nil {
		fmt.Println(json.Name)
		c.JSON(http.StatusOK, gin.H{"name":json.Name, "pwd":json.Pwd})
	} else {
		c.JSON(404, gin.H{"JSON=== status": "binding Json Error!"})
	}
}

func main() {

	//Db := db.Get()
	//defer Db.Close()
	//
	//push := model.GetPush(Db,32001)
	//
	//fmt.Printf("%v", push)
	//
	//Re := db.GetRedis()
	//defer Re.Close()
	//_, err := Re.Do("SET", "mykey", "superWang")
	//if err != nil {
	//
	//}
	//
	//value, err := redis.String(Re.Do("GET", "mykey"))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//} else {
	//	fmt.Printf("Get mykey: %v \n", value)
	//}
	////fmt.Printf("value %v", value)

	router := gin.Default()

	router.GET("/test1", func1)
	router.POST("/test2", func2)
	router.POST("/test3", func3)

	router.Run(":8888")
}