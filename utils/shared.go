package utils

import "strconv"

func StringToInt(s string) (bool, int) {
	if i, err := strconv.Atoi(s); err == nil {
		//log.Println("string is number")
		return true, i
	}
	return false, -1
}
