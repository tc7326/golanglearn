package main

import (
	"fmt"
	"math"
)

func init() {
	defer fmt.Println("hello3 init")
}

func hello3() {
	showUser()
	showPoint()
	showAssert()
}

// User 定义一个结构体 对象?
type User struct {
	name string
	age  int
	work string
}

// 结构体外声明函数 值修改
func (user User) testUser1() {
	user.work = "杰哥"
}

// 等效
func testUser1p(user User) {
	user.work = "杰哥"
}

// 结构体外声明函数 指针修改
func (user *User) testUser2() {
	user.name = "小瘪三"
}

// 等效
func testUser2p(user *User) {
	user.work = "太子"
}

func showUser() {

	user := User{"特拉法尔加-罗", 20, "医生"}
	fmt.Println(user.name, "今年", user.age, "岁，他的职业是", user.work)

	user.testUser1()
	fmt.Println(user.name, "今年", user.age, "岁，他的职业是", user.work)
	testUser1p(user) //等效
	fmt.Println(user.name, "今年", user.age, "岁，他的职业是", user.work)

	user.testUser2()
	fmt.Println(user.name, "今年", user.age, "岁，他的职业是", user.work)
	testUser2p(&user) //等效
	fmt.Println(user.name, "今年", user.age, "岁，他的职业是", user.work)

}

// Point 点 结构体
type Point struct {
	x, y float64
}

func (start Point) Distance(end Point) float64 {

	return math.Hypot(end.x-start.x, end.y-end.x)
}

func showPoint() {

	start := Point{1, 1}
	end := Point{8, 8}

	//使用选择器 选择函数
	selector := start.Distance
	fmt.Println(selector(end))
	fmt.Println(selector(start))

	//直接调用函数
	fmt.Println(start.Distance(end))
	fmt.Println(end.Distance(start))

	fmt.Println("===")

	//方法表达式
	distance := Point.Distance
	fmt.Println(distance(start, end))
	fmt.Printf("%T\n", distance)

	//方法表达式 传递指针类型
	distance2 := (*Point).Distance
	fmt.Println(distance2(&start, end))
	fmt.Printf("%T\n", distance2)

}

// 接口和断言 判断输入的值是否string类型
func whatUDoing(doing interface{}) string {
	result, ok := doing.(string)
	if !ok {
		fmt.Println("这不是string 返回 空字符串")
		return ""
	}
	fmt.Println("这是string 返回 对应的值")
	return result
}

// 判断数据类型
func switchDoing(doing interface{}) string {

	switch doing.(type) {
	case int:
		return "你输入的是 int"
	case string:
		return "你输入的是 string"
	default:
		return "未知类型"
	}

}

// 断言测试
func showAssert() {

	var one = 1
	fmt.Println(whatUDoing(one))
	var playing = "hi! suck u!!!"
	fmt.Println(whatUDoing(playing))

	fmt.Println(switchDoing(1))
	fmt.Println(switchDoing("MineCraft"))

}
