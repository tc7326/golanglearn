package main

import "fmt"

func init() {
	defer fmt.Println("hello2 init")
}

func hello2() {
	defers()
	slices()
	maps()
}

// 延迟函数 类似于栈 先进后出LIFO
func defers() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
}

// 切片 slice（动态数组）
func slices() {

	//声明一个定义切片（未指定大小的数组来）
	var nums []int
	//使用make()创建切片 并定义初始长度
	nums = make([]int, 66)
	fmt.Println(nums, len(nums), cap(nums))

	nums2 := make([]string, 188)
	nums2[0] = "这是数组"
	nums2[187] = "这是数组2"
	fmt.Println(nums2, len(nums2), cap(nums2))

}

// map map[key]value
func maps() {
	var maps map[int]string
	maps = make(map[int]string, 10)
	maps[0] = "巫师III"
	maps[1] = "行星边际2"
	maps[2] = "壁纸引擎"
	fmt.Println(maps, len(maps))

	games := map[string]string{
		"MC":    "我的世界",
		"PS2":   "行星边际2",
		"CS:GO": "反恐精英-全球攻势",
	}
	fmt.Println(games, len(games))

	//key是string value是map 的 map
	works := make(map[string]map[string]string)
	works["android"] = map[string]string{"name": "Android开发", "price": "20个W"}
	works["go"] = make(map[string]string, 6)
	works["go"]["title"] = "GO语言开发"
	works["go"]["prices"] = "30个W"
	works["go"]["year"] = "5年"

	fmt.Println(works, len(works))
}
