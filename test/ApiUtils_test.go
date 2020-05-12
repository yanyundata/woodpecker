package test

import (
	"github.com/yanyundata/woodpecker/apiap"
	"log"
	"testing"
)

type DemoDto struct {
	Test1 string `json:"test1"`
	Test2 string `json:"test2"`
}

var apiUtilsBaseUrl = "https://yan-yun.com:38085/lvyuan/admin/api"

type Model struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Test1 string `json:"test1"`
		Test2 string `json:"test2"`
	} `json:"data"`
	Time int64 `json:"time"`
}

func TestApiUtils(t *testing.T) {
	//gadata := apiap.Call(apiap.TypeA("GET", "/admin/test", "?test=123")).ToString()
	//if gadata == "123ok" {
	//	t.Log("GetA OK!!!\n")
	//}
	//
	//gbdata := apiap.Call(apiap.TypeB("GET", "/admin/test", "/100/200")).ToJson()
	//list := gbdata.Find("data").ToJsonObjectList()
	//for e := list.Front(); e != nil; e = e.Next() {
	//	t.Log("GetB OK!!!\n")
	//}
	//
	//padata := apiap.Call(apiap.TypeA("POST", "/admin/test/saveC2", "?test1=aaaa&test2=bbbb")).ToJson()
	//if padata.Find("data.test2").ToString() == "bbbb" {
	//	t.Log("PostA OK!!!\n")
	//}
	//
	//pcdata := apiap.Call(apiap.TypeC("POST", "/admin/test/saveC1", &DemoDto{Test1: "aaaa", Test2: "bbb"})).ToJson()
	//if pcdata.Find("data.test1").ToString() == "aaaa" {
	//	t.Log("PostC OK!!!\n")
	//}
	//
	//putadata := apiap.Call(apiap.TypeA("PUT", "/admin/test/update1", "?test1=aaaa&test2=bbbb")).ToJson()
	//if putadata.Find("data.test2").ToString() == "bbbb" {
	//	t.Log("PutA OK!!!\n")
	//}
	//
	//putcdata := apiap.Call(apiap.TypeC("PUT", "/admin/test/update2", &DemoDto{Test1: "aaaa", Test2: "bbb"})).ToJson()
	//if putcdata.Find("data.test1").ToString() == "aaaa" {
	//	t.Log("PutC OK!!!\n")
	//}
	//
	//dbdata := apiap.Call(apiap.TypeB("DELETE", "/admin/test/delete", "/100")).ToJson()
	//if dbdata.Find("msg").ToString() == "操作成功" {
	//	t.Log("DeleteB OK!!!\n")
	//}

	//gadata := apiap.GetA(apiUtilsBaseUrl+"/test", "?test=123").ToString()
	//if gadata == "123ok" {
	//	t.Log("GetA OK!!!\n")
	//}

	m := Model{}
	apiap.GetB(apiUtilsBaseUrl+"/test", "/100/200").ToModel(&m)
	log.Println(m.Data[0].Test1)

	//padata := apiap.PostA(apiUtilsBaseUrl+"/test/saveC2", "?test1=aaaa&test2=bbbb").ToJson()
	//if padata.Find("data.test2").ToString() == "bbbb" {
	//	t.Log("PostA OK!!!\n")
	//}
	//
	//pcdata := apiap.PostC(apiUtilsBaseUrl+"/test/saveC1", &DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	//if pcdata.Find("data.test1").ToString() == "aaaa" {
	//	t.Log("PostC OK!!!\n")
	//}
	//
	//putadata := apiap.PutA(apiUtilsBaseUrl+"/test/update1", "?test1=aaaa&test2=bbbb").ToJson()
	//if putadata.Find("data.test2").ToString() == "bbbb" {
	//	t.Log("PutA OK!!!\n")
	//}
	//
	//putcdata := apiap.PutC(apiUtilsBaseUrl+"/test/update2", &DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	//if putcdata.Find("data.test1").ToString() == "aaaa" {
	//	t.Log("PutC OK!!!\n")
	//}
	//
	//dbdata := apiap.DeleteB(apiUtilsBaseUrl+"/test/delete", "/100").ToJson()
	//if dbdata.Find("msg").ToString() == "操作成功" {
	//	t.Log("DeleteB OK!!!\n")
	//}
}
