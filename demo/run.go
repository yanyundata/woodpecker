package main

import (
	"github.com/yanyundata/woodpecker/apiUtils"
	"log"
)

func main() {
	//json := apiUtils.GetA("http://httpbin.org/anything", apiUtils.Param{map[string]string{"hello": "world"}}).toJson()
	//log.Print(json)
	r := apiUtils.TestA()
	log.Print(r)
}
