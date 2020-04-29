package test

import (
	"github.com/yanyundata/woodpecker/apiUtils"
	"testing"
)

type DemoDto struct {
	Test1 string `json:"test1"`
	Test2 string `json:"test2"`
}

func TestApiUtils(t *testing.T) {
	gadata := apiUtils.GetA("/admin/test", "?test=123").ToString()
	if gadata == "123ok" {
		t.Log("GetA OK!!!\n")
	}

	gbdata := apiUtils.GetB("/admin/test", "/100/200").ToJson()
	list := gbdata.Find("data").ToJsonObjectList()
	for e := list.Front(); e != nil; e = e.Next() {
		t.Log("GetB OK!!!\n")
	}

	padata := apiUtils.PostA("/admin/test/saveC2", "?test1=aaaa&test2=bbbb").ToJson()
	if padata.Find("data.test2").ToString() == "bbbb" {
		t.Log("PostA OK!!!\n")
	}

	pcdata := apiUtils.PostC("/admin/test/saveC1", &DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	if pcdata.Find("data.test1").ToString() == "aaaa" {
		t.Log("PostC OK!!!\n")
	}

	putadata := apiUtils.PutA("/admin/test/update1", "?test1=aaaa&test2=bbbb").ToJson()
	if putadata.Find("data.test2").ToString() == "bbbb" {
		t.Log("PutA OK!!!\n")
	}

	putcdata := apiUtils.PutC("/admin/test/update2", &DemoDto{Test1: "aaaa", Test2: "bbb"}).ToJson()
	if putcdata.Find("data.test1").ToString() == "aaaa" {
		t.Log("PutC OK!!!\n")
	}

	dbdata := apiUtils.DeleteB("/admin/test/delete", "/100").ToJson()
	if dbdata.Find("msg").ToString() == "操作成功" {
		t.Log("DeleteB OK!!!\n")
	}
}
