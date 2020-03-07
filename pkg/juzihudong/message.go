package juzihudong

import "github.com/fatelei/juzihudong-sdk/pkg/transport"

type MessageApi struct {
	Transport *transport.Transport
}


func NewMessageApi(endpoint string, token string) *MessageApi {
	transport := &transport.Transport{Endpoint:endpoint,Token:token}
	messageApi := &MessageApi{Transport:transport}
	return messageApi
}


func (p *MessageApi) SendTextMessage(chatId string, content string) bool {
	body := make(map[string]interface{})
	body["chatId"] = chatId
	body["messageType"] = 1
	body["payload"] = map[string]string{"text": content}
	_, err := p.Transport.Post("/message/send", body)
	if err != nil {
		return false
	}
	return true
}