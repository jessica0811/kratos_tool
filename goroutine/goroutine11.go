/*
 * @Author: qiuhua.lin
 * @Date: 2022-07-10 16:25:41
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)

func Crawling(url string) []string {
	fmt.Println("crawling: ", url)
	tokens <- struct{}{} // 获取令牌
	list, err := Extracting(url)
	<-tokens // 释放令牌
	if err != nil {
		log.Println(err)
	}
	return list
}

// 提取对指定的URL发出HTTP GET请求，进行解析并响应为HTML，并返回HTML文档中的链接
func Extracting(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as html:%v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	var n int
	n++
	fmt.Println("n: ", n)           // 用来记录任务列表中的任务个数
	worklist := make(chan []string) // 创建通道数组
	fmt.Println("os: ", os.Args[1:])
	go func() { worklist <- os.Args[1:] }() // 已经获取的链接
	picked := make(map[string]bool)
	for ; n > 0; n-- {
		fmt.Println("for n: ", n)
		list := <-worklist
		fmt.Println("list: ", list)
		for _, url := range list {
			if !picked[url] {
				fmt.Println("url: ", url)
				picked[url] = true
				n++
				fmt.Println("n....: ", n)
				go func(url string) { // 开启 goroutine
					worklist <- Crawling(url)
				}(url)
			}
		}
	}
}
