package IncubatorServer

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)



var engine *xorm.Engine

func SyncDbSchema() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./incubator.db")
	engine.ShowSQL(true)
	engine.SetMapper(core.SameMapper{})
	if err != nil {
		log.Println(err)
	}
	err = engine.Sync2(new(TbHelp), new(TbUpdate),new(TbUser), new(TbUserLogin),
		new(TbUserOperation),
		)

	if err != nil {
		log.Println(err)
	}
}
//2020.5.10
type TbHelp struct{
	ID int64
	Title string
	Content int
	CreateAt int64
}
type TbUpdate struct{
	ID int64
	Title string
	Content int
	CreateAt int64
}

type TbUser struct {
	ID int64
	UserName string
	Email string
	UserGroup string
	PhoneNumber string
	Description string
	UserPwd string
	Online int
	Uid string
	CreateAt int
}

type TbUserLogin struct {
	ID int64
	Time string
	User string
	Operation string
}


type TbUserOperation struct {
	ID int64
	Time string
	User string
	Operation string
}














































