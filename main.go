package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mygo/test"
	"net/http"
	"time"
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

	//var m map[string]string

	//m := make(map[string]string)
	//m["a"] = "b"
	//if m == nil {
	//	fmt.Println("m is empty")
	//}
	//fmt.Printf("m is : %v\n", m)
	//
	//s := make([]string, 1)
	//s[0] = "1234"
	//s = append(s, "2345")
	//fmt.Printf("s is : %v\n", s)
	//
	//a := [1]string{"12345"}
	//a[0] = "1234"
	//fmt.Printf("a is : %v\n", a)


	//i := 2
	//
	//if i > 1 {
	//	//var err error
	//	i, err := doDivision(i, 2)
	//	if err != nil {
	//		panic(err) //捕获异常
	//	}
	//	fmt.Println(i)
	//}
	//
	//fmt.Println(i)

	test.TestChuandi()

	//fanOut()
	//selectDrop()
	//withTimeout()
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

	//var arr []int
	//
	//for i := 1; i < 10; i++ {
	//	arr = append(arr, i)
	//}
	//
	//var wg sync.WaitGroup
	//
	////a := make(chan int)
	////b := make(chan int)
	//
	//
	//for j := 1; j <= 100; j++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		test.GetSum(arr)
	//	}()
	//}
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	go test.GetSum2(arr)
	//}()
	//
	//wg.Wait()

	//var sum1 int
	//var sum2 int
	//<- a
	//<- b
	//
	//fmt.Printf("Sum1:%d, Sum2:%d, all:%d\n", sum1, sum2, sum1 + sum2)
	//router := gin.Default()
	//
	//router.GET("/test1", func1)
	//router.POST("/test2", func2)
	//router.POST("/test3", func3)
	//
	////url 族群
	//group1 := router.Group("/g1")
	//group1.GET("/test4", func1)
	//
	////重定向
	//router.GET("test5", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	//})
	//
	//router.Run(":8888")
}


func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			}()
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		emps--
		}
	}


func selectDrop() {
	const cap = 5

	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
		}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
			case ch <- "paper":
				fmt.Println("manager : send ack")
			default:
				fmt.Println("manager : drop")
			}
		}
	close(ch)

	}


func withTimeout() {
	duration := 50 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch <- "paper"
		}()

	select {
		case p := <-ch:
			fmt.Println("work complete", p)

		case <-ctx.Done():
			fmt.Println("moving on")
		}
	}

func doDivision(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Input is invalid")
	}

	return x/y, nil
}