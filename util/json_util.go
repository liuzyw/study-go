package util

import (
	"encoding/json"
	"fmt"
	"sync"
)

type JsonUtil struct {
}

var jsonUtil *JsonUtil
var jsonUtilOnce sync.Once

func GetJsonUtil() *JsonUtil {
	jsonUtilOnce.Do(func() {
		jsonUtil = &JsonUtil{}
	})
	return jsonUtil
}

func (u *JsonUtil) jsonToStruct(j string, obj interface{}) {

	if j == "" {
		return
	}

	err := json.Unmarshal([]byte(j), &obj)
	if err != nil {
		fmt.Println(fmt.Sprintf("jsonToStruct Unmarshal fail, err:%v", err))
	}
}

func (u *JsonUtil) structToJson(obj interface{}) string {

	if obj == nil {
		return ""
	}

	bytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(fmt.Sprintf("structToJson Marshal fail, err:%v", err))
		return ""
	}

	return string(bytes)
}
