package userstory

import (
	"container/list"
	"fmt"
	"github.com/fatih/color"
	"strconv"
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

	busyBox := newBusyBox()

	for e := l.Front(); e != nil; e = e.Next() {
		tc := e.Value.(ExecutorData)

		fmt.Println(tc.name)
	}

	var noFail = true
	var sunLp = int64(0)
	for e := l.Front(); e != nil; e = e.Next() {
		ed := e.Value.(ExecutorData)
		name := ed.name

		sb := SandBox{testCaseName: name, busyBox: busyBox}

		if noFail {
			sb.run(ed.testCase)
			lpStr := strconv.FormatInt(sb.timeCost, 10)
			sunLp = sunLp + sb.timeCost

			if sb.pass {
				color.Green("用例【" + name + "】" + " PASS " + lpStr + "ms")
			} else {
				color.Red("用例【" + name + "】" + " FAIL " + sb.msg + lpStr + "ms")
				noFail = false
			}
		} else {
			color.Yellow("用例【" + name + "】" + " SKIP " + "0ms")
			continue
		}
	}

	//if noFail {
	//	color.Green("用户故事【" + us + "】" + " PASS " + strconv.FormatInt(sunLp, 10) + "ms")
	//} else {
	//	color.Red("用户故事【" + us.topic + "】" + " FAIL " + strconv.FormatInt(sunLp, 10) + "ms")
	//}

	return Report{}
}
