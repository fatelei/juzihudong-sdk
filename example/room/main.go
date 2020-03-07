package main

import (
	"fmt"
	"flag"
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
)


func main() {
	var endpoint string
	var token string

	flag.StringVar(&endpoint, "endpoint", "https://ex-api.botorange.com", "api endpoint")
	flag.StringVar(&token, "token", "","api token")
	flag.Parse()
	if len(token) == 0 {
		fmt.Println("token is required")
		return
	}
	roomApi := juzihudong.NewRoomApi(endpoint, token)
	resp := roomApi.GetRooms(0, 10)
	fmt.Printf("%+v\n", resp.Data)
}
