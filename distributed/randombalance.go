/*
 * @Author: qiuhua.lin
 * @Date: 2022-07-19 10:44:27
 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type RandomBalance struct {
	m        sync.Mutex
	curIndex int
	rss      []string
}

func New(rss []string) *RandomBalance {
	return &RandomBalance{rss: rss}
}

// 生成下一个随机字符串
func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.m.Lock()
	r.curIndex = rand.Intn(len(r.rss))
	r.m.Unlock()
	return r.rss[r.curIndex]
}

func main() {
	source := []string{"10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4"}
	b := New(source)
	wc := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		v := b.Next()
		fmt.Printf("%v\n", v)
	}
	wc.Wait()
}
