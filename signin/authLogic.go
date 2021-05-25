package main

import (
	"context"
)

type authLogic struct {
}

func newAuthLogic() *authLogic {
	d := authLogic{}
	return &d
}

func (l *authLogic) isValidUsernamePassword(ctx context.Context, username string, password string) bool {
	return false
}
