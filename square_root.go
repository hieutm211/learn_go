
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func abs(x float64) float64 {
    if x < 0 {
	return -x	
    } else {
	return x	
    }
}

func random(L, R int) int {
    return L + rand.Intn(R-L);
}

func Sqrt(x float64) float64 {
    var z float64 

    if x < 0 { return -1.0 }
    if x < 3 {
	z = float64(1)
    } else {
	z = float64(random(int(x)/3, int(x) - int(x)/3))
    }

    prev := -1.0

    for loop := 1; loop <= 100; loop++ {
	if abs(z*z - x) <= 0.000000001 { break }
	if z == prev { break }

	if loop % 2 != 0 {
	    prev = z
	}

	z -= (z*z-x)/(2*z)
    }

    if float64(int(z)*int(z)) == x {
	return float64(int(z))	
    } else {
	return z
    }
}

func main(){
    rand.Seed(time.Now().UTC().UnixNano());    

    var n float64

    fmt.Scanf("%v", &n) 
    fmt.Printf("%v", Sqrt(n))
}
