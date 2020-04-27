package main

import (
	"github.com/yanyundata/woodpecker/apiUtils"
	"github.com/yanyundata/woodpecker/demo/dto"
	"log"
)

func main() {
	//a=1;b=c;test1=100;test2=2344

	gadata := apiUtils.GetA("/admin/test", "?test=123").ToString()
	if gadata == "123ok" {
		log.Print("GetA OK!!!\n")
	}

	gbdata := apiUtils.GetB("/admin/test", "/100/200").ToJson()
	list := gbdata.Find("data").ToJsonObjectList()
	for e := list.Front(); e != nil; e = e.Next() {
		log.Print("GetB OK!!!\n")
	}

	padata := apiUtils.PostA("/admin/test/saveC2", "?test1=aaaa&test2=bbbb").ToJson()
	if padata.Find("data.test2").ToString() == "bbbb" {
		log.Print("PostA OK!!!\n")
	}

	pcdata := apiUtils.PostC("/admin/test/saveC1", &dto.DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	if pcdata.Find("data.test1").ToString() == "aaaa" {
		log.Print("PostC OK!!!\n")
	}

	putadata := apiUtils.PutA("/admin/test/update1", "?test1=aaaa&test2=bbbb").ToJson()
	if putadata.Find("data.test2").ToString() == "bbbb" {
		log.Print("PutA OK!!!\n")
	}

	putcdata := apiUtils.PutC("/admin/test/update2", &dto.DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	if putcdata.Find("data.test1").ToString() == "aaaa" {
		log.Print("PutC OK!!!\n")
	}

	dbdata := apiUtils.DeleteB("/admin/test/delete", "/100").ToJson()
	if dbdata.Find("msg").ToString() == "操作成功" {
		log.Print("DeleteB OK!!!\n")
	}

}
