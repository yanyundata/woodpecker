package userstory

import (
	"container/list"
	"log"
	"strconv"
	"time"
)

type UserStory struct {
	topic         string
	testCaseMap   map[string]ITestCase
	testCaseIndex *list.List
	session       Session
}

func New(topic string) UserStory {
	return UserStory{topic, make(map[string]ITestCase), list.New(), make(Session)}
}

func (us UserStory) Tell(name string, testCaseFun func(session Session) bool) UserStory {
	us.testCaseMap[name] = TestCase(testCaseFun)
	us.testCaseIndex.PushBack(name)
	return us
}

func (us UserStory) ThatSAll() {
	var noFail = true
	var sunLp = int64(0)
	for e := us.testCaseIndex.Front(); e != nil; e = e.Next() {
		name := e.Value.(string)
		if noFail {
			startTime := time.Now().Unix()
			var isPass = us.testCaseMap[name].Test(us.session)
			endTime := time.Now().Unix()

			lp := endTime - startTime
			lpStr := strconv.FormatInt(lp, 10)
			sunLp = sunLp + lp

			if isPass {
				log.Println("用例【" + name + "】" + " PASS " + lpStr + "ms")
			} else {
				log.Println("用例【" + name + "】" + " FAIL " + lpStr + "ms")
				noFail = false
			}
		} else {
			log.Println("用例【" + name + "】" + " SKIP " + "0ms")
			continue
		}
	}

	if noFail {
		log.Println("用户故事【" + us.topic + "】" + " PASS " + strconv.FormatInt(sunLp, 10) + "ms")
	} else {
		log.Println("用户故事【" + us.topic + "】" + " FAIL " + strconv.FormatInt(sunLp, 10) + "ms")
	}
}
