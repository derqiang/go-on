package me

import (
	"fmt"
	"sort"
	"strings"
)

var global = make([][]string, 0)

func PermCom() {
	// Test 1
	str := "a"
	permutationAndCombination([]rune(preSort(str)), make([]string, 0))
	fmt.Printf("Test1 Input : %v >> Count : %v , Values : %v\n", str, len(global), global)
	global = global[:0]
	// Test 2
	str = "BaC"
	permutationAndCombination([]rune(preSort(str)), make([]string, 0))
	fmt.Printf("Test2 Input : %v >> Count : %v , Values : %v\n", str, len(global), global)
	global = global[:0]
	// Test 3
	str = "AcDD"
	permutationAndCombination([]rune(preSort(str)), make([]string, 0))
	fmt.Printf("Test3 Input : %v >> Count : %v , Values : %v\n", str, len(global), global)
	global = global[:0]
	// Test 4
	str = "aaedc"
	permutationAndCombination([]rune(preSort(str)), make([]string, 0))
	fmt.Printf("Test4 Input : %v >> Count : %v , Values : %v\n", str, len(global), global)
}

/**
明确知道递归程序设计要点：
① 任务可分
② 分后可重复进行
③ 明确退出时机
④ 递归路径上，适当位置，完成目标逻辑
*/
func permutationAndCombination(arr []rune, collector []string) {
	if len(arr) == 0 {
		global = append(global, collector)
		return
	}

	for i, r := range arr {
		lc := make([]string, len(collector))
		copy(lc, collector)
		lc = append(lc, string([]rune{r}))

		tmp1 := make([]rune, len(arr[0:i]))
		copy(tmp1, arr[0:i])
		tmp2 := make([]rune, len(arr[i+1:]))
		copy(tmp2, arr[i+1:])
		tmp1 = append(tmp1, tmp2...)
		permutationAndCombination(tmp1, lc)
	}
}

func preSort(initStr string) string {
	arr := strings.Split(initStr, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}
