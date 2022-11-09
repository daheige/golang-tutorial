package main

import "fmt"

// 定义结构体
type Demo struct {
	// 小写表示不导出,包外不能引用
	a bool
	// 大写表示导出，包外能引用
	B byte
	C int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	D float32 // float64
	E string
	F []int
	G map[string]int
}

func main() {
	d := Demo{ // 创建一个 Demo 类型的结构体
		a: true,
		B: 'b',
		C: 1,
		D: 1.0,
		E: "E",
		F: []int{1},
		G: map[string]int{"GOLANG": 1},
	}

	fmt.Printf("%+v\n", d)

	// 结构体字段使用点号来访问
	d.a = false // 修改a字段的值

	fmt.Printf("%+v\n", d)

	fmt.Printf("dome.B: %c\n", d.B)
}
