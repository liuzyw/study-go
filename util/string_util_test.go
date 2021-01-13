package util

import (
	"fmt"
	"testing"
)

func Test_sd(t *testing.T) {

	fmt.Println(GetStringUtil().strToInt64("123"))

	fmt.Println(GetStringUtil().int64ToStr(32321))

}
