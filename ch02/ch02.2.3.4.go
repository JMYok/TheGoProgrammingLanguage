package main

func main() {
	println(func1() == func1())
	println(func2() == func2())
}

func func1() *int64 {
	var dummy int64 = 999
	return &dummy
}

func func2() *int64 {
	return new(int64)
}

/*
func func3() *[0]int {
	return &[0]int
}
*/
