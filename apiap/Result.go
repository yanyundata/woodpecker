package apiap

import (
	"encoding/json"
	"github.com/yanyundata/woodpecker/utils"
	"io"
	"io/ioutil"
	"log"
)

type Result struct {
	body io.ReadCloser
}

func (result Result) ToString() string {
	defer result.body.Close()

	bs, _ := ioutil.ReadAll(result.body)

	return string(bs)
}

func (result Result) ToJson() utils.JsonObject {
	defer result.body.Close()

	jsonMap := make(map[string]interface{})
	bs, _ := ioutil.ReadAll(result.body)
	err1 := json.Unmarshal(bs, &jsonMap)
	if err1 != nil {
		log.Println(err1)
	}

	jsonObject := utils.JsonObject{JsonMap: jsonMap}

	return jsonObject
}
