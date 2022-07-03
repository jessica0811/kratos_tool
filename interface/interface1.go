/*
 * @Author: qiuhua.lin
 * @Date: 2022-06-12 15:55:23
 */
package main

import "fmt"

type Programmer struct {
	Name string
}

type SkillInterface interface {
	Write()
	Read()
}

func (stu Programmer) Write() {
	fmt.Println("programmer write()")
}

func main() {
	var pro Programmer
	pro.Write()
}
