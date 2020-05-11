package utils

import (
	"container/list"
	"log"
	"math"
)

type DataAdapter struct {
	Data interface{}
}

func (d DataAdapter) ToString() string {
	var returnData = ""
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			returnData = "ToString Error"
		}
	}()

	returnData = d.Data.(string)
	return returnData
}

func (d DataAdapter) ToFloat() float64 {
	defer func() {
		recover()
	}()

	var returnData = math.MaxFloat64

	returnData = d.Data.(float64)
	return returnData
}

func (d DataAdapter) ToInt() int64 {
	defer func() {
		recover()
	}()
	return d.Data.(int64)
}

func (d DataAdapter) ToInterfaceList() []interface{} {
	defer func() {
		recover()
	}()

	return d.Data.([]interface{})
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

func (d DataAdapter) IsNil() bool {
	return d.Data == nil
}
