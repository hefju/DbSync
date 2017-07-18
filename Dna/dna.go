package Dna

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

func Donothing() {

}

var engine *xorm.Engine

func init() {

	path := "./Tsdata"
	if _, err := os.Stat(path); err == nil {
	} else {
		err := os.MkdirAll(path, 0711)

		if err != nil {
			log.Println(err)
			return
		}
	}

	var err error
	engine, err = xorm.NewEngine("sqlite3", "./Tsdata/dna.db")
	engine.ShowSQL(true)
	engine.SetMapper(core.SameMapper{})
	if err != nil {
		log.Println(err)
	}
	err = engine.Sync2( new(TLabRecords), new(E_DNAData),
	)

	if err != nil {
		log.Println(err)
	}
}

//实验记录
type TLabRecords struct {
	ID            int64   //
	Ltime         string  //时间
	Lname         string  //名称
	SamplesNumber int     //样品(个)
	Temperature   float32 //温度()
	WindSpeed     float32 //风速(m/s)
	Duration      float32 //时长(min)
	Curve         string  //DNA浓度曲线对比
	Operator      string  //操作人
	Details       string  //详情
}
type E_DNAData struct {
	ID                 int64   //
	M_CheckNum         string  //检查编号 ，nfc读取
	M_Num              int     //从下位机读取
	M_DNASamNum        string  //DNA样品编号=检查编号+（1~480）
	M_Lysate           float32 //裂解液用量
	M_Detergent        float32 //洗涤液用量
	M_Disinfectant     float32 //脱洗液用量
	M_FinalVolume      string  //最终体积
	M_DnaConcentration string   //DNA 浓度  下位机发送上来
	LabRecordsID       string  //实验记录ID
}
