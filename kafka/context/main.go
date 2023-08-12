package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.baidu.com/", nil)
	if err != nil {
		fmt.Println("http request failed,err :", err)
		return
	}
	go func() {
		res, err := client.Do(req)
		pack := Result{res, err}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		fmt.Println("timeout", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Println("server response is :", string(out))
	}
	return

}
func main() {
	go process()
	time.Sleep(2 * time.Second)

}
