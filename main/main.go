package main

import (
	"fmt"
	"github.com/golang-practice/grammar_practice"
	"github.com/golang/example/stringutil"
	"log"
)
func main()  {
	//print(1)
	fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))
	arr := []int32{3,6,8,1,2,4,0,9,7,5}
	sorted_arr := grammar_practice.MergeSort(arr)
	for i:=0; i < len(sorted_arr);i++ {
		print(sorted_arr[i])
		print(",")
	}

	arr = []int32 {2,3}
	log.Print("Print array ",arr,"\n")
	log.Println("Println array",arr)
	log.Printf("Printf array with item [%d,%d]\n",arr[0],arr[1])
}
