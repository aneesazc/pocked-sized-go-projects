package stuff

import "strconv"

var Name string = "Jim"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, v := range intArr {
		strArr = append(strArr, strconv.Itoa(v))
	}

	return strArr
}