package api

import (
	"messenger-server/internal/pkg/types"
	"strings"

	"github.com/CaoJiayuan/messenger"
	"github.com/buger/jsonparser"
	"github.com/enorith/http/content"
)

var (
	defaultEvent = "message"
	evStep       = "::"
)

type Broadcast struct {
	content.Request
	Channels []string `input:"channels" validate:"required"`
	Payload  []byte   `input:"payload" validate:"required"`
}

type IoHandler struct {
}

func (i IoHandler) Broadcast(b Broadcast, srv *messenger.Server) (types.Message, error) {
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
			srv.Broadcast(ev, types.JsonString(b.Payload), ce[0])
		} else {
			// string
			srv.Broadcast(ev, string(b.Payload), ce[0])
		}
	}

	return "ok", nil
}
