package main

import "fmt"

func main() {
	// fmt.Println("Hi, Charles")
	// var arr [4]int = [4]int{0, 1, 2, 3}
	// fmt.Printf("%d", arr[1:4])
	// var m1 map[int]string = map[int]string{}
	// m1[1] = "Tom"
	// println(m1[1])
	// m2 := map[int]string{1: "Tom", 2: "Lily"}
	// m2[0] = "John"
	// println(m2[1])
	// m3 := make(map[int]string)
	// m3[1] = "Tom"
	// println(m3[3])

	// println("开始循环遍历")
	// for k, v := range m2 {
	// 	println(k, "--->", v)
	// }

	// println("开始循环遍历key")
	// for k := range m2 {
	// 	println(k, "--->")
	// }

	// println("开始循环遍历value")
	// for _, v := range m2 {
	// 	println("--->", v)
	// }
	a := 13
	str := "new(string)"
	fmt.Printf("--%s--", str)
	fmt.Println("--", str, "--")
	fmt.Println("--", &a, "--")
	fmt.Println("--", a, "--")

	score := 91

	switch {
	case score > 90:
		fmt.Println("优秀")
		fallthrough
	case score > 80:
		fmt.Println("良好")
	case score > 60:
		fmt.Println("及格")
	case score > 50:
		fmt.Println("差")
		fallthrough
	case score > 10:
		fmt.Println("极差")
	default:
		fmt.Println("这啥成绩")
	}

	myarr := [...]string{"A", "B", "C"}
	for v := range myarr {
		fmt.Println(myarr[v])
	}
}
