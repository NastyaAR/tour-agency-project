package domain

import (
	"context"
	"github.com/sirupsen/logrus"
)

const (
	AuthorizeSuccess = iota
	NotAuthorizedError
	ParseTokenError
	ExtractUserIdError = -1
)

var TokenString = ""

type IAuthService interface {
	GetToken(login string, lg *logrus.Logger) (string, error)
	AddToken(login string, token string, lg *logrus.Logger) error
	CheckAccessRights(token string, needRoleLevel string, lg *logrus.Logger) (int, error)
	ExtractIdFromToken(tokenString string, lg *logrus.Logger) (int, error)
}

type IAuthRepo interface {
	GetToken(c context.Context, login string, lg *logrus.Logger) (string, error)
	AddToken(c context.Context, login string, token string, lg *logrus.Logger) error
}
