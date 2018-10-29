
package main

import "fmt"

func main() {
	fmt.Println("hehe")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	var a int = 1;
	defer fmt.Println(a);
	a =5 ;

	fmt.Println("done")
	fmt.Println("he")
}

