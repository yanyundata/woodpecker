package apiUtils

import (
	"strings"
)

type JsonObject struct {
	jsonMap map[string]interface{}
}

func (ap JsonObject) Find(path string) DataAdapter {
	paths := strings.Split(path, ".")

	var tempData interface{}
	for i := 0; i < len(paths); i++ {
		v := ap.jsonMap[paths[i]]
		tempData = v
	}

	da := DataAdapter{data: tempData}
	return da
}
