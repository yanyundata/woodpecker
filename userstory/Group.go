package userstory

import (
	"container/list"
)

type Group struct {
	name          string
	subGroupList  []*Group
	userStoryList []*UserStory
}

func (g Group) getExecutorData(executorDataList *list.List) {
	for i := 0; i < len(g.subGroupList); i++ {
		g.subGroupList[i].getExecutorData(executorDataList)
	}

	for i := 0; i < len(g.userStoryList); i++ {
		g.userStoryList[i].getExecutorData(executorDataList)
	}
}

func NewGroup(name string) *Group {
	return &Group{name: name, subGroupList: []*Group{}, userStoryList: []*UserStory{}}
}

func (g *Group) AddSubGroup(group *Group) *Group {
	g.subGroupList = append(g.subGroupList, group)

	return g
}

func (g *Group) AddUserStory(us *UserStory) *Group {
	g.userStoryList = append(g.userStoryList, us)

	return g
}
