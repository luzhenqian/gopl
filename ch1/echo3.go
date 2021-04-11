package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// 使用 strings.Join 是一种性能更强的方法。因为内部使用了 buffer 的方式对字符串进行拼接。
