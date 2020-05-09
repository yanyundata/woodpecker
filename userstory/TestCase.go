package userstory

type ITestCase interface {
	Test(session Session)
}

type TestCase func(session Session)

func (ts TestCase) Test(session Session) {
	ts(session)
}
