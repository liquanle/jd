// chengfa project main.go
package main

import (
	"fmt"
)

const PI float32 = 3.1415926

var nCount int

func add(x, y int) (sum float32) {
	sum = float32((x + y)) * PI
	return
}

//多值返回函数
func powfun(a, b, c int) (d, e, f int) {
	d = a * a
	e = b * b
	f = c * c
	return
}

type book struct {
	name   string
	author string
	price  int
}

//方法method
type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human
	school string
}

type employee struct {
	human
	commany string
}

func (h *human) sayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *employee) sayHi() {
	fmt.Println("我的公司名称是", e.commany)
}

func (s *student) sayHi() {
	fmt.Println("我的学校名称是", s.school)
}

func main() {
	//测试函数代码
	defer fmt.Printf("%d + %d = %f\n", 30, 59, add(30, 59))

	for i := 0; i < 5; i++ {
		nCount += i * i
		fmt.Println("结果", nCount)
	}
	fmt.Println("Hello World!")

	var marks = [5]int{35, 22, 38, 92, 87}
	for index, value := range marks {
		fmt.Printf("学号为%d的数学成绩为%d!\n", index, value)
		if curval := value + 5; curval > 90 {
			fmt.Println("他及格了！\n")
		}
	}

	for _, value := range marks {
		fmt.Printf("学号为的数学成绩为%d!\n", value)
		if curval := value + 5; curval > 90 {
			fmt.Println("他及格了！\n")
		}
	}

	fmt.Println("good luck")

	a, b, c := 3, 5, 8

	//注意这里是括号，不是大括号
	var (
		f int = 10
		g int = 15
		h int = 20
	)

	fmt.Println(powfun(a, b, c))
	fmt.Println(powfun(f, g, h))

	//指针
	fmt.Println("指针")
	var nSum int = 23
	var pSum *int = &nSum
	fmt.Printf("nsum 的地址是%x, 值是%d\n", pSum, *pSum)

	//switch
	fmt.Println("switch")
	var ooo int = 0
	for {

		ooo++

		fmt.Println("当前ooo的值", ooo)
		if ooo > 3 {
			break
		}
		switch ooo {
		case 1:
			fmt.Println("值 大于10")

		case 2:
			fmt.Println("值 大于20")

		case 3:
			fmt.Println("值大于90")

		case 100:
			fmt.Println("到100了，退出!")
		}
	}

	//结构体
	fmt.Println("结构体")
	var mybook = book{
		name:   "c++学习手册",
		author: "李三",
		price:  103,
	}

	fmt.Println(mybook)

	//slice
	fmt.Println("slice")
	//定义切片的方法有两种，一种定义一个无大小的数组，另一种是使用make
	var lis = []int{1, 2, 3, 99, 100}
	var lis2 = make([]int, 3)
	fmt.Printf("lis 的大小为%d，容量为%d\n", len(lis), cap(lis))
	lis2 = append(lis2, 123)
	lis2 = append(lis2, 643)
	lis2 = append(lis2, 232)
	lis2 = append(lis2, 532, 673, 77, 21)

	var lis3 = lis2[4:6]
	for in, ele := range lis3 {
		fmt.Printf("slice[%d]的值为%d\n", in, ele)
	}

	fmt.Println(lis2)
	fmt.Println(lis3)

	//map
	fmt.Print("map的使用")
	var myMap = make(map[string]int)
	myMap["liquanle"] = 99
	myMap["wangwei"] = 77
	myMap["wudi"] = 22
	myMap["zhangchuan"] = 86

	fmt.Println("图长度为", len(myMap))

	fmt.Println(myMap)
	for name, mark := range myMap {
		fmt.Printf("mymap[%s]的值为%d\n", name, mark)
	}

	if _, ok := myMap["wudi"]; ok {
		fmt.Println("zhangww元素存在！", myMap["zhangww"])
	} else {
		fmt.Println("zhangww元素不存在！")
	}

	delete(myMap, "wudi")
	fmt.Println(myMap)

	//方法method
	fmt.Println("method===========================")

	var hm1 human = human{name: "lql", age: 40, phone: "13582217261"}
	var st1 student = student{human: hm1, school: "大张庄小学"}
	var com1 employee = employee{human: hm1, commany: "jindi company"}

	hm1.sayHi()
	st1.sayHi()
	com1.sayHi()
}
