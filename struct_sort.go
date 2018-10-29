
package main

import "fmt"

type pair struct{
    first, second int    
}

func swap(a, b *pair){
    tmp := *a
    *a = *b
    *b = tmp
}

func qsort(a []pair, L, H int){
    if L >= H { return }
    var i, j, mid int = L, H, a[(L+H)/2].first
    for i <= j {
	for a[i].first < mid { i++ }
	for a[j].first > mid { j-- }
	if i <= j {
	    if i < j {
		swap(&a[i], &a[j])
	    }	    
	    i++
	    j--
	}
    }
    qsort(a, L, j) 
    qsort(a, i, H)
}

func main(){
    var n int

    fmt.Scanf("%d", &n)

    var a []pair = make([]pair, n, 100005)

    for i := 0; i < n; i++ {
	fmt.Scanf("%d", &a[i].first)	
	a[i].second = i
    }
    
    qsort(a, 0, n-1)
    for i := 0; i < n; i++ {
	fmt.Println(a[i])	
    }
}
