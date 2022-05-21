/*
 * @Author: qiuhua.lin
 * @Date: 2022-05-20 23:02:09
 */
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Instance struct {
	InstanceId string `xorm:"InstanceId"`
	Regiond    string `xorm:"Regiond"`
	Status     string `xorm:"Status"`
}

func conn() (engin *xorm.Engine, err error) {
	engine, err := xorm.NewEngine("mysql", "root:Dachangtui123$%^@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		return engin, err
	}
	return engine, nil
}

func other(engine *xorm.Engine) error {
	var ids []Instance
	err := engine.Table("data_map_instances").Where("InstanceId = 'rm-uf6iumn5wrhzt44'").Find(&ids)
	if err != nil {
		return err
	}
	return nil
}

func query(engine *xorm.Engine) error {
	var name string
	err := engine.DB().QueryRow("select InstanceId, Regiond, Status from data_map_instances where InstanceId = ?", "rm-uf6iumn5wrhzt44").Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows { // 应用程序代码通常不会将空结果视为错误，如果不检查错误是否为这个特殊常量，就会导致意想不到的应用程序代码错误。

		} else {
			return err
		}
	}
	return nil
}

func main() {
	engine, err := conn()
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Close()
	err = other(engine)
	fmt.Println("other: ", err)
	err = query(engine)
	fmt.Println("query: ", err)
}
