package myDatabase

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {

	var err error
	Engine, err = xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = Engine.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

}
