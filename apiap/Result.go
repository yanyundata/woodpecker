package apiap

import (
	"encoding/json"
	"github.com/yanyundata/woodpecker/utils"
	"io"
	"io/ioutil"
)

type Result struct {
	body io.ReadCloser
}

func (result Result) ToString() string {
	defer result.body.Close()

	bs, _ := ioutil.ReadAll(result.body)

	return string(bs)
}

func (result Result) ToMap() map[string]interface{} {
	defer result.body.Close()

	jsonMap := make(map[string]interface{})
	bs, _ := ioutil.ReadAll(result.body)
	json.Unmarshal(bs, &jsonMap)

	return jsonMap
}

func (result Result) ToJson() utils.JsonObject {
	defer result.body.Close()

	jsonObject := utils.JsonObject{JsonMap: result.ToMap()}

	return jsonObject
}

func (result Result) ToModel(model interface{}) {
	defer result.body.Close()

	bs, _ := ioutil.ReadAll(result.body)
	json.Unmarshal(bs, model)
}
