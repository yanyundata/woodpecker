package main

import (
	"fmt"
	"github.com/yanyundata/woodpecker/apiUtils"
	"log"
)

func main() {
	getAData := apiUtils.GetA("https://yan-yun.com:38085/lvyuan/admin/test", apiUtils.Param{Data: map[string]string{"test": "123"}}).ToString()
	if getAData == "123ok" {
		log.Print("GetA OK!!!\n")
	}

	getBData := apiUtils.GetB("https://yan-yun.com:38085/lvyuan/admin/test", apiUtils.Param{Data: map[string]string{"test1": "1", "test2": "2"}}).ToJson()
	fmt.Print(getBData)
}
