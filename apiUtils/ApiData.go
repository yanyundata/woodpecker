package apiUtils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

type ApiData struct {
	body io.ReadCloser
}

func (ap ApiData) ToString() string {
	defer ap.body.Close()

	bs, _ := ioutil.ReadAll(ap.body)

	return string(bs)
}

func (ap ApiData) ToJson() JsonObject {
	defer ap.body.Close()

	jsonMap := make(map[string]interface{})
	bs, _ := ioutil.ReadAll(ap.body)
	err1 := json.Unmarshal(bs, &jsonMap)
	if err1 != nil {
		log.Println(err1)
	}
	data := jsonMap["data"].(interface{})
	log.Print(data)

	jsonObject := JsonObject{jsonMap}

	return jsonObject
}
