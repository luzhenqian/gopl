// 这是一个从命令行输出 hello,world. 的例子。

package main

import "fmt"

func main() {
	fmt.Println("hello,world.")
}

// package 表示文件属于哪个包
// 用 main 来命名包是特殊的，它属于可独立执行的程序，而非代码库。

// import 表示导入的包列表
// import 不可以导入未使用的包
// import 必须在 package 关键字的后面

// package 和 import 后面可以是函数（func）、变量（var）、常量（const）、类型（type），一般来说它们的顺序无所谓。

// 用 main 为函数命名也是特殊的，表示程序的入口函数。
// 函数由 func 关键字、函数名、函数参数列表（可省略）、返回值参数列表（可省略）和函数体（可为空） 5 部分组成。

// 函数体的 { 必须和函数名在同一行
// 语句不需要分号结尾
