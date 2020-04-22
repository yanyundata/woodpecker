package main

import (
	"log"
)

func main() {
	m := map[string]string{"hello": "world"}
	//param := apiUtils.Param{m}
	//json := apiUtils.GetA("http://httpbin.org/anything", param).ToJson()
	log.Print(m)

}
