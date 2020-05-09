package test

import (
	"github.com/yanyundata/woodpecker/apiap"
	"github.com/yanyundata/woodpecker/userstory"
	"log"
	"testing"
)

func TestUserStory(t *testing.T) {
	us := userstory.New("UserStory Demo1")
	us.Tell("测试GetA接口", func(session userstory.Session) bool {
		gadata := apiap.GetA("/admin/test", "?test=123").ToString()
		if gadata == "123ok" {
			session["gadata"] = gadata
			return true
		} else {
			return false
		}
	}).Tell("测试Session读取", func(session userstory.Session) bool {
		if session["gadata"].(string) == "123ok" {
			return true
		} else {
			return false
		}
	}).Tell("假装FAIL", func(session userstory.Session) bool {
		return false
	}).Tell("空用例", func(session userstory.Session) bool {
		return true
	}).ThatSAll()
}

func TestUserStoryCase(t *testing.T) {
	us := userstory.New("用户查询测试")
	us.Tell("测试具体查询", func(session userstory.Session) bool {
		gadata := apiap.GetA("https://yan-yun.com:38085/gotest/get/two/san/12", "").ToJson()
		log.Println("====================")
		log.Println(gadata)
		log.Println("====================")
		return true
	}).ThatSAll()
}
