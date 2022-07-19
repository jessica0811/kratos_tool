/*
 * @Author: qiuhua.lin
 * @Date: 2022-07-19 09:18:39
 */
package main

import (
	"errors"
	"fmt"
	"sync"
)

type RoundRobinBalancer struct {
	m     sync.Mutex
	next  int
	items []interface{}
}

func New(items []interface{}) *RoundRobinBalancer {
	return &RoundRobinBalancer{items: items}
}

func (b *RoundRobinBalancer) Pick() (interface{}, error) {
	if len(b.items) == 0 {
		return nil, errors.New("没有可用项")
	}
	b.m.Lock()
	r := b.items[b.next]
	b.next = (b.next + 1) % len(b.items)
	fmt.Println("next: ", b.next)
	b.m.Unlock()
	return r, nil
}

func main() {
	source := []interface{}{"10.0.0.1", "10.0.0.2", "10.0.0.3"}
	b := New(source)
	wc := sync.WaitGroup{}
	for i := 0; i < 8; i++ {
		wc.Add(1)
		go func() {
			v, _ := b.Pick()
			fmt.Printf("%v\n", v.(string))
			fmt.Println()
			wc.Done()
		}()
	}
	wc.Wait()
}
