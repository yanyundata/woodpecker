package userstory

import (
	"fmt"
	"testing"
)

func TestUserStory(t *testing.T) {
	us := New()
	us.Tell("用例1", TestCase(func() {
		fmt.Println("YES")
	})).Tell("用例2", TestCase(func() {
		fmt.Println("YES,YES")
	})).ThatSAll()
}
