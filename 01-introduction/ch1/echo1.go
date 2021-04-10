// 该程序演示在控制台输出命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// 使用 var 声明变量，格式为 var 变量命 变量类型
// 变量分为声明和初始化，在声明时可以不进行初始化，那么变量会隐式赋值为该类型零值，字符串的零值是 ""
// 如果声明多个变量，使用 , 隔开

// := 是短变量声明，它会根据变量的值给变量推断出合适的类型
// for 是 go 语言唯一的循环关键字，格式为 for initialization; condition; post {}
// 与其他语言相比，for 两侧不需要小括号，但左大括号必须与 for 在同一行
// 传统的 while 等同于 for condition {}
// 死循环等同于 for {}
// 循环可以使用 return 和 coutinue 来终止

// 命令行参数使用 os.Args 变量获取，它是一个 slice 结构体
// 通俗来说，slice 是一个动态扩容的数组。
// 常见操作有通过索引取值 s[i]、获取连续子区间 s[m:n]、获取长度 len(s)
// += 是一种运算操作，等同于 s = s + sep + os.Args[i]
