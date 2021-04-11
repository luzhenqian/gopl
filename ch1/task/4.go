// 记录命令行参数中的文件内容中出现的每一行的次数

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	type CountInfo struct {
		fileNames []string
		count     int
	}
	counts := make(map[string]*CountInfo)
	files := os.Args[1:]
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			temp := strings.Split(file, "/")
			fileName := temp[len(temp)-1]
			if _, ok := counts[line]; ok {
				counts[line].count++
				exist := false
				for _, name := range counts[line].fileNames {
					if name == fileName {
						exist = true
						continue
					}
				}
				if !exist {
					counts[line].fileNames = append(counts[line].fileNames, fileName)
				}
			} else {
				counts[line] = &CountInfo{
					count:     1,
					fileNames: []string{fileName},
				}
			}
		}
	}

	for line, n := range counts {
		fmt.Printf("fileName:%s\t%d\t%s\n", strings.Join(n.fileNames, " "), n.count, line)
	}
}
