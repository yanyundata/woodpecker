package userstory

type ITestCase interface {
	Test(session BusyBox)
}

type TestCase func(busybox BusyBox)

func (ts TestCase) Test(busybox BusyBox) {
	ts(busybox)
}
