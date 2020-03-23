package exception_test

import (
	"fmt"
	"testing"
)

func no_cache(){
	defer func() {
		println("defer ,do nothing, will throw")
	}()
	println("func no_cache begin ")
	arr := []int32{1,78}
	println(arr[3])
	println("func no_cache end ")

}

func TestThrow(t *testing.T)  {

	println("test start ")
	no_cache()
	//达不到这里，直接在上一步抛出，未捕获，crash
	println("test end ")
}

func catch()  {
	defer func() {
		println("defer start")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		println("defer end ")
	}()
	println("func catch begin ")
	arr := []int32{1,78}
	println(arr[3])
	println("func catch end ")
}

func  TestCatch(t *testing.T)  {
	println("test start ")
	catch()
	println("test end ")
	//上一步捕获异常，程序走到结束
}


func mypanic(){
	defer func() {
		println("defer start ")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		println("defer end ")
	}()
	println("func mypanic start ")

	panic("exception")


	println("func mypanic  end ")


}

//异常 执行defer后并不会回到本函数，与java的 try catch 程序继续执行catch之后的代码不一样

func  TestMypanic(t *testing.T)  {
	println("test start ")
	mypanic()
	println("test end ")
	//上一步捕获异常，程序走到结束
}

func  TestDefer(t *testing.T)  {
	println(c())
	println(f())
	println(f2())
	println(f3())

}

func c() (i int) {
	defer func() { i++ }()
	return 1
}


func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}


func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
