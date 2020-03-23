package io_practice

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os"
	"testing"
)

//首先需要熟悉一下io及ioutil包中的接口与方法
/*
io来自哪里？
1、控制台
2、网络
3、文件
4、管道
 */
func TestFileIO(t *testing.T)  {
	//"你好，世界！"的GBK编码
	testBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	var testStr string
	utfStr := "你好，世界！"
	var dec mahonia.Decoder
	var enc mahonia.Encoder

	testStr = string(testBytes)
	println(testStr)
	println("==========")
	//dec.NewReader()
	gbkbytes ,_ := ioutil.ReadFile("gbk.txt")
	//n, utf8bytes,err := dec.Translate(testBytes,true)
	//println(n)
	println(string(gbkbytes))
	//println(err1)

	dec = mahonia.NewDecoder("gbk")

	utf8str := dec.ConvertString(string(gbkbytes))
	//_, utf8bytes,_ := dec.Translate(gbkbytes,false)
	println("==========")
	println(utf8str)
	//println(string(utf8bytes))
	fileName := "utf8.txt"
	filePointer, _ := os.Create(fileName)
	defer filePointer.Close()
	filePointer.WriteString(utf8str)
	filePointer.Close()

	ret := dec.ConvertString(testStr)
	fmt.Println("GBK to UTF-8: ", ret, " bytes:", testBytes)


	enc = mahonia.NewEncoder("gbk")
	ret = enc.ConvertString(utfStr)
	fmt.Println("UTF-8 to GBK: ", ret, " bytes: ", []byte(ret))

	return
}

func TestFileTranslate(t *testing.T)  {
	gbkContent,_ := ioutil.ReadFile("gbk.txt")
	//gbkfile,_ := os.Open("gbk.txt")
	//gbkfile.
	//println(string(gbkContent))
	decoder := mahonia.NewDecoder("gbk")
	//eader := decoder.NewReader(gbkfile)
	_,utf8bytes,_ := decoder.Translate(gbkContent,true)
	//var buffer []byte
	filePointer,_ := os.Create("utf8.txt")
	defer filePointer.Close()

	//filePointer.w
	//n, _ := reader.Read(buffer)
//	println(n)
	//utf8bytes := buffer[0:n]
	println(string(utf8bytes))
	filePointer.WriteString(string(utf8bytes))
}