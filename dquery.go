
package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
    "strconv"
    "strings"
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

func readInput() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    var n int
    var nTemp int64
    var err error

    nTemp, err = strconv.ParseInt(readLine(reader), 10, 64)
    checkErr(err)

    n = int(nTemp)
    a = make([]int, n)
    
    arrTemp := strings.Split(readLine(reader), " ")

    for i := 0; i < len(a); i++ {
	arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	checkErr(err)

	a[i] = int(arrItemTemp)
    }

    nTemp, err = strconv.ParseInt(readLine(reader), 10, 64)
    checkErr(err)

    n = int(nTemp)

    query = make([]qtype, n)
    ans = make([]int, n)
    
    for i := 0; i < len(query); i++ {
	pairTemp := strings.Split(readLine(reader), " ")
	checkErr(err)

	query[i].id = i
	var temp int64
	temp, err = strconv.ParseInt(pairTemp[0], 10, 64)
	query[i].first = int(temp) -1
	temp, err = strconv.ParseInt(pairTemp[1], 10, 64)
	query[i].second = int(temp) -1
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
	return ""
    }
    return strings.TrimRight(string(str), "\r\n")
}

func checkErr(err error) {
    if err != nil {
	panic(err)
    }
}

func writeOutput() {
    out := bufio.NewWriter(os.Stdout)

    for i := range ans {
	fmt.Fprintf(out, "%d\n", ans[i])
	if i%1000 == 0 {
	    out.Flush()
	}
    }
    out.Flush()
}

func main(){
    var u, v, id int
    var fw fenwick
    var prev [1000001]int

    readInput()

    compress(a)

    sort_qtype(query, 0, len(query)-1)

    fw.Init(len(a))

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

    writeOutput()

}
