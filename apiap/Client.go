package apiap

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var BaseUrl = ""

func init() {
	BaseUrl = os.Getenv("BaseUrl")
}

func aB(method string, uri string, param string) Result {
	log.Println(method + ":" + BaseUrl + uri)
	prefix := ""
	if strings.Contains(param, "?") {
		prefix = "?"
		param = param[1:]
	}
	escape := url.PathEscape(param)
	uri = uri + prefix + escape
	req, _ := http.NewRequest(method, BaseUrl+uri, nil)
	res, _ := http.DefaultClient.Do(req)

	return Result{res.Body}
}

func c(method string, url string, object interface{}) Result {
	log.Println(method + ":" + url + "\n")
	j, _ := json.Marshal(object)

	req, _ := http.NewRequest(method, BaseUrl+url, bytes.NewBuffer(j))
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	return Result{res.Body}
}

func GetA(url string, param string) Result {
	return aB("GET", url, param)
}

func GetB(url string, param string) Result {
	return aB("GET", url, param)
}

func PostA(url string, param string) Result {
	return aB("POST", url, param)
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
