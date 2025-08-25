package main

import (
	_ "go-study/main/study/lib/lib1"
	// mylib2 "go-study/main/study/lib/lib2"
	. "go-study/main/study/lib/lib2"
)

func main() {
	// lib1.Lib1Test()
	// lib2.Lib2Test()
	// mylib2.Lib2Test()
	Lib2Test()
}
