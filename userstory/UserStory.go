package userstory

import (
	"container/list"
)

type UserStory struct {
	topic            string
	executorDataList *list.List
}

func New(topic string) UserStory {
	return UserStory{topic: topic, executorDataList: list.New()}
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
