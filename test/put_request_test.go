package test

import (
	"github.com/yanyundata/woodpecker/apiUtils"
	"log"
	"testing"
)

//测试put请求: http://localhost:9000/test/put/{one}/{two}
func TestPutA(t *testing.T) {
	putdata := apiUtils.PutA("/test/put",
		make(apiUtils.PathParam).
			Set("one", "a").
			Set("two", "b")).
		ToJson()

	if putdata.Find("data").ToString() == "one:a- two:b" {
		log.Println("PUTA is Ok!")
	}
}

//put 请求 请求路径参数
//put: https://ip:port/test/put/two?param1=xxx&param2=xxxx
func TestPutB(t *testing.T) {
	putdata := apiUtils.PutB("/test/put/two",
		make(apiUtils.PathParam).
			Set("param1", "one").Set("param2", "two")).
		ToJson()
	if putdata.Find("data.param1").ToString() == "one" {
		log.Println("PutB is Ok")
	}
}

//put 请求 请求体为Json
// put : https://ip:port/test/put/three
type TestDto struct {
	Param1 string `json:"param1"`
	Param2 string `json:"param2"`
}
func TestPutC(t *testing.T) {
	putdata := apiUtils.PutC("/test/put/three",
		TestDto{Param1: "参数一", Param2: "参数二"}).
		ToJson()
	if putdata.Find("data.param1").ToString()=="参数一" {
		log.Println("PutC is Ok")
	}
}

