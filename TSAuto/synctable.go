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
	engine.ShowSQL(true)
	engine.SetMapper(core.SameMapper{})
	if err != nil {
		log.Println(err)
	}
	err = engine.Sync2(new(Coordinate), new(PageProfile),new(Microplate), new(MicroplateInfo),new(MicroplateInfo2), new(SystemImaging), new(IndependentImaging),
	new(TaskSchedule),new(SystemConfig),new(AutomaticParam)	,	new(BoxInfo),	new(CellPlateInfo),	new(CellSelectInfo),
	new(JvUSER),new(USER),new(Power),new(Charge),new(Log),new(ProductCategory),
		//2017TS新项目
		new(WorkItem),new(CommandS),
	)

	if err != nil {
		log.Println(err)
	}
}
//工作流程
type WorkItem struct{
	ID int64
	Code string
	Name string
	CommandID int64
	Orders float32
}
//执行命令
type CommandS struct {
	ID int64
	Command string
	Param string
	ParentId int64
}

type ProductCategory struct {
	ProductCategoryID int64
	Name string
	Other string
}
//2017.1.22
type Log struct {
	ID int64
	Date time.Time
	Thread string
	Level string
	Logger string
	Msg string
}


//2016.12.10 ldh充电量
type Charge struct {
	ID int64
	FYear int
	FMonth int
	FDay int
	Pile string  //电桩
	Card string //卡号
	Line string  //线路
	Car string   //车号
	Chminute string//分钟
	Money float32//金额
	Pw float32 //电量
	Soc1 int
	Soc2 int
	StDate string//开始时间
	EdDate string//结束时间
	Chtime string//充电时间
	FromFile string
}
type Power struct {
	ID int64
	FYear int
	FMonth int
	FDay int
	Line string
	Car string
	Pw float32
	Soc1 int
	Soc2 int
	FromFile string
}

//2016.11.23 学习java
type JvUSER struct {
	ID int64
	Name string
	Password string
}
type USER struct {
	ID int64
	Name string
	Password string
}

//箱子 2016.11.20
type BoxInfo struct {
	ID int64
	BoxId string  //1,2,3, 28,29, 不是主键
	BoxNo string  //01C, 01A
	ProjectId string
	ProjectName string
	Statu int
}
//板子的信息
type  CellPlateInfo struct  {
	ID int64
	BoxNo string  //外键, 01A,01C
	PlateId string //不是主键
	PlateName string
	PlatePos string
	Tid string
	StartTime string
	Statu int
}
//选中的孔
type  CellSelectInfo struct  {
	ID int64
	PlateId string     //外键  联合主键
	Tid string
	SelectedCell string //联合主键
	//IsSelected string
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

//坐标
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
//下拉框
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
//
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
//系统拍照
type SystemImaging struct {
	ID                int64
	SourcePlateName   string
	SourceMicroplates string
	Incubators        string
	Project           string
	StartTime         string
	EndTime           string
	PhotoPath	string//保存路径
	ImagingType string//是system,还是idependent

}
//独立成像
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

//任务列表
type TaskSchedule struct {
	ID int64
	BoardID string
	Tid string     //rfid
	TaskName string//拍照
	Description string
	TimeBegin string
	TimeEnd string
	Status string//Unstart, Running, Stopp, Finish

}
//系统配置
type SystemConfig struct{
	ID int64
	Category string
	Key string
	Value string
	CreateAt string//time.Time `xorm:"created"`
}
//自动执行的参数
type AutomaticParam struct {
	ID int64
	BoardID string //板子的ID
	Tid string
	TaskName string//任务名称
	FocusX int  //对焦x
	FocusY int  //对焦y
	FocusZ int//对焦z
	Lensel int//放大倍数, 4,10,20
	Filtersel int//滤镜
	FluExposureTime string//曝光时间
	TaskParam string//下发参数
}













































