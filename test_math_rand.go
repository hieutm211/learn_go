
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main(){
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Random1 = ", rand.Intn(10))
}
