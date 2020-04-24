package main

import (
	"fmt"
	"github.com/yanyundata/woodpecker/apiUtils"
	"log"
)

func main() {
	getAData := apiUtils.GetA("/admin/test", apiUtils.Param{Data: map[string]string{"test": "123"}}).ToString()
	if getAData == "123ok" {
		log.Print("GetA OK!!!\n")
	}

	getBData := apiUtils.GetB("/admin/test", apiUtils.Param{Data: map[string]string{"test1": "100", "test2": "2"}}).ToJson()
	list := getBData.Find("data").ToJsonObjectList()
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(apiUtils.JsonObject).Find("test1").ToString())
	}

}
