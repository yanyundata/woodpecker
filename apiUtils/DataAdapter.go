package apiUtils

import (
	"container/list"
	"math"
)

type DataAdapter struct {
	data interface{}
}

func (d DataAdapter) ToString() string {
	var returnData = ""
	defer func() {
		if err := recover(); err != nil {
			returnData = "ToString Error"
		}
	}()

	returnData = d.data.(string)
	return returnData
}

func (d DataAdapter) ToFloat() float64 {
	defer func() {
		recover()
	}()

	var returnData = math.MaxFloat64

	returnData = d.data.(float64)
	return returnData
}

func (d DataAdapter) ToInterfaceList() []interface{} {
	defer func() {
		recover()
	}()

	return d.data.([]interface{})
}

func (d DataAdapter) ToJsonObjectList() *list.List {
	defer func() {
		recover()
	}()

	interfaceList := d.ToInterfaceList()
	JsonObjectList := list.New()

	for _, intercate := range interfaceList {
		object := JsonObject{intercate.(map[string]interface{})}
		JsonObjectList.PushBack(object)
	}

	return JsonObjectList
}
