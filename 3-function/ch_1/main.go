package main

import "fmt"

// 定义一个 print 函数, 它接受 0 个参数，返回 0 个值
func print() {
	fmt.Println("hello world")
}

// 定义一个 add 函数, 它接受两个 int 类型的参数,返回 一个 int 类型的值 （参数的声明请看下一章）
func add(x int, y int) int {
	return x + y
}

// 定义一个 swap 函数, 它接受两个 int 类型的参数,返回 两个 int 类型的值
func swap(x, y int) (int, int) {
	return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
func swap2(x, y int) (a int, b int) {
	a = y
	b = x
	return
}

// 传递必定长参数
func variableCut(x int, y ...int) int {
	for _, v := range y {
		x += v
	}
	return x
}

// 值类型参数
func modifyValue(x int) {
	x = x * 10
}

// 指针类型参数
func modifyPointer(x *int) {
	*x = (*x) * 10
}

func main() {
	print()
	fmt.Printf("add() return: %d\n", add(1, 2))
	x, y := swap(1, 2)
	fmt.Printf("swap() x: %d, y: %d \n", x, y)
	fmt.Printf("advariableCutd() return: %d\n", variableCut(1, 2, 3, 4, 5))

	// 值类型传递与指针类型传递的区别
	//		值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
	//		指针传递(引用传递)是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
	x = 1
	modifyValue(x)
	fmt.Printf("modifyValue() return: %d\n", x)

	x = 1
	modifyPointer(&x)
	fmt.Printf("modifyPointer() return: %d\n", x)
}
