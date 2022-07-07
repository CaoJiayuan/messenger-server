package socket

import (
	"github.com/enorith/http/contracts"
	"github.com/enorith/http/pipeline"
)

type AuthMiddleware struct {
	Token string
}

func (am AuthMiddleware) Handle(request contracts.RequestContract, next pipeline.PipeHandler) contracts.ResponseContract {

	return next(request)
}
