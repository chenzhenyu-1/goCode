package test

import (
	"class2/tool"
	"fmt"
)

func MyTest() int64 {
	id := tool.NewIdInstance()
	idValue := id.GetID()
	fmt.Println(idValue)
	return idValue
}
