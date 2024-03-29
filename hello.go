// 本套学习教程基于 https://www.yuque.com/aceld/mo95lb/pz6nag
package main //包名

import (
	"fmt" //导入依赖
	"strconv"
	"unsafe"
)

// 初始化函数
func init() {
	fmt.Println("init")
}

// 主函数
func main() {

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

	//round10000()

	//模拟的错误方法
	//recoverTest(10)
	//产生错误后 程序还能继续
	//fmt.Println("程序继续执行...")

	//slice测试
	//sliceTest()

	//map测试
	mapTest()

}

// 声明变量 不赋值
var empty string

// 定义变量 并赋值
var name = "wuta"

func varShow() {

	//定义变量 省略 var 只能定义到方法里
	novar := 998
	fmt.Println(empty, name, novar)

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
	y = 99
	fmt.Print("坐标", x, ", ", y)

	//这种分解的写法,一般用于声明全局变量
	var (
		top1 int
		topb bool
	)
	top1 = 10
	topb = true
	fmt.Print("全局变量", top1, topb)

	var p1, p2, p3 int = 1, 2, 3
	fmt.Print("指定类型 ", p1, p2, p3)

	var pi, pb, ps = 1, false, "dd"
	fmt.Print("不指定类型 ", pi, pb, ps)

	//不带声明格式的只能在函数体内实现
	g, h := 123, "ddd"
	fmt.Println(g, h)

	//实际上7的赋值被废弃，变量 _  不具备读特性
	_, value := 7, 5
	//fmt.Println(_) //_变量的是读不出来的
	fmt.Println(value) //5

	fmt.Printf("%T", value) //5

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

	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	println(a, b, c)

	//
	const (
		l0 = iota
		l1
		l2
		l3
	)

}

func plus(x int, y int) int {
	return x + y
}

func xySwap(x, y string) (string, string) {
	return y, x
}

func replace(x, y string) (r1 string, r2 string) {
	r1 = x
	r2 = y
	return
}
func replace2(x, y string) (r1, r2 string) {
	r1 = x
	r2 = y
	return
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

// 修改值
func change(p *int) {
	*p = 50
}

// 验证
func test() {
	i = 100    //定义一个变量i值为100
	change(&i) //传递i的内存地址
	fmt.Println("修改后: ", i)
}

func recoverTest(i int) {

	//定义10个元素的数组
	var arr [10]int

	//错误拦截要在产生错误前设置
	defer func() {
		//设置拦截错误信息
		err := recover()
		//产生panic异常 打印错误信息
		if err != nil {
			fmt.Println(err)
		}

	}()

	//根据函数参数为数组元素赋值
	//如果i的值超过数组下标 会报错误：数组下标越界
	arr[i] = 10
}

func sliceTest() {

	var slice4 []string
	slice4 = []string{"aaa", "b"}

	//slice4 := []int{1, 2, 3}

	for i, val := range slice4 {
		fmt.Println("第 ", i, " 个值是: ", val)
	}

	//追加
	slice5 := []string{"a", "bb"}
	//给切片新增一个值
	slice6 := append(slice5, "ddd")

	for i, val := range slice5 {
		fmt.Println("第 ", i, " 个值是: ", val)
	}

	for i, val := range slice6 {
		fmt.Println("第 ", i, " 个值是: ", val)
	}

	fmt.Println("len: ", len(slice6), " cap: ", cap(slice6))

	//截取
	slice7 := []string{"a", "b", "c", "d", "e", "f", "g"}
	//从下标 2 开始 到 下标 5
	slice8 := slice7[2:5] //c,d,e
	fmt.Println("截取: ", slice8)

	//copy
	n1 := []int{1, 2, 3}
	var n2 = make([]int, 5)
	copy(n2, n1)
	fmt.Println("拷贝: ", n2)

}

func mapTest() {

	//先声明map类型
	var map0 map[string]int
	//再开辟空间
	map0 = make(map[string]int, 5)
	map0["dd"] = 1
	map0["a"] = 2
	fmt.Println("map0: ", map0)

	//直接开辟
	map1 := make(map[int]string)
	map1[1] = "wang"
	map1[9] = "kk"
	fmt.Println("map1: ", map1)

	//简写
	map2 := map[string]string{
		"key": "value",
		"ddd": "bili",
		"bi":  "滴滴",
	}
	fmt.Println("map2: ", map2)

	//遍历 和切片类似
	for i, val := range map2 {
		fmt.Println("map2: ", i, "-", val)
	}

}
