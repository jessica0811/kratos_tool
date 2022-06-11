/*
 * @Author: qiuhua.lin
 * @Date: 2022-06-11 10:26:47
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BarrierResponse struct {
	Err    error
	Resp   string
	Status int
}

func doRequest(out chan<- BarrierResponse, url string) {
	fmt.Println("out: ", out)
	res := BarrierResponse{}
	// 设置 http 客户端
	client := http.Client{
		Timeout: time.Duration(20 * time.Second),
	}
	// 执行 get 请求
	resp, err := client.Get(url)
	if resp != nil {
		res.Status = resp.StatusCode
	}
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	byt, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	// 将获取的结果数据放入通道
	res.Resp = string(byt)
	out <- res
}

func Berrier(urls ...string) {
	requestNum := len(urls)
	in := make(chan BarrierResponse, requestNum)
	response := make([]BarrierResponse, requestNum)
	defer close(in)
	for i, urls := range urls {
		fmt.Println(i, in, urls)
		go doRequest(in, urls)
	}
	var hasError bool
	for i := 0; i < requestNum; i++ {
		fmt.Println(i)
		resp := <-in
		if resp.Err != nil {
			fmt.Println("error: ", resp.Err, resp.Status)
			hasError = true
		}
		response[i] = resp
		fmt.Println("------")
	}
	if !hasError {
		for _, resp := range response {
			fmt.Println(resp.Status)

		}
	}
}

func main() {
	Berrier([]string{"https://www.baidu.com", "https://www.weibo.com", "https://www.shirdon.com"}...)
}
