package grammar_practice

import (
	"fmt"
	"strings"
	"testing"
)



func TestMergeSort(t *testing.T) {
	arr := []int32{3,6,8,1,2,4,0,9,7,5}
	sorted_arr := MergeSort(arr)
	for i:=0; i < len(sorted_arr);i++ {
		print(sorted_arr[i])
		print(",")
	}
	//fmt.f
}

func TestSlice(t *testing.T){
	arr := []int32{3,6,8,1,2,4,0,9,7,5}
	s := make([]int32,3)
	println(len(s))
	println(cap(s))
	println(s[0])
	s = arr[0:2]
	println(s[0])
	arr[0] = 1
	println(s[0])

	in := f3()
	println(in)
	//需要先导入Strings包
	s1 := "字符串"
	s2 := "拼接"
	var build strings.Builder
	build.WriteString(s1)
	build.WriteString(s2)
	s3 := build.String()

	ss := fmt.Sprintf("%s %s",s3,s3)
	println(ss)
}
