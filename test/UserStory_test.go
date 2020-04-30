package test

import (
	"github.com/yanyundata/woodpecker/userstory"
	"github.com/yanyundata/woodpecker/utils"
	"testing"
)

func TestUserStory(t *testing.T) {
	us := userstory.New()
	us.Tell("用例1", func(session userstory.Session) bool {
		session["abc"] = "123"
		return true
	}).Tell("用例2", func(session userstory.Session) bool {
		println(utils.DataAdapter{Data: session["abc"]}.ToString())
		return true
	}).ThatSAll()
}
