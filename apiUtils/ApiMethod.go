package apiUtils

import (
	"fmt"
	"net/http"
)

var BaseUrl = "https://yan-yun.com:38085/lvyuan"

func GetA(url string, param Param) ApiData {
	url = param.applyA(url)
	fmt.Print("GetA:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func GetB(url string, param Param) ApiData {
	url = param.applyB(url)
	fmt.Print("GetB:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func PostC1(url string, param Param) ApiData {
	url = param.applyB(url)
	fmt.Print("GetB:" + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
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
