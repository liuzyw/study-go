package util

import "sync"

type TypeUtil struct {
}

var typeUtil *TypeUtil
var typeUtilOnce sync.Once

func GetTypeUtil() *TypeUtil {
	typeUtilOnce.Do(func() {
		typeUtil = &TypeUtil{}
	})
	return typeUtil
}

