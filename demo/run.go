package main

import (
	"github.com/yanyundata/woodpecker/apiUtils"
	"log"
)

func main() {
	json := apiUtils.GetA("http://httpbin.org/anything", apiUtils.Param{Data: map[string]string{"hello": "world"}}).ToJson()
	log.Print(json)
}
