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

//------值传递
type person struct {
	name string
	age int
	isDead bool
}

func TestChuandi()  {
	p1 := person{name:"gongyao", age:20}
	p2 := person{name:"wanghui", age:50}
	p3 := person{name:"yaoke", age:100}
	//
	//people := []person{p1, p2, p3}
	//whoIsDead1_(people)
	//
	//for _, p := range people {
	//	if p.isDead {
	//		fmt.Printf("Who Is dead? %s\n", p.name)
	//	}
	//}


	//people2 := make([]*person, 0)
	//people2 = append(people2, &p1)
	//people2 = append(people2, &p2)
	//people2 = append(people2, &p3)
	//whoIsDead2(people2)
	//for _, p := range people2 {
	//	if p.isDead {
	//		fmt.Printf("Who Is dead? %s\n", p.name)
	//	}
	//}

	m_people := make(map[string]*person)
	m_people[p1.name] = &p1
	m_people[p2.name] = &p2
	m_people[p3.name] = &p3
	whoIsDead_m1(m_people)

	for _, p := range m_people {
		if p.isDead {
			fmt.Printf("Who Is dead? %s\n", p.name)
		}
	}

	//m2_people := map[string]person{
	//	p1.name : p1,
	//	p2.name : p2,
	//	p3.name : p3,
	//}
	
}

func whoIsDead_m1(persons map[string]*person)  {
	for _, per := range persons {
		if per.age < 50 {
			//per
			persons[per.name].isDead = true
		}
	}
}

func whoIsDead1(people []person)  {
	for _, p := range people {
		if p.age < 50 {
			p.isDead = true
		}
	}
}

func whoIsDead1_(people []person)  {
	for i := 0; i < len(people); i++ {
		if people[i].age < 50 {
			people[i].isDead = true
		}
	}
}

func whoIsDead2(people []*person)  {
	for _, p := range people { //和 php foreach as 异曲同工，都是 复制
		if p.age < 50 {
			p.isDead = true
		}
	}
}