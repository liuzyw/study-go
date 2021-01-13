package util

import (
	"fmt"
	"sync"
)

type MapUtil struct {
}

var mapUtil *MapUtil
var mapUtilOnce sync.Once

func GetMapUtil() *MapUtil {
	mapUtilOnce.Do(func() {
		mapUtil = &MapUtil{}
	})
	return mapUtil
}

func (u *MapUtil) mapToJson(m map[string]interface{}) string {
	return GetJsonUtil().structToJson(m)
}

func (u *MapUtil) jsonToMap(j string) map[string]interface{} {

	if j == "" {
		return map[string]interface{}{}
	}

	m := map[string]interface{}{}

	GetJsonUtil().jsonToStruct(j, &m)

	return m
}

func (u *MapUtil) putNotEmptyKeyValue(m map[string]interface{}, key string, v interface{}) {
	if key == "" {
		return
	}

	if value, ok := v.(string); ok {
		if value == "" {
			return
		}
	}
	m[key] = v
}

func (u *MapUtil) putIfAbsent(m map[string]interface{}, key string, v interface{}) {
	if key == "" {
		return
	}

	if v, ok := m[key]; ok {
		fmt.Println(fmt.Sprintf("putIfAbsent key exsit value:%v", v))
		return
	}
	m[key] = v
}
