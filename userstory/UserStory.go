package userstory

import "container/list"

type UserStory struct {
	topic   string
	busyBox BusyBox

	executorDataList *list.List
}

func New(topic string) UserStory {
	return UserStory{topic: topic, busyBox: newBusyBox(), executorDataList: list.New()}
}

func (us UserStory) Tell(name string, testCaseFun func(busybox BusyBox)) UserStory {
	us.executorDataList.PushBack(NewExecutorData(name, TestCase(testCaseFun)))
	return us
}

func (us UserStory) Error(errMsg string) {
	panic(errMsg)
}

func (us UserStory) getExecutorData(executorDataList *list.List) {
	executorDataList.PushBackList(us.executorDataList)
}

//func (us UserStory) ThatSAll() {
//	var noFail = true
//	var sunLp = int64(0)
//	for e := us.testCaseIndex.Front(); e != nil; e = e.Next() {
//		name := e.Value.(string)
//		sb := SandBox{testCaseName: name, busyBox: us.busyBox}
//		if noFail {
//			sb.run(us.testCaseMap[name])
//			lpStr := strconv.FormatInt(sb.timeCost, 10)
//			sunLp = sunLp + sb.timeCost
//
//			if sb.pass {
//				color.Green("用例【" + name + "】" + " PASS " + lpStr + "ms")
//			} else {
//				color.Red("用例【" + name + "】" + " FAIL " + sb.msg + lpStr + "ms")
//				noFail = false
//			}
//		} else {
//			color.Yellow("用例【" + name + "】" + " SKIP " + "0ms")
//			continue
//		}
//	}
//
//	if noFail {
//		color.Green("用户故事【" + us.topic + "】" + " PASS " + strconv.FormatInt(sunLp, 10) + "ms")
//	} else {
//		color.Red("用户故事【" + us.topic + "】" + " FAIL " + strconv.FormatInt(sunLp, 10) + "ms")
//	}
//
//}
