// 本套学习教程基于 https://www.yuque.com/aceld/mo95lb/pz6nag
package main //包名

import (
	"fmt" //导入依赖
	"strconv"
)

// 初始化函数
func init() {
	fmt.Println("init")
}

// 主函数
func main() {
	//先声明
	var str string
	var i int
	//再赋值
	str = "ddd"
	i = 188
	fmt.Print("我是 str: ", str, "我是 i: ", i)

	//自动判断类型
	var b = false
	var i2 = 99
	fmt.Print("我是 b: ", b, "我是 i2: ", i2)

	//:=
	dev := "Go Developers"
	fmt.Println("Go Hello World!", dev) //自动换行
	fmt.Print("Go Hello World!", dev)   //不换行

	//多变量声明
	var x, y int
	//赋值
	x = 10
	y = 100
	fmt.Print("坐标", x, ", ", y)

	var ()

	//varShow()
	//varListShow()
	//constShow()
	//doubleReturnShow()
	//showSwap()
	//
	//hello2()
	//
	//hello3()
	//
	//hello4()

	round10000()
}

// 声明变量 不赋值
var empty string

// 定义变量 并赋值
var name = "wuta"

func varShow() {

	//定义变量 省略 var 只能定义到方法里
	novar := 998

	fmt.Println(empty, name, novar)

}

// 一次声明多个变量
var a, b, c, d = 1, 2, 3, 4

// 声明多个不同类型的变量
var i, s, f = 123, "字", 16.5

func varListShow() {
	fmt.Println(a, b, c, d, i, s, f)
}

// 声明常量
const bookTitle = "百科全书"
const price float32 = 6.66

// 枚举
const (
	BOY  = 1
	GIRL = 0
)

func constShow() {
	fmt.Println(bookTitle, "的价格是", price)
	fmt.Println(BOY, GIRL)
}

// 函数返回多个值
func doubleReturn(input0 int) (string, string) {

	//int显式转换string
	result0 := strconv.Itoa(input0 * input0)
	result1 := "ddd" + strconv.Itoa(input0)

	return result0, result1
}

func doubleReturnShow() {
	a, b := doubleReturn(3)
	fmt.Println(a, b)
}

// 类似于go的直接改内存地址
func swap(x *int, y *int) {
	var temp = *x // 保存 x 地址上的值
	*x = *y       // 将 y 值赋给 x
	*y = temp     // 将 temp 值赋给 y
}

func showSwap() {
	a := 1
	b := 2

	fmt.Println("地址：", "a是", &a, "\t", "b是", &b)
	fmt.Println("值：", "a是", a, "\t", "b是", b)
	swap(&a, &b)
	fmt.Println("地址：", "a是", &a, "\t", "b是", &b)
	fmt.Println("值：", "a是", a, "\t", "b是", b)
}

func goFunc(i int) {
	fmt.Println("goroutine ", i, " ...")
}

func round10000() {
	//time.Sleep(time.Second)
	for i := 0; i < 10000; i++ {
		go goFunc(i) //开启一个并发协程
	}
	//time.Sleep(time.Microsecond * 1)
	fmt.Println("goroutine  end")
}
