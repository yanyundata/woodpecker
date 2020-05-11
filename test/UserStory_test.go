package test

import (
	"github.com/yanyundata/woodpecker/apiap"
	"github.com/yanyundata/woodpecker/userstory"
	"log"
	"testing"
)

func TestUserStory(t *testing.T) {
	us := userstory.New("UserStory Demo1")
	us.Tell("测试GetA接口", func(session userstory.Session) {
		gadata := apiap.GetA("/admin/test", "?test=123").ToString()
		if gadata == "123ok" {
			session["gadata"] = gadata
		} else {
			us.Error("数据读取失败")
		}
	}).Tell("测试Session读取", func(session userstory.Session) {
		if session["gadata"].(string) != "123ok" {
			us.Error("session读取失败")
		}
	}).Tell("假装FAIL", func(session userstory.Session) {
		us.Error("～～～")
	}).Tell("空用例", func(session userstory.Session) {
	}).ThatSAll()
}

func TestUserStoryCase(t *testing.T) {
	us := userstory.New("用户查询测试")
	us.Tell("测试具体查询", func(session userstory.Session) {
		gadata := apiap.GetA("https://yan-yun.com:38085/gotest/get/two/san/12", "").ToJson()
		log.Println("====================")
		log.Println(gadata)
		log.Println("====================")
	}).ThatSAll()
}
