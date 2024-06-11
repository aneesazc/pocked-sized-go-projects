package main

import (
	stuff "ex/project/mypackage"
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Hello", stuff.Name)
	intArr := []int{1, 2, 3, 4, 5}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}