package main

import (
	"flag"
	"fmt"
	"github.com/fatelei/juzihudong-sdk/pkg/juzihudong"
)


func main() {
	var endpoint string
	var token string
	var messageId string
	var chatId string

	flag.StringVar(&endpoint, "endpoint", "https://ex-api.botorange.com", "api endpoint")
	flag.StringVar(&token, "token", "","api token")
	flag.StringVar(&messageId, "message_id", "","message id")
	flag.StringVar(&chatId, "chat_id", "","chat id")
	flag.Parse()
	if len(token) == 0 || len(messageId) == 0 || len(chatId) == 0{
		fmt.Println("token | message id | chat id is required")
		return
	}
	messageApi := juzihudong.NewMessageApi(endpoint, token)
	resp := messageApi.GetArtworkImage(chatId, messageId)
	fmt.Printf("%+v\n", resp.Data)
}
