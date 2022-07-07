package types

import (
	"encoding/json"
	"net/http"
)

type Message string

func (m Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"status":  200,
		"message": string(m),
	})
}

type Status int

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"status":  int(s),
		"message": http.StatusText(int(s)),
	})
}

type JsonString []byte

func (j JsonString) MarshalJSON() ([]byte, error) {
	return j, nil
}
