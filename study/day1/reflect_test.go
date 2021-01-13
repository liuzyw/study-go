package day1

import (
	"fmt"
	"reflect"
	"testing"
)

type Turbo struct {
	Name string
	Age  int
}

func Test_1232(t *testing.T) {
	turbo := &Turbo{
		Name: "迈莫coding",
		Age:  1,
	}
	value := reflect.ValueOf(turbo)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fmt.Printf("字段类型：%v, 字段值：%v\n", field.Type(), field.Interface())
		}
		st := value.FieldByName("Name")
		fmt.Printf("%v\n", st.Interface())
	}

}

func Test_dasdasd(t *testing.T) {

}
