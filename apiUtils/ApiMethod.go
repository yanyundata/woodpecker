package apiUtils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//var BaseUrl = "https://yan-yun.com:38085/lvyuan"
var BaseUrl = "http://localhost:9000"

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

//put 请求 请求路径参数
// put: https://ip:port/test/put/{one}/{two}
func PutA(url string, param string) ApiData {
	url = url + param
	fmt.Printf("PutA:" + BaseUrl + url + "\n")
	req, _ := http.NewRequest("PUT", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}

//put 请求 请求路径参数
//put: https://ip:port/test/put/two?param1=xxx&param2=xxxx
func PutB(url string, param string) ApiData {
	url = url + param
	fmt.Println("PutB:" + BaseUrl + url)
	req, _ := http.NewRequest("PUT", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}

//put 请求 请求体参数
//put : https://ip:port/test/put/three
func PutC(url string, object interface{}) ApiData {
	fmt.Print("PutC:" + url + "\n")
	j, _ := json.Marshal(object)

	req, _ := http.NewRequest("PUT", BaseUrl+url, bytes.NewBuffer(j))
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}

//delete 请求 请求路径参数
// delete https://ip:port/test/delete?id=xxx
func DeleteA(url string, param string) ApiData {
	fmt.Println("DeleteA:" + BaseUrl + url)
	req, _ := http.NewRequest("DELETE", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}

//delete 请求，请求路径参数
//delete https://ip:port/test/delete/{id}
func DeleteB(url string, param string) ApiData {
	fmt.Println("DeleteB:" + BaseUrl + url)
	req, _ := http.NewRequest("DELETE", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}
