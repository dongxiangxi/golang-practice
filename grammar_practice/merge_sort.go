package grammar_practice

//import "fmt"

func MergeSort(arr []int32) []int32{
	if arr == nil {
		return []int32{}
	}
	lenArr := len( arr)
	if lenArr == 1 {
		return  arr
	}
	mid := lenArr /2
	left := arr[0 :mid]
	righe := arr[mid :lenArr]

	return_left := MergeSort(left)
	return_right := MergeSort(righe)

	result := merge(return_left,return_right)
	return  result
}

func merge(a []int32, b []int32) []int32{
	if a == nil {
		a = []int32{}
	}
	if b == nil {
		b = []int32{}
	}
	lenA := len(a)
	lenB := len(b)

	m,n:= 0,0

	return_arr := []int32{}
	for  m < lenA && n < lenB {
		if a[m] < b[n] {
			return_arr =  append(return_arr,a[m])
			m++
		}else {
			return_arr =  append(return_arr,b[n])
			n++
		}
	}

	if m < lenA {
		for ; m < lenA ; m++ {
			return_arr =  append(return_arr,a[m])
		}
	} else {
		for ; n < lenB ; n++ {
			return_arr =  append(return_arr,b[n])
		}
	}

	return  return_arr

}
//
//func main(){
//
//
//}
