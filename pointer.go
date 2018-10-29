
package main

import "fmt"

func main(){
    var i int = 1
    var p *int = &i
    fmt.Println(*p)
    *p = 5
    fmt.Println(i)
    fmt.Println(p)
}
