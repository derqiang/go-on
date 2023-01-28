package slice

import "fmt"

type Runner int64

func (cur Runner) Run() {
	println(">>> Chapter 4.2 Array P_123 ~ P_134 >>>")
	println("StartIndex" + fmt.Sprintf("%v", cur))

	fmt.Println("")
	months := [...]string{1: "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December"}
	fmt.Printf("Slice 基本使用 (13 - 1， 取不到13，左闭右开): months[1:13] = %v \n", months[1:13])

	// 北方的夏天
	summer := months[6:9]
	// 第二季度
	Q2 := months[4:7]
	fmt.Println(Q2)
	fmt.Println(summer)
	checkOverlap(Q2, summer)

	// 超出 lens会自动扩展新slice； 超出容量，则会panic
	// fmt.Println(summer[:20])

	endlessSummer := summer[:5]
	fmt.Printf("endlessSummer : %v\n", endlessSummer)

	arr := []int{1, 2, 3, 4, 5}
	reverse(arr[:2])
	reverse(arr[2:])
	reverse(arr)

	var aPtr *[3]int
	var aPtr2 = [3]int{1, 2, 3}
	aPtr = &aPtr2
	fmt.Printf("arr 数组地址: %p\n%p\n", &aPtr2, aPtr)
	reverseArray(aPtr)

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
	//fmt.Printf(" Q2 VS summer : %v\n", &Q2[2] == &summer[0])
	println("<<< Chapter 4.2 Array P_123 ~ P_134 <<<")
}

func checkOverlap(s1 []string, s2 []string) {
	for _, r1 := range s1 {
		for _, r2 := range s2 {
			if r1 == r2 {
				fmt.Printf("%s apperars in both\n", r1)
			}
		}
	}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	//fmt.Printf("i : %v, j : %v\n", arr[i], arr[j])
	fmt.Printf("reverse : %v\n", arr)
}

func reverseArray(arr *[3]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Printf("reverseArray : %v\n", *arr)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < len(x)*2 {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
