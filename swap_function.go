
package main

import "fmt"

func swap(a, b *int){
    tmp := *a
    *a = *b
    *b = tmp
}

func main(){
    var a, b int
    a = 1
    b = 2
    swap(&a, &b)
    fmt.Println(a)
    fmt.Println(b)
}
