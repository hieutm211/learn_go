
package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
)

var n int

type segmentTree struct {
    val, lazy []int
}

func max(a, b int) int {
    if a >= b {
	return a
    }
    return b
}

func (node *segmentTree) Init(n int) {
    node.val = make([]int, 4*n+1)
    node.lazy = make([]int, 4*n+1)
}

func (node *segmentTree) down(k int) {
    node.val[2*k] += node.lazy[k]
    node.lazy[2*k] += node.lazy[k]
    node.val[2*k+1] += node.lazy[k]
    node.lazy[2*k+1] += node.lazy[k]
    node.lazy[k] = 0
}

func (node *segmentTree) Update(k, l, r, u, v, x int) {
    if v < l || r < u { return }
    if u <= l && r <= v {
	node.val[k] += x
	node.lazy[k] += x
	return
    }

    node.down(k)

    mid := (l+r)/2
    node.Update(2*k, l, mid, u, v, x)
    node.Update(2*k+1, mid+1, r, u, v, x)

    node.val[k] = max(node.val[2*k], node.val[2*k+1])
}

func (node *segmentTree) Get(k, l, r, u, v int) int {
    if v < l || r < u { return 0 }
    if u <= l && r <= v {
	return node.val[k]
    }

    node.down(k)

    mid := (l+r)/2

    return max(node.Get(2*k, l, mid, u, v), node.Get(2*k+1, mid+1, r, u, v))
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
	return ""
    }
    return string(str)
}

func main(){
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024) 
    writer := bufio.NewWriter(os.Stdout)

    var m, c, u, v, x int
    var st segmentTree
    var temp string

    temp = readLine(reader)
    fmt.Sscanf(temp, "%d%d", &n, &m)

    st.Init(n)

    for ; m > 0; m-- {
	temp = readLine(reader)

	fmt.Sscanf(temp, "%d%d%d%d", &c, &u, &v, &x)
	if c == 0 {
	    st.Update(1, 1, n, u, v, x)
	} else {
	    fmt.Fprintf(writer, "%d\n", st.Get(1, 1, n, u, v))
	}

	if m%1000 == 0 {
	    writer.Flush()
	}
    }
    writer.Flush()
}

