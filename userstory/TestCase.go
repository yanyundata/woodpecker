package userstory

type ITestCase interface {
	Test() bool
}

type TestCase func() bool

func (ts TestCase) Test() bool {
	return ts()
}
