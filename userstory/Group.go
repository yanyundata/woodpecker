package userstory

type Group struct {
	name          string
	subGroupList  []Group
	userStoryList []UserStory
}

func NewGroup(name string) *Group {
	return &Group{name: name, subGroupList: []Group{}, userStoryList: []UserStory{}}
}

func (g Group) AddSubGroup(group Group) *Group {
	g.subGroupList = append(g.subGroupList, group)

	return &g
}
