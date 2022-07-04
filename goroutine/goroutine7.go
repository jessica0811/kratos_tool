/*
 * @Author: qiuhua.lin
 * @Date: 2022-07-04 08:19:07
 */
package main

import "fmt"

//管道模式 2

func Generator(max int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			fmt.Println("i: ", i)
			out <- i
		}
		close(out)
	}()
	return out
}

func Square(in <-chan int) <-chan int {
	fmt.Println("square: ", <-in)
	out := make(chan int, 100)
	go func() {
		for v := range in {
			fmt.Println("square1: ", v)
			out <- v * v
		}
		close(out)
	}()
	return out
}

func Sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var Sum int
		for v := range in {
			fmt.Println("sum1: ", v)
			Sum += v
		}
		out <- Sum
		close(out)
	}()
	return out
}

func main() {
	// 1. 生成数组
	arr := Generator(5)
	// 2. 求数组每一个元素的平方
	squ := Square(arr)
	sum := <-Sum(squ)
	fmt.Println(sum)
}
