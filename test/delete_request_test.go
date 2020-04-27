package test

import (
	"github.com/yanyundata/woodpecker/apiUtils"
	"log"
	"testing"
)

//http://localhost:9000/test/delete?id=xx
func TestDeleteOne(t *testing.T) {

	deletedata := apiUtils.DeleteA("/test/delete",
		make(apiUtils.PathParam).Set("id", "1")).ToJson()
	if deletedata.Find("data").ToString()=="1" {
		log.Print("DeleteA is Ok!")
	}
}
//http://localhost:9000/test/delete/xxx
func TestDeleteTwo(t *testing.T)  {
	deletedata := apiUtils.DeleteB("/test/delete", make(apiUtils.PathParam).Set("id", "2")).ToJson()
	if deletedata.Find("data").ToString()=="2" {
		log.Println("DeleteB is Ok!")
	}

}