package apiUtils

import (
	"strings"
)

type JsonObject struct {
	JsonMap map[string]interface{}
}

func (ap JsonObject) Find(path string) DataAdapter {
	paths := strings.Split(path, ".")

	var findData interface{}
	var tempData map[string]interface{}
	for i := 0; i < len(paths); i++ {
		var index = paths[i]
		if i == 0 {
			tempData = ap.JsonMap
		}

		if len(paths) == i+1 {
			findData = tempData[index].(interface{})
		} else {
			tempData = tempData[index].(map[string]interface{})
		}

	}

	da := DataAdapter{data: findData}
	return da
}
