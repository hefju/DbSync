package IncubatorServer

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./incubator.db")
	engine.ShowSQL(true)
	engine.SetMapper(names.SameMapper{})
	if err != nil {
		log.Println(err)
	}
	err = engine.Sync2(new(TbHelp), new(TbUpdate), new(TbUser), new(TbUserLogin),
		new(TbUserOperation), new(TbUserManagement), new(TbIncubatorOperation),
		new(TbIncubatorAudit), new(TbIncubatorNotification), new(TbProjects),
	)

	if err != nil {
		log.Println(err)
	}
}

//2020.5.10
type TbHelp struct {
	ID       int64
	Title    string
	Content  string
	CreateAt int64
}
type TbUpdate struct {
	ID       int64
	Title    string
	Content  int
	CreateAt int64
}

type TbUser struct {
	ID          int64
	UserName    string
	Email       string
	UserGroup   string
	PhoneNumber string
	Description string
	UserPwd     string
	Online      int
	Uid         string
	CreateAt    int
}

type TbUserLogin struct {
	ID        int64
	Time      string
	User      string
	Operation string
}

type TbUserOperation struct {
	ID        int64
	Time      string
	User      string
	Operation string
}

type TbUserManagement struct {
	ID        int64
	UserName  string
	Email     string
	UserGroup string

	PhoneNumber string
	Description string
	Uid         string
	CreateAt    int
}

type TbIncubatorOperation struct {
	ID         int64
	Time       string
	Incubators string
	User       string
	Operation  string
}

type TbIncubatorAudit struct {
	ID              int64
	Time            string
	ProjectName     string
	CellName        string
	MediumName      string
	IncubatorNumber string
	Approver        string
	Results         string
}

type TbIncubatorNotification struct {
	ID           int64
	Time         string
	Notification string
}

type TbProjects struct {
	ID               int64
	ProjectName      string
	Type             string
	Template         string
	StartTime        string
	IncubationPeriod string
	User             string
	Incubators       string
	Status           int //0代表已结束,1代表在运行

}

func MockData() {
	//return
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 72; i++ {

		//	x:=TbUserOperation{Time:time.Now().AddDate(0,0,-1*i).Format("2006-01-02 15:04:05"),User:"abab",Operation:"add time"}
		//x:=TbUserManagement{UserName:"abac",Email:"15827700360@qq.com",UserGroup:"senior", PhoneNumber:"1380013800",Description:"i am a experimenter,i am a exp"}
		//x:=TbHelp{Title:"WindowStartupLocation-CenterScreen",Content:"Balabalabalabalabalabalabalabalabalabalabalabalabalabalabalabalabalabalabalabalabalabalabala"}

		//x:=TbIncubatorOperation{Time:time.Now().Add(time.Hour*-1).Format("2006-01-02 15:04:05"), Incubators:"5B",User:"Adam",Operation:"Open door"}
		mytime := time.Now().Add(time.Hour * (-4) * time.Duration(i)).Format("2006-01-02 15:04:05")
		//x:=TbIncubatorAudit{Time:mytime,ProjectName:"Embryonic stem cell research",CellName:"293E",MediumName:"acid borth",
		//	IncubatorNumber:"13",Approver:"Adam",Results:"Agree"}
		//x:=TbIncubatorNotification{Time:mytime,Notification:"Embryonic stem cell research"}
		x := TbProjects{ProjectName: "Balabalabalabalaba", Type: "Production", Template: "xxx", StartTime: mytime, IncubationPeriod: "12day", User: "Adam", Incubators: "3A,6A,5B,6C", Status: 0}

		_, err = session.Insert(&x)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(affected, err )
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}
}
