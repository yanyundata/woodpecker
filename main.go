package main

import (
	"context"
	"fmt"
	"github.com/yanyundata/woodpecker/userstory"
	pb "gitlab.com/yanyundata/module/go/proto-cicd-demo"
	"log"
)

func Hello(ch chan int) {
	fmt.Println("hello everybody , I'm asong")
	ch <- 1
}

func main() {
	us := userstory.New("Grpc Test Demo")
	us.Tell("请求Grpc", func(busybox userstory.BusyBox) {
		c := pb.NewMyServiceClient(busybox.GrpcHelper().Conn("121.22.244.194:38099"))
		r, _ := c.SayHi(context.Background(), &pb.HelloRequest{Name: "tomwu"})
		busybox.Session()["result"] = r.Message
	}).Tell("打印结果", func(busybox userstory.BusyBox) {
		log.Printf("Result：%s ", busybox.Session()["result"])
	}).ThatSAll()
}
