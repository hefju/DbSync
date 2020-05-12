//这个工具的作用是数据库同步.
//由于我会在多个数据库之间切换, 于是建表, 更新表的时候就变得非常麻烦了.
//我的想法是通过xorm的同步功能, 更新多个数据库
//首先根据配置文件来打开要更新数据库, 然后执行同步,配置文件可以设定不同步
package main

import (
    "fmt"
    "github.com/hefju/DbSync/CSharp"
    "github.com/hefju/DbSync/IncubatorServer"
    "github.com/hefju/DbSync/TSAuto"
)
func main(){


    test2020()

    fmt.Println("main end")
}


func test2020(){
 //   IncubatorServer.SyncDbSchema()
    IncubatorServer.MockData()
}


func test2017(){
    CSharp.Donothing()
    // CSharp.Convert()
    TSAuto.Donothing()//   TSAuto.SyncDb()
    // TSAuto.SyncDb()
    TSAuto. Update0508()
}
