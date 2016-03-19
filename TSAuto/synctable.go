package TSAuto

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"log"
)
func Donothing(){

}
var engine *xorm.Engine
func init(){
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	engine.ShowSQL = true
	engine.SetMapper(core.SameMapper{})
	if err != nil {
		log.Println(err)
	}
	err = engine.Sync2(new(Coordinate))
	if err != nil {
		log.Println(err)
	}
}

func SyncDb() {
//	var err error

//	engine, err = xorm.NewEngine("sqlite3", "./test.db")
//	log.Println("end")
}

func Insert()int64{
	//SyncDb()
	c:=Coordinate{FormOrder:1,FormName:"first form"}
	count,err:=engine.Insert(&c)//InsertOne
	if err!=nil{
		log.Fatal(err)
	}
	return count
}

type Coordinate struct {
	ID                int64  //`xorm:pk,autoincr`
	FormOrder         int
	FormName          string
	CoordinateType    string
	PointOrder        float64
	CoordinateNumber  string
	CoordinateName_EN string
	CoordinateName_CN string
	XAxis             float64
	YAxis             float64
	ZAxis             float64
	AAxis             float64
	BAxis             float64
	CAxis             float64
	CreateAt time.Time`xorm:"created"`
}
