
package main

import "fmt"

func qsort(arr []int, L, H int){
    if L >= H { return }
    i, j, mid := L, H, arr[(L+H)/2]
    for i <= j {
	for arr[i] < mid { i++ }
	for arr[j] > mid { j-- }
	if i <= j {
	    if i < j {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
	    }
	    i++
	    j--
	}
    }
    qsort(arr, L, j)
    qsort(arr, i, H)
}

func compress(arr []int){
    m := make(map[int]int)
    for i := range arr {
	m[arr[i]] = 0
    }

    keys := make([]int, len(m))
    cnt := 0
    for i := range m {
	keys[cnt] = i
	cnt++
    }
    qsort(keys, 0, len(keys)-1)

    cnt = 0
    for _, v := range keys {
	m[v] = cnt
	cnt++
    }

    for i := range arr {
	arr[i] = m[arr[i]]
    }
}

func main(){
    var n int

    fmt.Scanf("%d", &n)

    arr := make([]int, n)
    for i := range arr {
	fmt.Scanf("%d", &arr[i])
    }

    compress(arr)

    for _, v := range arr {
	fmt.Printf("%d ", v)
    }
    fmt.Println()
}
