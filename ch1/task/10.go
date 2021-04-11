// 并发 fetch
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	file, err := os.Create("./ch1/fetch_cache/" + strings.TrimSuffix(strings.TrimPrefix(url, "http://"), "/") + "_" + start.Format("2006-01-02 15:04:05"))
	if err != nil {
		ch <- fmt.Sprintf("create file: %s : %v", url, err)
		return
	}
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading: %s : %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// goroutine 是一个并发执行的函数
// 通道 chan 是一种允许某一例程向另一个例程传递指定类型的值的通信机制
// 使用 io.Copy 获取 body 的 byte 长度，并将内容丢弃到 ioutil.Discard
// 从通道接受数据时会造成阻塞
