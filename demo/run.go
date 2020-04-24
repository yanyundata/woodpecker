package main

import (
	"github.com/yanyundata/woodpecker/apiUtils"
	"log"
)

func main() {
	getAData := apiUtils.GetA("https://yan-yun.com:38085/lvyuan/admin/test", apiUtils.Param{Data: map[string]string{"test": "123"}}).ToString()
	if getAData == "123ok" {
		log.Print("GetA OK!!!\n")
	}

	getBData := apiUtils.GetB("https://yan-yun.com:38085/lvyuan/admin/test", apiUtils.Param{Data: map[string]string{"test1": "100", "test2": "2"}}).ToJson()
	list := getBData.Find("data").ToList()
	for _, object := range list {
		value := object.(map[string]interface{})

		var test1 = value["test1"].(string)
		if "100" == test1 {
			log.Print("GetB OK!!!\n")
		}
	}

}
