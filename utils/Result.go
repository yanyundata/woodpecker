package utils

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Result struct {
	response *resty.Response
	err      error
}

func (result Result) ToString() string {
	return string(result.response.Body())
}

func (result Result) ToMap() map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal((*(&result.response)).Body(), &jsonMap)

	return jsonMap
}

func (result Result) ToJson() JsonObject {
	jsonObject := JsonObject{JsonMap: result.ToMap()}

	return jsonObject
}

func (result Result) ToModel(model interface{}) {
	json.Unmarshal((*(&result.response)).Body(), model)
}
