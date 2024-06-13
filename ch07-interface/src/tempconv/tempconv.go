package tempconv

import (
	"flag"
	"fmt"
	"tempconv0"
)

// *celsiusFlag满足flag.Value接口
type celsiusFlag struct{ tempconv0.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) //无须检查错误
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv0.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv0.FToC(tempconv0.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag 定义了一个Celsius标志，返回了标志指针
// 标志必须包含一个数值和一个单位，比如“100C”
func CelsiusFlag(name string, value tempconv0.Celsius, usage string) *tempconv0.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
