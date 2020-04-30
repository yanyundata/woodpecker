package userstory

import "container/list"

type UserStory struct {
	testCaseMap   map[string]ITestCase
	testCaseIndex *list.List
	session       Session
}

func New() UserStory {
	return UserStory{make(map[string]ITestCase), list.New(), make(Session)}
}

func (us UserStory) Tell(name string, testCaseFun func(session Session) bool) UserStory {
	us.testCaseMap[name] = TestCase(testCaseFun)
	us.testCaseIndex.PushBack(name)
	return us
}

func (us UserStory) ThatSAll() {
	for e := us.testCaseIndex.Front(); e != nil; e = e.Next() {
		name := e.Value.(string)
		if us.testCaseMap[name].Test(us.session) {
			println("用例【" + name + "】" + " PASS")
		} else {
			println("用例【" + name + "】" + " FAIL")
		}
	}
}
