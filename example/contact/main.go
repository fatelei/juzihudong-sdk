package main

import (
	"fmt"
	"flag"
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
)


func main() {
	var endpoint string
	var token string
	var wxid string

	flag.StringVar(&endpoint, "endpoint", "https://ex-api.botorange.com", "api endpoint")
	flag.StringVar(&token, "token", "","api token")
	flag.StringVar(&wxid, "wxid", "", "wxid")

	flag.Parse()
	if len(token) == 0 {
		fmt.Println("token is required")
		return
	}

	contactApi := juzihudong.NewContactApi(endpoint, token)
	resp := contactApi.GetContact(0, 10, wxid)
	fmt.Printf("%+v\n", resp.Data)
}
