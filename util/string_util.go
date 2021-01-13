package util

import (
	"strconv"
	"sync"
)

type StringUtil struct {
}

var stringUtil *StringUtil
var stringUtilOnce sync.Once

func GetStringUtil() *StringUtil {
	stringUtilOnce.Do(func() {
		stringUtil = &StringUtil{}
	})
	return stringUtil
}

func (u *StringUtil) strToInt64(str string) int64 {
	if str == "" {
		return 0
	}

	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}

	return num
}

func (u *StringUtil) int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}
