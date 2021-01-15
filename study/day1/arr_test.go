package day1

import (
	"fmt"
	"testing"
)

func Test_Arr(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	fmt.Println(a)

	b := []int{1, 2, 3}
	b = append(b, 2)
	fmt.Println(b)

	c := make([]int, 0)
	fmt.Println(c)
	c = append(c, b...)
	fmt.Println(c)

}
