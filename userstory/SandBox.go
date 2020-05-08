package userstory

import (
	"github.com/yanyundata/woodpecker/utils"
	"time"
)

type SandBox struct {
	testCaseName string
	session      Session

	pass     bool
	timeCost int64
	msg      string
}

func (sb *SandBox) Run(testCase ITestCase) {
	startTime := time.Now().Unix()
	defer func() {
		sb.timeCost = time.Now().Unix() - startTime
		if err := recover(); err != nil {
			sb.msg = utils.DataAdapter{Data: err}.ToString()
			sb.pass = false
		} else {
			sb.pass = true
		}
	}()

	testCase.Test(sb.session)
}
