
package main

import (
    "fmt"
)

type pair struct {
    first, second int
}

type qtype struct {
    id, first, second int
}

type fenwick struct {
    val []int 
}

func (node *fenwick) Init(n int) {
    node.val = make([]int, n+1)
}

func (node *fenwick) Update(i, x int) {
    for i = i+1; i < len(node.val); i += i & -i {
	node.val[i] += x
    }
}

func (node *fenwick) Get(i int) int {
    ans := 0
    for i = i+1; i > 0; i -= i & -i {
	ans += node.val[i]
    }
    return ans
}

func sort_qtype(q []qtype, L, H int) {
    if L >= H { return }
    i, j, mid := L, H, q[(L+H)/2].second
    
    for i <= j {
	for q[i].second < mid { i++ } 
	for q[j].second > mid { j-- }
	if i <= j {
	    if i < j {
		tmp := q[i]
		q[i] = q[j]
		q[j] = tmp
	    }
	    i++
	    j--
	}
    }
    sort_qtype(q, L, j)
    sort_qtype(q, i, H)
}

func sort_pair(p []pair, L, H int) {
    if L >= H { return }
    i, j, mid := L, H, p[(L+H)/2].first
    for i <= j {
	for p[i].first < mid { i++ }
	for p[j].first > mid { j-- }
	if i <= j {
	    if i < j {
		tmp := p[i]
		p[i] = p[j]
		p[j] = tmp
	    }
	    i++
	    j--
	}
    }
    sort_pair(p, L, j)
    sort_pair(p, i, H)
}

var a, ans []int
var query []qtype

func compress(a []int) {
    b := make([]pair, len(a))
    for i := range a {
	b[i].first = a[i]
	b[i].second = i
    }

    sort_pair(b, 0, len(b)-1)

    cnt := 1
    a[b[0].second] = cnt 
    for i := 1; i < len(a); i++ {
	if b[i-1].first != b[i].first { cnt++ }
	a[b[i].second] = cnt
    }
}

func main(){
    var n, u, v, id int
    var fw fenwick
    var prev [1000001]int

    fmt.Scanf("%d", &n)
    a = make([]int, n)
    fw.Init(n)

    for i := 0; i < len(a); i++ {
	fmt.Scanf("%v", &a[i])
    }
    fmt.Scanf("\n")
    fmt.Scanf("%d", &n)

    query = make([]qtype, n)
    ans = make([]int, n)

    for i := 0; i < len(query); i++ {
	query[i].id = i
	fmt.Scanf("%v%v%v", &query[i].first, &query[i].second)
	query[i].first--
	query[i].second--
    }


    compress(a)

    sort_qtype(query, 0, len(query)-1)

    for i := range prev {
	prev[i] = -1
    }

    i := 0
    for _, x := range query {
	u = x.first
	v = x.second
	id = x.id

	for ; i <= v; i++ {
	    if prev[a[i]] != -1 {
		fw.Update(prev[a[i]], -1)
	    }
	    prev[a[i]] = i
	    fw.Update(i, +1)
	}

	ans[id] = fw.Get(v) - fw.Get(u-1)
    }

    for _, x := range ans {
	fmt.Println(x)
    }
}
