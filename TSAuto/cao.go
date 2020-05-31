package TSAuto

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

//var engine *xorm.Engine
func Update0508() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	engine.SetMapper(names.SameMapper{})
	if err != nil {
		fmt.Println(err)
	}
	cs := make([]Coordinate, 0)
	err = engine.Where("FormName='Frm05_8'").Find(&cs)
	if err != nil {
		fmt.Println(err)
	}

	ps := make([]PageProfile, 0)
	err = engine.Where("FormName='Frm05_8'").Find(&ps)
	if err != nil {
		fmt.Println(err)
	}

	for _, coor := range cs {
		cname := coor.CoordinateName
		for _, pro := range ps {
			pname := pro.ItemName_CN + pro.Display_CN
			if cname == pname {
				coor.CustomID = pro.CustomID
				_, err := engine.ID(coor.ID).Cols("CustomID").Update(&coor)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

}
