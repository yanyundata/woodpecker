package main

import "github.com/yanyundata/woodpecker/apiUtils"

func main() {
	apiUtils.test123()
	json := apiUtils.GetA(("http://httpbin.org/anything", apiUtils.Param{map[string]string{"hello": "world"}}).toJson()
	//log.Print(json)
}
