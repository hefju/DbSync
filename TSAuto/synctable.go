package TSAuto

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func Donothing() {

}

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	//engine.ShowSQL = true
	engine.SetMapper(core.SameMapper{})
	if err != nil {
		log.Println(err)
	}
	err = engine.Sync2(new(Coordinate),new(PageProfile))
	if err != nil {
		log.Println(err)
	}
}

func SyncDb() {
	//	var err error

	//	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	//	log.Println("end")
}

func Insert() int64 {
	//SyncDb()
	c := Coordinate{COrder: 1, CoordinateName: "first form"}
	count, err := engine.Insert(&c) //InsertOne
	if err != nil {
		log.Fatal(err)
	}
	return count
}

type Coordinate struct {
	ID                int64 //`xorm:pk,autoincr`
	CoordinateID     int
	IsZero            int
	ZkyID int
	COrder int
	CoordinateName  string
	XAxis             float64
	YAxis             float64
	ZAxis             float64
	AAxis             float64
	BAxis             float64
	CAxis             float64
	Settled           int
	UpdatedAt time.Time `xorm:"updated"`
	CreateAt          time.Time `xorm:"created"`
}
type PageProfile struct {
	ID          int64
	FormName    string
	ItemType    string
	ItemOrder   int
	ItemNumber  string
	ItemName_CN string
	ItemName_EN string
	UpdatedAt time.Time `xorm:"updated"`
	CreateAt          time.Time `xorm:"created"`
}
