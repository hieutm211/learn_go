
package main

import (
    "fmt"
    "io"
    "os"
    "bufio"
)

const readPortionSize = 1024
const readerSize = readPortionSize<<3

func openFile(path string) *os.File {
    file, err := os.Open(path)
    if err != nil {
	panic(err)
    }
    return file
}

func main(){
    file := openFile("test.txt")
    defer file.Close()

    reader := bufio.NewReaderSize(file, readerSize)

    var buf = make([]byte, readPortionSize)
    var tmp string = ""
    var arr = make([]string, 0)

    for {
	nBytes, e := reader.Read(buf); 

	if e != nil && e != io.EOF {
	    panic(e)
	} else if e == io.EOF {
	    break
	}

	for i := 0; i < nBytes; i++ {
	    if '0' <= buf[i] && buf[i] <= '9' {
		tmp += string(buf[i])
	    } else if tmp != "" {
		arr = append(arr, tmp)
		tmp = ""
	    }
	}
    }

    fmt.Println(arr)
}
