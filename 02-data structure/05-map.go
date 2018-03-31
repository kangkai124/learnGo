package main

import "fmt"

/*
 * 创建：make(map[string]int)
 * 获取元素：m[key]
 * key不存在，获取value初始值，即zero value
 * value, ok := m[key]判断是否存在key
 * delete删除一个key
 * range遍历map
 * 不保证遍历顺序，如需顺序，需手动对key排序
 * 使用len获取元素个数
 *
 * map的key
 * map使用哈希表，必须可以比较相等
 * 除了slice,map,function的内奸类型，都可以作为key
 */

func main() {
	m := map[string] string {
		"name": "kk",
		"course": "go",
		"company": "k2data",
	}

	m2 := make(map[string] int)	 // m2 == empty map

	var m3 map[string] int  // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v :=range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	causeName := m["cause"]
	fmt.Println(causeName)	// zero value

	name , ok := m["name"]
	fmt.Println(name, ok)
	wrongName, ok := m["nam"]
	fmt.Println(wrongName, ok)

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	delete(m, "name")
	name , ok = m["name"]
	fmt.Println(name, ok)
}
