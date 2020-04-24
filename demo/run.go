package main

import (
	"fmt"
	"github.com/yanyundata/woodpecker/apiUtils"
	"github.com/yanyundata/woodpecker/demo/dto"
	"log"
)

func main() {
	//a=1;b=c;test1=100;test2=2344

	getAData := apiUtils.GetA("/admin/test", make(apiUtils.PathParam).Set("test", "123")).ToString()
	if getAData == "123ok" {
		log.Print("GetA OK!!!\n")
	}

	getBData := apiUtils.GetB("/admin/test", make(apiUtils.PathParam).Set("test1", "100").Set("test2", "200")).ToJson()
	list := getBData.Find("data").ToJsonObjectList()
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(apiUtils.JsonObject).Find("test1").ToString())
	}

	PostC1Data := apiUtils.PostC1("/admin/test/saveC1", &dto.DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	c1v := PostC1Data.Find("data.test1").ToString()
	fmt.Print(c1v)
}
