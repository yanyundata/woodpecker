package main

import (
	"github.com/yanyundata/woodpecker/userstory"
)

func main() {
	rootGroup := userstory.NewGroup("快车道APP测试")

	ddzc := userstory.NewGroup("注册登陆")

	rootGroup.AddSubGroup(ddzc)
	shoujihaozhuce := userstory.NewGroup("手机号注册登陆")
	weixindenglu := userstory.NewGroup("微信登陆")
	wangjimima := userstory.NewGroup("忘记密码")
	ddzc.AddSubGroup(shoujihaozhuce).AddSubGroup(weixindenglu).AddSubGroup(wangjimima)

	grzx := userstory.NewGroup("个人中心")
	sy := userstory.NewGroup("首页")
	tj := userstory.NewGroup("体检")
	dt := userstory.NewGroup("地图")
	jcyu := userstory.NewGroup("检车预约")

	rootGroup.AddSubGroup(ddzc).AddSubGroup(grzx).AddSubGroup(sy).AddSubGroup(tj).AddSubGroup(dt).AddSubGroup(jcyu)

	us := userstory.New("首页").Tell("显示视频", func(busybox userstory.BusyBox) {

	}).Tell("显示新闻", func(busybox userstory.BusyBox) {

	}).Tell("显示。。。", func(busybox userstory.BusyBox) {

	})
	shoujihaozhuce.AddUserStory(&us)

	te := userstory.TestExecutor{}
	te.Test(rootGroup)
}
