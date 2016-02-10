package main
import (
	"log"
	"fmt"
	"encoding/json"
	"os"
)
func main(){
	array:=make( []*Config,5)
	array[0]=&Config{Key:"sqlite",Value:"local.db"}
	array[1]=&Config{Key:"mysql",Value:"server:192.168.1.100"}
	array[2]=&Config{Key:"postgre",Value:"server:192.168.1.100"}


	b,err:=json.Marshal(array)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("json:",b)

	f,err:=os.Create("config.json")
	if err!=nil{
		log.Fatal(err)
	}
	f.Write(b)
	defer f.Close()
	//fmt.Println(c)
	fmt.Println("end")
}
type Section  struct{
	Comments string
	Cf Config
}
type Config struct{
	Key string
	Value string
}

