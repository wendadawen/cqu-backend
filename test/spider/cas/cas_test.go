package cas

import (
	"cqu-backend/src/spider/cas"
	"cqu-backend/test"
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	auth := cas.NewAuth(test.StudId1, test.CasPwd1)
	auth.Login()
	Url := "https://self.cqu.edu.cn/getUser"
	client := auth.GetClient()
	post, err := client.R().Post(Url)
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println(post.String())
}
