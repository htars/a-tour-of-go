package main

import "fmt"

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	
	return func() int {
		tmp := f0
		f0 = f1
		f1 = f0 + tmp
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
