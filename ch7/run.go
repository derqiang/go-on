package ch7

import (
	"flag"
	"fmt"
	"go_try/ch2/cf/tempconv"
	"go_try/me"
)

//var period = flag.Duration("period", 1*time.Second, "sleep period")

var temp = CelsiusFlag("temp", 20.0, "the temperature")

type celsiusFlag struct {
	tempconv.Celsius
}

func (ff *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		ff.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		ff.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

type ch7 int64

func (ff ch7) Run() {
	flag.Parse()
	//flag.Parse()
	//fmt.Printf("Sleeping for %v...\n", *period)
	//time.Sleep(*period)
	//fmt.Println()
	fmt.Println(*temp)
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func GetRunner() []me.Runner {
	return []me.Runner{
		//ch7(1),
		//InterfaceValue(2),
		//SortRunner(3),
		//ServerRunner(4),
		//Server2Runner(5),
		TypeAssertRunner(6),
	}
}
