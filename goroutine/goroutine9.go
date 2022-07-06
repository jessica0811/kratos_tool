/*
 * @Author: qiuhua.lin
 * @Date: 2022-07-06 08:16:52
 */
package main

import (
	"fmt"
	"sync"
)

// 管道模式 1

func Buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			fmt.Println("配件: ", i)
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

func Build(in <-chan string) <-chan string {
	fmt.Println("1111")
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range in {
			fmt.Println("组装：", v)
			out <- fmt.Sprint("组装", v)
		}
	}()
	return out
}

func Pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			fmt.Println("打包: ", c)
			out <- fmt.Sprint("打包", c)
		}
	}()
	return out
}

func Merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	// 把一个通道中的数据发送到 out 中
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	// 扇入，需要启动多个 goroutine 用于处理多个通道中的数据
	for _, cs := range ins {
		go p(cs)
	}
	// 等待所有输入的数据 ins 处理完，再关闭输出 out
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// 采购 12 套配件
	accessories := Buy(12)
	// 组装 12 台计算机
	fmt.Println("1111")
	computers1 := Build(accessories)
	fmt.Println("2222")
	computers2 := Build(accessories)
	fmt.Println("3333")
	computers3 := Build(accessories)
	fmt.Println("4444")
	// 汇聚 3 个通道成一个
	computers := Merge(computers1, computers2, computers3)
	// 打包他们以便售卖
	fmt.Println("5555")
	packs := Pack(computers)
	fmt.Println("6666")
	// 输出测试
	for p := range packs {
		fmt.Println(p)
	}
}
