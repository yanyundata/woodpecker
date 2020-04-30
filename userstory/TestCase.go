package userstory

type ITestCase interface {
	Test(session Session) bool
}

type TestCase func(session Session) bool

func (ts TestCase) Test(session Session) bool {
	return ts(session)
}
