package response

import (
	"github.com/siuvlqnm/bookmark/model"
)

type CliUserResponse struct {
	User model.CliUser `json:"user"`
}

type CliLoginResponse struct {
	User      model.CliUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
