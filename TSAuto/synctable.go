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
	err = engine.Sync2(new(Coordinate), new(PageProfile),
		new(Microplate), new(MicroplateInfo),new(MicroplateInfo2), new(SystemImaging), new(IndependentImaging),
	new(TaskSchedule),new(SystemConfig))

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
	c := Coordinate{CoordinateOrder: 1, CoordinateName: "first form"}
	count, err := engine.Insert(&c) //InsertOne
	if err != nil {
		log.Fatal(err)
	}
	return count
}

type Coordinate struct {
	ID              int64 //`xorm:pk,autoincr`
	CoordinateID    int
	CustomID        string
	FormName        string
	ItemType        string
	CoordinateOrder int
	AreaName        string
	CoordinateName  string
	CoordinateDesc  string
	XAxis           float64
	YAxis           float64
	ZAxis           float64
	AAxis           float64
	BAxis           float64
	CAxis           float64
	IsZero          int
	Settled         int
	UpdatedAt       time.Time `xorm:"updated"`
	CreateAt        time.Time `xorm:"created"`
}
type PageProfile struct {
	ID          int64
	CustomID    string
	FormName    string
	ItemType    string
	ItemOrder   int
	AreaName    string
	ItemNumber  string
	ItemName_CN string
	Display_CN  string
	ItemName_EN string
	Display_EN  string
	UpdatedAt   time.Time `xorm:"updated"`
	CreateAt    time.Time `xorm:"created"`
}

type Microplate struct {
	ID             int64
	MicroplateName string
}

type MicroplateInfo struct {
	ID           int64
	MicroplateID int
	Manufacturer string
	Numwell      int
	Rows         int
	Cols         int
	Height       float32
	Width        float32
	Margin_x     float32
	Margin_y     float32
	Pitch_x      float32
	Pitch_y      float32
	Radius       float32
	Description  string
}


type MicroplateInfo2 struct {
	ID           int64
	MicroplateID int
	Manufacturer string
	Wells6 string
	Wells12 string
	Wells24 string
	Wells48 string
	Wells96 string

}
type SystemImaging struct {
	ID                int64
	SourcePlateName   string
	SourceMicroplates string
	Incubators        string
	Project           string
	StartTime         string
	EndTime           string
}

type IndependentImaging struct {
	ID                int64
	Name              string
	SourceMicroplates string
	//    Incubators string
	//    Project string
	StartTime string
	EndTime   string
	User      string
}

type TaskSchedule struct {
	ID int64
	BoardID string
	TaskName string//拍照
	Description string
	TimeBegin string
	TimeEnd string
	Status string//Unstart, Running, Stopp, Finish

}
type SystemConfig struct{
	ID int64
	Category string
	Key string
	Value string
	CreateAt time.Time `xorm:"created"`
}
//自动执行的参数
type AutomaticParam struct {
	ID int64
	BoardID string //板子的ID
	TaskName string//任务名称
	TaskParam string//下发参数
}













































