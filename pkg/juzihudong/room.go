package juzihudong

import (
	"encoding/json"
	"github.com/fatelei/juzihudong-sdk/pkg/model"
	"github.com/fatelei/juzihudong-sdk/pkg/transport"
	"strconv"
)

type RoomApi struct {
	Transport *transport.Transport
}

func NewRoomApi(endpoint string, token string) *RoomApi {
	transport := &transport.Transport{Endpoint:endpoint,Token:token,}
	roomApi := &RoomApi{Transport:transport,}
	return roomApi
}

type RoomListReponse struct {
	Code int64 `json:"code,omitempty"`
	Data *[]model.Room `json:"room,omitempty"`
}


func (p *RoomApi) GetRooms(current int, pageSize int) *RoomListReponse {
	param := make(map[string]string)
	param["current"] = strconv.Itoa(current)
	param["pageSize"] = strconv.Itoa(pageSize)
	body, err := p.Transport.Get("/room/list", &param)
	if err != nil {
		panic(err)
	}
	resp := RoomListReponse{}
	if err := json.Unmarshal(body, &resp); err == nil {
		return &resp
	}
	return nil
}