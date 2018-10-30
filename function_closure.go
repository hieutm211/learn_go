
package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		b = a+b
		a = b-a
		return b-a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

