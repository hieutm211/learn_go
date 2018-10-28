
package main

import (
    "fmt"
    "math"
)

func is_prime(n int) bool {
    if n < 2 {
	return false
    }

    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
	if n%i == 0 {
	    return false
	}
    }

    return true
}

func main(){
    var n int
    fmt.Scanf("%v", &n)
    if is_prime(n){
	fmt.Println(n, "is a prime!")
    } else {
	fmt.Println(n, "is not a prime!")
    }
}
