package utils

import (
	"container/list"
	"math"
)

type DataAdapter struct {
	Data interface{}
}

func (d DataAdapter) ToString() string {
	return d.Data.(string)
}

func (d DataAdapter) ToFloat() float64 {
	var returnData = math.MaxFloat64

	returnData = d.Data.(float64)
	return returnData
}

func (d DataAdapter) ToInt() int64 {
	return d.Data.(int64)
}

func (d DataAdapter) ToInterfaceList() []interface{} {
	return d.Data.([]interface{})
}

func (d DataAdapter) ToJsonObjectList() *list.List {
	interfaceList := d.ToInterfaceList()
	JsonObjectList := list.New()

	for _, intercate := range interfaceList {
		object := JsonObject{intercate.(map[string]interface{})}
		JsonObjectList.PushBack(object)
	}

	return JsonObjectList
}

func (d DataAdapter) IsNil() bool {
	return d.Data == nil
}
