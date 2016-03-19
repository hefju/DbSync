package CSharp

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

//为了避免修改import,做一个无用的方法
func Donothing(){

}
//将c#的类转换成golang的struct
func ConvertToStruct(class string) {
	reader := bytes.NewBufferString(class)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		dealwithonline(strings.Trim(scanner.Text(), " "))//去掉空格
	}
	//fmt.Println("Convert finish")
}
func dealwithonline(line string) {//处理每一行代码
	if len(line) < 1 {//如果是空白就跳过
		return
	}
	if strings.Contains(line, "class") {//将class转成struct
		fmt.Println(strings.Replace(line, "public class", "type", -1) + " struct{")
	} else {
		if line == "}" {
			fmt.Println("CreateAt time.Time`xorm:\"created\"`")
			fmt.Println("}")
		} else {//处理public int ID { get; set; }
			list := strings.Split(line, " ")
			if len(list) > 1 {
				fmt.Println(list[2] + " " + GetGoType( list[1]))
			}
		}
	}

}
func GetGoType(cstype string)string {
	switch cstype {
		case "float":
		return "float64"
		default:
		break;
	}
	return cstype
}
func Convert() {
	var class = `
    public class Coordinate
    {
        public int ID { get; set; }

        public int FormOrder { get; set; }  //窗口顺序
        public string FormName { get; set; } //窗口名

        public string CoordinateType { get; set; } //坐标分类
        public float PointOrder { get; set; }  //坐标顺序
        public string CoordinateNumber { get; set; }   //1,2-1,2-2,3,4,5
        public string CoordinateName_EN { get; set; }
        public string CoordinateName_CN { get; set; }//1未抓取安全点,2-1未抓取上位点,2-3未抓取下位点,

        public float XAxis { get; set; }
        public float YAxis { get; set; }
        public float ZAxis { get; set; }
        public float AAxis { get; set; }
        public float BAxis { get; set; }
        public float CAxis { get; set; }
    }
    `
	ConvertToStruct(class)
}
