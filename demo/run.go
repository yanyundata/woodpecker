package main

import (
	"fmt"
	"github.com/yanyundata/woodpecker/apiUtils"
	"log"
)

func main() {
	json := apiUtils.GetA("https://yan-yun.com:38085/lvyuan/admin/test", apiUtils.Param{Data: map[string]string{"test": "123"}}).ToString()
	if json == "123ok" {
		fmt.Print("OK!!!")
	}
	log.Print(json)
}
