package apiUtils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var BaseUrl = ""

func init() {
	BaseUrl = os.Getenv("BaseUrl")
}

func GetA(url string, param string) ApiData {
	url = url + param
	log.Println("GetA:" + BaseUrl + url)

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func GetB(url string, param string) ApiData {
	url = url + param
	log.Println("GetB:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("GET", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func PostC(url string, object interface{}) ApiData {
	log.Println("GetB:" + url + "\n")
	j, _ := json.Marshal(object)

	req, _ := http.NewRequest("POST", BaseUrl+url, bytes.NewBuffer(j))
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

func PostA(url string, param string) ApiData {
	url = url + param
	log.Println("PostA:" + BaseUrl + url + "\n")

	req, _ := http.NewRequest("POST", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}

//put 请求 请求路径参数
// put: https://ip:port/test/put/{one}/{two}
func PutA(url string, param string) ApiData {
	url = url + param
	log.Println("PutA:" + BaseUrl + url + "\n")
	req, _ := http.NewRequest("PUT", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}

//put 请求 请求路径参数
//put: https://ip:port/test/put/two?param1=xxx&param2=xxxx
func PutB(url string, param string) ApiData {
	url = url + param
	log.Println("PutB:" + BaseUrl + url)
	req, _ := http.NewRequest("PUT", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}

//put 请求 请求体参数
//put : https://ip:port/test/put/three
func PutC(url string, object interface{}) ApiData {
	log.Println("PutC:" + url + "\n")
	j, _ := json.Marshal(object)

	req, _ := http.NewRequest("PUT", BaseUrl+url, bytes.NewBuffer(j))
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}

//delete 请求 请求路径参数
// delete https://ip:port/test/delete?id=xxx
func DeleteA(url string, param string) ApiData {
	log.Println("DeleteA:" + BaseUrl + url)
	req, _ := http.NewRequest("DELETE", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}

//delete 请求，请求路径参数
//delete https://ip:port/test/delete/{id}
func DeleteB(url string, param string) ApiData {
	url = url + param
	log.Println("DeleteB:" + BaseUrl + url)
	req, _ := http.NewRequest("DELETE", BaseUrl+url, nil)
	res, _ := http.DefaultClient.Do(req)
	return ApiData{res.Body}
}
