package apiUtils

import (
	"fmt"
	"net/http"
)

func GetA(url string, param Param) ApiData {
	url = param.applyA(url)
	fmt.Print("GetA:" + url + "\n")

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	return ApiData{res.Body}
}
