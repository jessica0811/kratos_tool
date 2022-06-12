// /*
//  * @Author: qiuhua.lin
//  * @Date: 2022-06-11 15:13:18
//  */
// package main

// import (
// 	"fmt"
// )

// /*
// 在日常编程中，可能会遇到这样一个场景，一个任务可能有好几件事需要去做，并且这些事是完全可以并发执行的，
// 除此之外，我们还需要得到其执行结束后的结果，并以此来进行后续的一些处理。
// 这个时候就可以考虑去使用Go编程当中的Future模式
// */

// type FutureTask struct {
// 	args chan interface{}
// 	res  chan interface{}
// }

// func execFutureTask(futureTask *FutureTask) {
// 	// 读取传入的参数
// 	fmt.Println("goroutine 读取到的参数: ", <-futureTask.args)
// 	// 这里可以执行具体的业务逻辑
// 	result := "执行完业务逻辑后得到的结果"
// 	// 将结果进行保存
// 	futureTask.res <- result
// }

// func main() {
// 	args := []string{"main 线程传入的参数", "helloworld", "111", "222", "333"}
// 	futureTask := FutureTask{make(chan interface{}, len(args)), make(chan interface{}, len(args))}
// 	defer close(futureTask.args)
// 	defer close(futureTask.res)
// 	// 读取线程执行的
// 	for i := 0; i < len(args); i++ {
// 		go execFutureTask(&futureTask)
// 		// 向FutureTask传入参数，如果不传的话会死锁
// 		futureTask.args <- args[i]
// 	}
// 	for i := 0; i < len(args); i++ {
// 		fmt.Println("主线程读取 future 模式下goroutine的结果: ", <-futureTask.res)
// 	}
// }
