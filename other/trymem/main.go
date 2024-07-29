package main

import "runtime"

func main() {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	println(stat.HeapSys/1024, "MB")
}
