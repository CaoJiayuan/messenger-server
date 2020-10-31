package main

import (
	"encoding/json"
	"github.com/CaoJiayuan/messenger"
	env "github.com/enorith/environment"
	"github.com/enorith/framework"
	"github.com/enorith/framework/http"
	"github.com/enorith/framework/http/router"
	http2 "net/http"
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
		"message": http2.StatusText(int(s)),
	})
}

type JsonString []byte

func (j JsonString) MarshalJSON() ([]byte, error) {
	return j, nil
}

func main() {
	srv, _ := messenger.NewServer()
	srv.Cors()
	srv.RegisterEvents().ServeIo()
	masterKey := env.GetString("MASTER_KEY", "master")
	exp := env.GetInt("JWT_TTL", 24*60*60)
	jwtKey := env.GetString("JWT_SECRET", "somerandomstring")
	defer srv.Close()
	framework.Serve(":3000", func(ro *router.Wrapper, k *http.Kernel) {
		k.OutputLog = true
		k.KeepAlive()
		k.Handler = http.HandlerNetHttp
		k.SetMiddlewareGroup(map[string][]http.RequestMiddleware{
			"auth": {AuthMiddleware{JwtKey: jwtKey}},
		})
		ro.RegisterAction(router.ANY, "/socket.io", srv)
		ro.Post("/login", AuthHandler{masterKey, exp, jwtKey}.Login)
		ro.Post("/broadcast", IoHandler{srv}.Broadcast).Middleware("auth")

	}, framework.StandardAppStructure{BasePath: "."})
}
