package apiap

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func aB(method string, uri string, param string) Result {
	log.Println(method + ":" + uri)
	//A方式，对参数Param进行加密
	if len(param) > 0 && strings.Contains(param, "?") {
		//如果传递了参数
		param = param[1:]
		param = url.PathEscape(param)
		param = "?" + param
	}
	uri = uri + param
	req, _ := http.NewRequest(method, uri, nil)
	res, _ := http.DefaultClient.Do(req)

	return Result{res.Body}
}

func c(method string, url string, object interface{}) Result {
	log.Println(method + ":" + url + "\n")
	j, _ := json.Marshal(object)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(j))
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	return Result{res.Body}
}

func GetA(uri string, param string) Result {
	return aB("GET", uri, param)
}

func GetB(url string, param string) Result {
	return aB("GET", url, param)
}

func PostA(uri string, param string) Result {
	return aB("POST", uri, param)
}

func PostC(url string, object interface{}) Result {
	return c("POST", url, object)
}

func PutA(url string, param string) Result {
	return aB("PUT", url, param)
}

func PutB(url string, param string) Result {
	return aB("PUT", url, param)
}

func PutC(url string, object interface{}) Result {
	return c("PUT", url, object)
}

func DeleteA(url string, param string) Result {
	return aB("DELETE", url, param)
}

func DeleteB(url string, param string) Result {
	return aB("DELETE", url, param)
}
