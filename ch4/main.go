package main

import (
	"go_try/ch4/array"
	_map "go_try/ch4/map"
	"go_try/ch4/slice"
	_struct "go_try/ch4/struct"
)

type Runner interface {
	Run()
}

func main() {
	println("复合类型： 1. 数组； 2. Slice切片； 3. 结构体； 4. JSON； 5. 文本和HTML模板")
	var runners = [...]Runner{array.Runner(0), slice.Runner(1), _map.MapContainer(2), _struct.StructRunner(3)}
	//for _, r := range runners {
	//r.Run()
	//}

	runners[3].Run()
}
