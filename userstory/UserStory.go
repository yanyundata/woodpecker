package userstory

import "container/list"

type UserStory struct {
	testCaseList *list.List
}

func New() UserStory {
	return UserStory{list.New()}
}

func (us UserStory) Tell(name string, testCase ITestCase) UserStory {
	us.testCaseList.PushBack(testCase)
	return us
}

func (us UserStory) ThatsAll() {
	for tc := us.testCaseList.Front(); tc != nil; tc = tc.Next() {
		tc.Value.(ITestCase).Test()
	}

}
