/*
 * @Author: qiuhua.lin
 * @Date: 2022-07-11 22:13:43
 */
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type WorkProcessData struct { // 创建过程数据的存储容器
	Data map[string]string
	mux  sync.RWMutex
}

// 添加数据到容器
func (s *WorkProcessData) AddData(key, value string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.Data == nil { // map 未初始化值为 nil
		s.Data = make(map[string]string)
	}
	s.Data[key] = value
}

// 从容器获取数据
func (s *WorkProcessData) GetData() string {
	return s.Data["1"] + "," + s.Data["2"]
}

// 工作过程
func workProcessUnit(name string, ch chan string) {
	var wd = &WorkProcessData{}
	var group = sync.WaitGroup{}
	group.Add(2)

	go process1(&group, wd)
	go process2(&group, wd)

	group.Wait()
	ch <- name + ":" + wd.GetData()
}

// 工作过程1
func process1(group *sync.WaitGroup, gData *WorkProcessData) {
	defer group.Done()
	time.Sleep(time.Microsecond * 1)
	gData.AddData("1", strconv.Itoa(rand.Intn(10)))
}

// 工作过程2
func process2(group *sync.WaitGroup, gData *WorkProcessData) {
	defer group.Done()
	time.Sleep(time.Microsecond * 2)
	gData.AddData("2", strconv.Itoa(rand.Intn(10)))
}

// 任务生产者函数
func TaskProducer(ch chan string) {
	name := "Task"
	// 启动 8 个任务
	for i := 1; i <= 8; i++ {
		go workProcessUnit(name+strconv.Itoa(i), ch)
	}
}

// 任务消费者函数
func TaskConsumer(ch chan string, finished chan bool) {
	// 消费 8 个任务的返回值
	var result string
	i := 0
	for value := range ch {
		result += value + "\n"
		if i++; i == 8 {
			break
		}
	}
	finished <- true
	fmt.Println(result)
}

func main() {
	var ch = make(chan string, 2)
	// 结束标志
	var finished = make(chan bool)
	go TaskProducer(ch)
	go TaskConsumer(ch, finished)
	<-finished
}
