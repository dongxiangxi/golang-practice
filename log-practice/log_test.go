package log_practice

import (
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T)  {
	println("test begin ")
	log.Fatalln("Come with fatal,exit with 1 ")
	println("test end ")

}

func  Test2(t *testing.T)  {

	defer func() {
		if err := recover();err != nil {
			log.Println("this is recover")
			log.Println("e is :",err)
			log.Println("recover end ")
		}
	}()
	arr := []int {2,3}
	//log.SetFlags(log.Ldate)
	log.Print("Print array ",arr,"\n")
	//log.Panicln("this is Panicln")
	log.Println("Println array",arr)
	log.Printf("Printf array with item [%d,%d]\n",arr[0],arr[1])
}

func  Test3(t *testing.T)  {
	logFileName := "test.log"
	logFile,err := os.Create(logFileName)
	defer  logFile.Close()
	if err != nil {
		log.Fatalln("create file failed")
	}
	//	都是2的倍数，进行或运算 |
	myLogger := log.New(logFile,"[DEBUG]",log.Ldate | log.Ltime)

	myLogger.Println("this is myLogger debug")





}