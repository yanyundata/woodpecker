package userstory

import (
	"container/list"
	"fmt"
)

type ExecutorData struct {
	name     string
	testCase ITestCase
}

func NewExecutorData(name string, testCase ITestCase) ExecutorData {
	return ExecutorData{name: name, testCase: testCase}
}

type ITestAble interface {
	getExecutorData(executorDataList *list.List)
}

type Report struct {
}

type TestExecutor struct {
	report Report
}

func (ta TestExecutor) Test(us ITestAble) Report {
	l := list.New()
	us.getExecutorData(l)

	fmt.Println(l.Len())
	return Report{}
}
