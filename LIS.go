package main

import "fmt"

const maxN = 300001
var n int
var a = make([]int, 0, maxN)

type fenwick struct {
    val []int
}

func (fw *fenwick) Init(n int){
    fw.val = make([]int, n)
}

func (fw *fenwick) Update(i, x int){
    for ; i < len(fw.val); i += i &-i {
	if fw.val[i] < x {
	    fw.val[i] = x
	}
    }
}

func (fw *fenwick) Get(i int) int {
    ans := 0
    for ; i > 0; i -= i & -i {
	if ans < fw.val[i] {
	    ans = fw.val[i]
	}
    }
    return ans
}

func qsort(a []int, L, H int){
    if L >= H-1 { return }
    i, j, mid := L, H-1, a[(L+H-1)/2]
    for i <= j {
	for a[i] < mid { i++ }
	for a[j] > mid { j-- }
	if i <= j {
	    if (i < j){
		tmp := a[i]
		a[i] = a[j]
		a[j] = tmp
	    }
	    i++
	    j--
	}
    }
    qsort(a, L, j+1)
    qsort(a, i, H)
}

func compress(a []int){
    m := make(map[int]int)
    var cnt int

    for i := range a {
	m[a[i]] = 0
    }

    keys := make([]int, len(m))
    cnt = 0
    for i := range m {
	keys[cnt] = i
	cnt++
    }
    qsort(keys, 0, len(keys))

    cnt = 1
    for i := range keys {
	m[keys[i]] = cnt
	cnt++
    }

    for i := range a {
	a[i] = m[a[i]]
    }
}

func main(){
    fmt.Scanf("%v", &n)

    a = a[:n]

    for i := range a {
	fmt.Scanf("%v", &a[i])
    }

    compress(a)

    var fw fenwick
    fw.Init(n+2)

    ans := 0
    for i := range a {
	f := fw.Get(a[i]-1) + 1
	if f > ans {
	    ans = f
	}
	fw.Update(a[i], f)
    }
    fmt.Println(ans)
}
