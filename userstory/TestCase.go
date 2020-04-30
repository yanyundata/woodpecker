package userstory

type ITestCase interface {
	Test()
}

type TestCase func()

func (ts TestCase) Test() {
	ts()
}
