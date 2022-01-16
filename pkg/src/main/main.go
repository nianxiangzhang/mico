package main

import (
	"fmt"
	"reflect"
)

func main() {

	//testSlice()
	TestTag()

}

type MyType struct {
	Name string `json:"name"`
}

func TestTag() {
	mt := MyType{Name: "test"}
	myType :=reflect.TypeOf(mt)
	name :=myType.Field(0)
	tag :=name.Tag.Get("json")
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
