package main

import (
	"fmt"
	"flag"
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
)


func main() {
	var endpoint string
	var token string
	var wxId string

	flag.StringVar(&endpoint, "endpoint", "https://ex-api.botorange.com", "api endpoint")
	flag.StringVar(&token, "token", "","api token")
	flag.StringVar(&wxId,"wxid","","weixin id")
	flag.Parse()
	if len(token) == 0 {
		fmt.Println("token is required")
		return
	}

	if len(wxId) == 0 {
		fmt.Println("wxid is required")
		return
	}

	roomApi := juzihudong.NewRoomApi(endpoint, token)
	resp := roomApi.GetRooms(0, 10)
	fmt.Printf("%v\n", resp)
}
