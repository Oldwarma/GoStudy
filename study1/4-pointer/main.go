package main

import "fmt"

//一般不用var 声明指针，会因内存问题报错
//* 可以获取指针的，此操作成为指针的解引用， * 也叫取值操作符
//& 可以获取非指针变量的地址，叫取地址操作符。
//sp:=new([]int)指针
func main() {

	n := 10

	getPointer(&n)

	//i := -10
	//j := 25
	//pI := &i
	//pJ := &j
	//fmt.Println("pI memory:", pI)
	//fmt.Println("pJ memory:", pJ)
	//fmt.Println("pI value:", *pI)
	//fmt.Println("pJ memory:", *pJ)
}
func getPointer(n *int) {
	fmt.Println(" ", n)
	fmt.Println("*", *n)
	fmt.Println("&", &n)
	fmt.Println("&*", &*n)
	fmt.Println("*&", *&n)
	fmt.Println("**&", **&n)
	fmt.Println("*&*", *&*n)           //
	fmt.Println("*&*&*&*&", *&*&*&*&n) //a8
}
func returnPointer(n int) *int {
	v := n * n
	fmt.Println(&v)
	return &v
}
