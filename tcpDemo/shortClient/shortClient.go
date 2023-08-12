package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 关闭长连接,使用短连接进行http请求（请求不频繁的情况,一天只使用几次）
// 如果是复用连接,可以使用true来使用长连接
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func main() {
	date := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:8088")
	date.Set("name", "sb")
	date.Set("age", "111")
	queryStr := date.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	res, err := http.NewRequest("get", urlObj.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w, err := client.Do(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer w.Body.Close()
	byt, err := io.ReadAll(w.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byt))
}
