package apiUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var BaseUrl = "https://yan-yun.com:38085/lvyuan"

func GetA(url string, param PathParam) ApiData {
	url = param.applyA(url)
	fmt.Print("GetA:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func GetB(url string, param PathParam) ApiData {
	url = param.applyB(url)
	fmt.Print("GetB:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func PostC1(url string, object interface{}) ApiData {
	fmt.Print("GetB:" + url + "\n")
	j, _ := json.Marshal(object)
	fmt.Print(string(j))

	req, _ := http.NewRequest("POST", BaseUrl+url, bytes.NewBuffer(j))
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func PostC2(url string, param Param) ApiData {
	url = param.applyB(url)
	fmt.Print("GetB:" + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}
