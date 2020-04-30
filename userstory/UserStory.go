package userstory

type UserStory struct {
	testCaseMap map[string]ITestCase
	Session     map[string]interface{}
}

func New() UserStory {
	return UserStory{make(map[string]ITestCase), make(map[string]interface{})}
}

func (us UserStory) Tell(name string, testCase ITestCase) UserStory {
	us.testCaseMap[name] = testCase
	return us
}

func (us UserStory) ThatSAll() {
	for name := range us.testCaseMap {
		if us.testCaseMap[name].Test(us.Session) {
			println(name + " 测试通过")
		} else {
			println(name + " 测试失败")
		}
	}
}
