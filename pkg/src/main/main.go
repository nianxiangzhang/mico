package main

import (
	"fmt"
	"reflect"
)

func main() {

	//testSlice()
	//TestTag()
	//DoOP(1, increase)
	TestMakeNew()

}

/**
New 返回指针地址
Make返回第一个元素，可预设内存空间，避免未来的内存拷贝
*/

func TestMakeNew() {

	n1 := new([]int)
	n2 := make([]int, 0)
	n3 := make([]int, 10)
	n4 := make([]int, 10, 20)
	println(n1)
	println(n2)
	println(n3)
	println(n4)

}

func TestClosePackage() {
	defer func() {
		if r := recover(); r != nil {
			println("recovered ")
		}
	}()
}

func DoOP(x int, f func(int, int)) {
	f(x, 1)
}

func increase(a, b int) {
	println("increase result is:", a+b)
}

func decrease(a, b int) {
	println("decrease result is:", a-b)
}

type MyType struct {
	Name string `json:"name"`
}

func TestTag() {
	mt := MyType{Name: "test"}
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")
	fmt.Println(tag)
}

// go 语言全是值传递
func testSlice() {
	mySlice := []int{10, 20, 30, 40, 50}
	// 这里相当于创建了一个临时变量value，将value临时变量*2，不会影响切片的值
	for _, value := range mySlice {
		value *= 2
		fmt.Println(value)
	}
	fmt.Printf("slice: %+v\n", mySlice)

	for index, _ := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("slice: %+v\n", mySlice)
}
