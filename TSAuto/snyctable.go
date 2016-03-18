package TSAuto

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
}

type Coordinate struct {
	ID                int
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
}
