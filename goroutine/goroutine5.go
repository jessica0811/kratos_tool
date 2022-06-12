/*
 * @Author: qiuhua.lin
 * @Date: 2022-06-12 08:05:25
 */
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Publisher struct {
	subscribers map[Subscriber]TopicFunc
	buffer      int           // 订阅者缓冲区长度
	timeout     time.Duration // publisher 发送消息的超时时间
	m           sync.RWMutex
	/*
		m 用来保护 subscribers,
		当修改 subscribers 的时候（即新加订阅者或删除订阅者）使用写锁，
		当向某个订阅者发送消息的时候（即向某个 Subscriber channel 中写入数据），使用读锁
	*/
}

// 订阅者通道
type Subscriber chan interface{}

// 主题函数
type TopicFunc func(v interface{}) bool

//实例化
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[Subscriber]TopicFunc),
	}
}

// 发布订阅者方法
func (p *Publisher) Subscribe() Subscriber {
	return p.SubscribeTopic(nil)
}

// 发布者订阅主题
func (p *Publisher) SubscribeTopic(topic TopicFunc) Subscriber {
	ch := make(Subscriber, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 删除某个订阅者
func (p *Publisher) Delete(sub Subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub) // map 的删除
	close(sub)
}

// 发布者发布消息
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭 Publisher, 删除所有订阅者
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// 发送主题
func (p *Publisher) sendTopic(sub Subscriber, topic TopicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("v: ", v)
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func main() {
	// 实例化
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()
	// 订阅者订阅所有消息
	all := p.Subscribe()
	fmt.Println("all: ", len(all))
	// 订阅者仅订阅包含 golang 的消息
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	// 发布消息
	p.Publish("hello, world!")
	p.Publish("hello, golang!")
	// 加锁
	var wg sync.WaitGroup
	wg.Add(2)
	// 开启 goroutine
	go func() {
		for msg := range all {
			fmt.Println("msg: ", msg)
			_, ok := msg.(string)
			fmt.Println(ok)
		}
		wg.Done()
	}()
	// 开启 goroutine
	go func() {
		for msg := range golang {
			v, ok := msg.(string)
			fmt.Println(v)
			fmt.Println(ok)
		}
		wg.Done()
	}()
	p.Close()
	wg.Wait()
}
