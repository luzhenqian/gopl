// 获取 URL
// 输出状态码
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) || !strings.HasPrefix(url, httpsPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%v\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
