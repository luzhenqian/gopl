// 打印参数索引和参数值，每组一行

package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, arg)
	}
}
