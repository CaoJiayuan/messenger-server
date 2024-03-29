package requests

import (
	"messenger-server/internal/pkg/auth"
	"github.com/enorith/http/content"
)

type LoginRequest struct {
	content.Request
}

func (lr LoginRequest) Rules() map[string][]interface{} {
	return map[string][]interface{}{
		auth.UsernameField: {"required"},
		auth.PasswordField: {"required"},
	}
}
