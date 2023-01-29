package _struct

import (
	"fmt"
	"time"
)

type StructRunner int64

func (sr StructRunner) Run() {
	//EmployeeByID(0).Salary = 0
	//Sort([]int{5, 3, 1, 2, 4, 1})

	anonymousMembers()
}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func EmployeeByID(id int) *Employee {

	return nil
}

//// 二叉树实现插入排序
type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func Sort(values []int) {
	fmt.Printf(" 输入数据 : %v\n", values)
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Printf(" 排序结果 : %v\n", appendValues(values[:0], root))
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 匿名成员使用

type Point struct {
	X int
	Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	StructRunner
	Spokes int
	X      int
}

func anonymousMembers() {
	var w Wheel
	w.Circle.Point.X = 8
	w.Circle.Point.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	var wa Wheel
	wa.X = 8
	wa.Y = 8
	wa.Radius = 5
	wa.Spokes = 20
	fmt.Printf("%#v\n", w)

	w = Wheel{
		Circle: Circle{
			Point:  Point{8, 9},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)

}
