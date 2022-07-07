package socket

import (
	"fmt"
	"log"

	"github.com/CaoJiayuan/messenger"
	"github.com/enorith/container"
	"github.com/enorith/framework"
	"github.com/enorith/http/contracts"
)

var Server *messenger.Server

type Config struct {
	Port  int    `yaml:"port" default:"3003"`
	Token string `yaml:"token" default:""`
}

type Service struct {
}

//Register service when app starting, before http server start
// you can configure service, prepare global vars etc.
// running at main goroutine
func (s *Service) Register(app *framework.App) error {

	var conf Config
	app.Configure("messenger", &conf)

	Server = messenger.NewServer()

	Server.Cors()
	app.Bind(func(ioc container.Interface) {
		ioc.BindFunc(&messenger.Server{}, func(c container.Interface) (interface{}, error) {
			return Server, nil
		}, true)
	})

	app.Daemon(func(exit chan struct{}) {
		log.Println("[socket] socket server start at port", conf.Port)

		go func() {
			Server.Serve(fmt.Sprintf(":%d", conf.Port))
		}()

		<-exit
	})

	return nil
}

//Lifetime container callback
// usually register request lifetime instance to IoC-Container (per-request unique)
// this function will run before every request handling
func (s *Service) Lifetime(ioc container.Interface, request contracts.RequestContract) {
}

func NewService() *Service {
	return &Service{}
}
