package test

import (
	"testing"
)

func TestUserStory(t *testing.T) {
	//conn, err := grpc.Dial("121.22.244.194:38099", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("Connect failed: %v", err)
	//}
	//
	//c := pb.NewMyServiceClient(conn)
	//
	//hr := &pb.HelloRequest{Name: "tomwu"}
	//r, err := c.SayHi(context.Background(), hr)
	//
	//if err != nil {
	//	log.Fatalf("Fail to call SayHi：%v", err)
	//}
	//log.Printf("Result：%s ", r.Message)

	//us := userstory.New("UserStory Demo1")
	//us.Tell("测试GetA接口", func(session userstory.Session) {
	//	gadata := apiap.GetA(baseUrl+"/test", "test=123").ToString()
	//	if gadata == "123ok" {
	//		session["gadata"] = gadata
	//	} else {
	//		us.Error("数据读取失败")
	//	}
	//}).Tell("测试Session读取", func(session userstory.Session) {
	//	if session["gadata"].(string) != "123ok" {
	//		us.Error("session读取失败")
	//	}
	//}).Tell("假装FAIL", func(session userstory.Session) {
	//	us.Error("～～～")
	//}).Tell("空用例", func(session userstory.Session) {
	//}).ThatSAll()
}
