package apiUtils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func TestA() string {
	return "abc"
}

func GetA(url string, param Param) ApiData {
	url = param.applyA(url)
	fmt.Print("GetA:" + url + "\n")

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

type ApiData struct {
	body io.ReadCloser
}

func (ap ApiData) toString() string {
	defer ap.body.Close()

	bs, _ := ioutil.ReadAll(ap.body)

	return string(bs)
}

func (ap ApiData) toJson() map[string]interface{} {
	defer ap.body.Close()

	jsonMap := make(map[string]interface{})
	bs, _ := ioutil.ReadAll(ap.body)
	err1 := json.Unmarshal(bs, &jsonMap)
	if err1 != nil {
		log.Println(err1)
	}

	return jsonMap
}
