
package main

import "fmt"

func swap(a, b int) (int, int){
    return b, a
}

func main(){
    var a, b int
    a = 1
    b = 2
    a, b = swap(a, b)
    fmt.Println(a)
    fmt.Println(b)
}
