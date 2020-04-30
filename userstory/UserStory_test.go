package userstory

import (
	"log"
	"testing"
)

type Case1 struct {
}

func (c Case1) Test() {
	log.Println("YES")
}

func TestUserStory(t *testing.T) {
	us := New()
	us.Tell("用例1", Case1{}).ThatsAll()
}
