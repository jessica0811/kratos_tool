/*
 * @Author: qiuhua.lin
 * @Date: 2022-07-10 18:16:26
 */
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("开始倒计时.....")
	tick := time.Tick(1 * time.Second)
	for coundown := 5; coundown > 0; coundown-- {
		fmt.Println(coundown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("abort...!")
			return
		}
	}
}
