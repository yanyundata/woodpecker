package userstory

import (
	"log"
	"time"
)

type SandBox struct {
	testCaseName string
	session      Session

	pass     bool
	timeCost int64
	msg      string
}

func (sb *SandBox) run(testCase ITestCase) {
	startTime := time.Now().Unix()
	defer func() {
		sb.timeCost = time.Now().Unix() - startTime
		if err := recover(); err != nil {
			switch err.(type) {
			case string:
				sb.msg = err.(string)
			default:
				log.Println(err)
			}

			sb.pass = false
		} else {
			sb.pass = true
		}
	}()

	testCase.Test(sb.session)
}
