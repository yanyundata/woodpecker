package apiUtils

import "container/list"

type DataAdapter struct {
	data interface{}
}

func (d DataAdapter) ToString() string {
	return d.data.(string)
}

func (d DataAdapter) ToFloat() float64 {
	return d.data.(float64)
}

func (d DataAdapter) ToInterfaceList() []interface{} {
	return d.data.([]interface{})
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
