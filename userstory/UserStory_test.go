package userstory

import (
	"log"
	"testing"
)

type case1 struct {
}

func (c case1) Test() {
	log.Println("YES")
}

func TestUserStory(t *testing.T) {
	us := New()
	us.Tell("用例1", &case1{}).ThatsAll()
}
