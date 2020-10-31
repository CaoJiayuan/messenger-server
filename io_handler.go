package main

import (
	"github.com/CaoJiayuan/messenger"
	"github.com/buger/jsonparser"
	"github.com/enorith/framework/http/content"
	"strings"
)


var (
	defaultEvent = "message"
	evStep       = "::"
)


type Broadcast struct {
	content.Request
	Channels []string `input:"channels"`
	Payload  []byte   `input:"payload"`
}

type IoHandler struct {
	Srv *messenger.Server
}

func(i IoHandler) Broadcast(b Broadcast) (Message,error) {
	s, err := jsonparser.ParseString(b.Payload)
	if err != nil {
		return "", err
	}

	for _, channel := range b.Channels {
		ce := strings.Split(channel, evStep)
		var ev string
		if len(ce) > 1 {
			ev = ce[1]
		} else {
			ev = defaultEvent
		}
		if strings.HasPrefix(s, "{") || strings.HasPrefix(s, "[") {
			// assume json
			i.Srv.Broadcast(ev, JsonString(b.Payload), ce[0])
		} else {
			// string
			i.Srv.Broadcast(ev, string(b.Payload), ce[0])
		}
	}
	return "ok", nil
}