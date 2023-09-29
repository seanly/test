package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gitlab"
	password = "gitlab123"
	dbName   = "hello"
)

func main() {
	engine := getDBEngine()

	err := engine.Sync2(new(UserTbl))
	if err != nil {
		log.Fatal(err)
	}

	user := &UserTbl{
		Id:       1,
		Username: "Windows",
		Sex:      1,
		Info:     "操作系统",
	}

	sessionUserTest(user)

}

func getDBEngine() *xorm.Engine {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	engine, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	engine.ShowSQL()

	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("connect postgresql success")
	return engine
}

type UserTbl struct {
	Id       int
	Username string
	Sex      int
	Info     string
}

func selectAll() {
	var user []UserTbl
	engine := getDBEngine()
	engine.SQL("select * from user_tbl").Find(&user)

	fmt.Println(user)
}

func InsertUser(user *UserTbl) bool {
	engine := getDBEngine()
	rows, err := engine.Insert(user)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if rows == 0 {
		return false
	}

	return true
}

func selectUser(name string) {
	var user []UserTbl
	engine := getDBEngine()
	engine.Where("user_tbl.username=?", name).Find(&user)

	fmt.Println(user)
}

func selectOne(id int) {
	var user UserTbl
	engine := getDBEngine()
	engine.ID(id).Get(&user)
	//engine.Alias("u").Where("u.id=?", id).Get(&user)
	fmt.Println(user)
}

func sessionUserTest(user *UserTbl) {
	engine := getDBEngine()
	session := engine.NewSession()

	session.Begin()

	_, err := session.Insert(user)
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}

	err = session.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
