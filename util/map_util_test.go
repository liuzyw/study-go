package util

import (
	"fmt"
	"testing"
	"time"
)

func Test_map(t *testing.T) {

	fmt.Println(GetMapUtil().mapToJson(map[string]interface{}{
		"dsdd":     123,
		"dsdas":    "dsdas",
		"dsddffff": time.Now(),
	}))

	j := `{"dsdas":"dsdas","dsdd":123,"dsddffff":"2021-01-13T15:04:59.718036+08:00"}`
	m := GetMapUtil().jsonToMap(j)
	fmt.Println(m)
	GetMapUtil().putNotEmptyKeyValue(m, "dsd", "dsadas")
	GetMapUtil().putNotEmptyKeyValue(m, "aaaaaaaa", 213123)
	GetMapUtil().putIfAbsent(m, "aaaaaaaa", 4)
	GetMapUtil().putIfAbsent(m, "bbbbbbbb", 4)

	fmt.Println(m)

}
