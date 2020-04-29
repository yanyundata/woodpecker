package test

import (
	"github.com/yanyundata/woodpecker/apiap"
	"testing"
)

type DemoDto struct {
	Test1 string `JsonObject:"test1"`
	Test2 string `JsonObject:"test2"`
}

func TestApiUtils(t *testing.T) {
	gadata := apiap.GetA("/admin/test", "?test=123").ToString()
	if gadata == "123ok" {
		t.Log("GetA OK!!!\n")
	}

	gbdata := apiap.GetB("/admin/test", "/100/200").ToJson()
	list := gbdata.Find("data").ToJsonObjectList()
	for e := list.Front(); e != nil; e = e.Next() {
		t.Log("GetB OK!!!\n")
	}

	padata := apiap.PostA("/admin/test/saveC2", "?test1=aaaa&test2=bbbb").ToJson()
	if padata.Find("data.test2").ToString() == "bbbb" {
		t.Log("PostA OK!!!\n")
	}

	pcdata := apiap.PostC("/admin/test/saveC1", &DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	if pcdata.Find("data.test1").ToString() == "aaaa" {
		t.Log("PostC OK!!!\n")
	}

	putadata := apiap.PutA("/admin/test/update1", "?test1=aaaa&test2=bbbb").ToJson()
	if putadata.Find("data.test2").ToString() == "bbbb" {
		t.Log("PutA OK!!!\n")
	}

	putcdata := apiap.PutC("/admin/test/update2", &DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	if putcdata.Find("data.test1").ToString() == "aaaa" {
		t.Log("PutC OK!!!\n")
	}

	dbdata := apiap.DeleteB("/admin/test/delete", "/100").ToJson()
	if dbdata.Find("msg").ToString() == "操作成功" {
		t.Log("DeleteB OK!!!\n")
	}
}
