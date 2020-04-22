package apiUtils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func test123() {
	fmt.Print("test")
}

func GetA(url string, param Param) ApiData {
	url = param.applyA(url)
	fmt.Print("GetA:" + url + "\n")

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

type Param struct {
	param map[string]string
}

func (p Param) applyA(url string) string {
	if len(p.param) > 0 {
		url = url + "?"

		isFirst := true
		for k, v := range p.param {
			if isFirst {
				url = url + k + "=" + v
				isFirst = false
			} else {
				url = "&" + url + k + "=" + v
			}
		}
	}

	return url
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
