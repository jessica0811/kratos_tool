/*
 * @Author: qiuhua.lin
 * @Date: 2022-05-20 19:02:44
 */
package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type QueryError struct {
	Query string
	Err   error
}

func (e QueryError) Error() string {
	return e.Query + ":" + e.Err.Error()
}

func New(text string, err error) error {
	return QueryError{Query: text, Err: err}
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

func ReadFile(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return &QueryError{"文件没有打开", err}
	}
	return nil
}

func ReadConfig() error {
	home := os.Getenv("home")
	path := filepath.Join(home + ".settings.xml")
	err := ReadFile(path)
	return &QueryError{"eof", err}
}

func main() {
	// err := ReadFile("a.txt")
	// ErrNotFound := errors.New("not found")
	// if e, ok := err.(*QueryError); ok {
	// 	fmt.Println("something not found", e.Query)
	// }
	err := ReadConfig()
	// var e *QueryError
	// if errors.As(err, &e) {
	// 	fmt.Println(e.Err)
	// }
	fmt.Println(err)

	ErrNotFound := New("eof", errors.New("文件没有打开:open .settings.xml: no such file or directory"))
	fmt.Println(ErrNotFound)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("进来了")
	} else {
		fmt.Println("2222")
	}

}
