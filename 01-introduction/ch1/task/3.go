// 测试三种实现之间的性能差异

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	echo1()
	elapsed1 := time.Since(t1)
	fmt.Println("echo1 elapsed: ", elapsed1)

	t2 := time.Now()
	echo2()
	elapsed2 := time.Since(t2)
	fmt.Println("echo2 elapsed: ", elapsed2)

	t3 := time.Now()
	echo3()
	elapsed3 := time.Since(t3)
	fmt.Println("echo3 elapsed: ", elapsed3)
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
