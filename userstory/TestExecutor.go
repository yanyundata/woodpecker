package userstory

type ExecutorData struct {
	name     string
	tastCase ITestCase
}

type ITestAble interface {
	getTestCaseList() []ExecutorData
}

type Report struct {
}

type TestExecutor struct {
	report Report
}

func (ta TestExecutor) Test(us ITestAble) Report {

	return Report{}
}
