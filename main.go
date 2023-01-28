package main

import (
	"fmt"
	tempconv2 "go_try/ch2/cf/tempconv"
)

func main() {

	// Chapter 1 , P3
	//ch1.BenchTest()

	// 字符串递归全排
	//me.PermCom()

	// Chapter 1.3, P6
	//ch1.ChapterFirstToThree()

	// Chapter 1.3, P30
	//ch1.ChapterDup2()

	// Chapter 1.3, P33
	//ch1.ChapterDup3()

	// Chapter 1.4, P35
	//ch1.ChapterLissajous()

	// Chapter 1.5, P38
	//ch1.Fetch()
	//ch1.FetchAll()

	//ch1.Server1()

	/* Chapter 2 */
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g °F = %g °C \n", freezingF, tempconv2.FToC(freezingF))
	fmt.Printf("%g °F = %g °C \n", boilingF, tempconv2.FToC(boilingF))
	fmt.Printf("Brrr! %v\n", tempconv2.AbsoluteZeroC)
	tempconv2.Run()
	//
}
