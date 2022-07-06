/*
 * @Author: qiuhua.lin
 * @Date: 2022-07-04 08:30:29
 */
package main

import "fmt"

// 管道模式 1

func Buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

func Build(in <-chan string) <-chan string {
	fmt.Println("00000")
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range in {
			fmt.Println(v)
			out <- fmt.Sprint("组装", v)
		}
	}()
	return out
}

func Pack(in <-chan string) <-chan string {
	fmt.Println("=======")
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- fmt.Sprint("打包", c)
		}
	}()
	return out
}

func main() {
	accessories := Buy(6)
	fmt.Println("11111")
	computers := Build(accessories)
	fmt.Println("22222")
	packs := Pack(computers)
	fmt.Println("33333")
	for p := range packs {
		fmt.Println(p)
	}
}
