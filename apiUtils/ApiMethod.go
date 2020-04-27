package apiUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var BaseUrl = "https://yan-yun.com:38085/lvyuan"

func GetA(url string, param string) ApiData {
	url = url + param
	fmt.Print("GetA:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func GetB(url string, param string) ApiData {
	url = url + param
	fmt.Print("GetB:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func PostC(url string, object interface{}) ApiData {
	fmt.Print("GetB:" + url + "\n")
	j, _ := json.Marshal(object)

	req, _ := http.NewRequest("POST", BaseUrl+url, bytes.NewBuffer(j))
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func PostA(url string, param string) ApiData {
	url = url + param
	fmt.Print("PostA:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("POST", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}
