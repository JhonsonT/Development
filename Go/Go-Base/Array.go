package main

import "fmt"

// 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
//java 中 数组是引用类型

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArr2(arr [5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	UserInfo := [...]struct {
		name string
		age  int
	}{
		{"gyf", 18},
		{"tjs", 18},
		{"syw", 18},
		{"mkk", 18},
	}

	fmt.Println(UserInfo)
	fmt.Println(UserInfo[0].name)
	fmt.Println(UserInfo[1].name)
	fmt.Println(&UserInfo[1].name)
	fmt.Println(&UserInfo[2].name)

	var arr1 [5]int
	fmt.Println("source", arr1)

	printArr(&arr1)
	fmt.Println(arr1)

	arr2 := [...]int{2, 4, 6, 8, 10}
	fmt.Println("source", arr2)

	printArr2(arr2)
	fmt.Println(arr2)

	printArr(&arr2)
	fmt.Println(arr2)

	var slice0 = [5]int{0, 1, 2, 3, 4}
	fmt.Println(slice0)

	var slice1 = make([]int, 5, 10)
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))

	for i := 5; i < 10; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))

	for i := 10; i < 15; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))

	for i := 15; i < 20; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))

	for i := 20; i < 25; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))
}
