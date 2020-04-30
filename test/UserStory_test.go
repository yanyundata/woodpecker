package test

import (
	"github.com/yanyundata/woodpecker/userstory"
	"testing"
)

func TestUserStory(t *testing.T) {
	us := userstory.New()
	us.Tell("用例1", userstory.TestCase(func(session map[string]interface{}) bool {
		return true
	})).Tell("用例2", userstory.TestCase(func(session map[string]interface{}) bool {
		return true
	})).ThatSAll()
}
