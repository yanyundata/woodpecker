package userstory

type ITestCase interface {
	Test(session map[string]interface{}) bool
}

type TestCase func(session map[string]interface{}) bool

func (ts TestCase) Test(session map[string]interface{}) bool {
	return ts(session)
}
