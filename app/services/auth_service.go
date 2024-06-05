package services

import (
	"app/domain"
	"app/pkg"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

type authService struct {
	authRepo domain.IAuthRepo
	timeout  time.Duration
}

func CreateNewAuthService(authRepo domain.IAuthRepo, t time.Duration) domain.IAuthService {
	return &authService{
		authRepo: authRepo,
		timeout:  t,
	}
}

func (a *authService) GetToken(login string, lg *logrus.Logger) (string, error) {
	if login == "" {
		lg.Warnf("bad empty login")
		return "", xerrors.Errorf("auth service: gettoken error: empty login")
	}
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	token, err := a.authRepo.GetToken(ctx, login, lg)
	if err != nil {
		lg.Warnf("auth service: gettoken error: %v", err.Error())
		return "", xerrors.Errorf("auth service: gettoken error: %v", err.Error())
	}

	return token, nil
}

func (a *authService) AddToken(login string, token string, lg *logrus.Logger) error {
	if login == "" {
		lg.Warnf("bad empty login")
		return xerrors.Errorf("auth service: addtoken error: empty login")
	}
	if token == "" {
		lg.Warnf("bad empty token")
		return xerrors.Errorf("auth service: addtoken error: empty token")
	}
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	err := a.authRepo.AddToken(ctx, login, token, lg)
	if err != nil {
		lg.Warnf("auth service: addtoken error: %v", err.Error())
		return xerrors.Errorf("auth service: addtoken error: %v", err.Error())
	}

	return nil
}

func (a *authService) CheckAccessRights(tokenString string, needRoleLevel string, lg *logrus.Logger) (int, error) {
	if tokenString == "" && needRoleLevel != domain.GuestUser {
		lg.Warnf("insufficient level of rights")
		return domain.NotAuthorizedError, xerrors.Errorf("auth service: check access rights error: %v")
	}
	if tokenString == "" && needRoleLevel == domain.GuestUser {
		return domain.AuthorizeSuccess, nil
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(pkg.Key), nil
	})

	if err != nil {
		lg.Warnf("parse token string error")
		return domain.ParseTokenError, xerrors.Errorf("auth service: check access rights error: %v", err.Error())
	}
	if !token.Valid {
		lg.Warnf("parse token string error")
		return domain.ParseTokenError, xerrors.Errorf("auth service: check access rights error: %v", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		lg.Warnf("bad extract claims of token")
		return domain.ParseTokenError, xerrors.Errorf("auth service: check access rights error: %v", err.Error())
	}
	role, ok := claims["role"].(string)
	if !ok {
		lg.Warnf("bad extract role from claims")
		return domain.ParseTokenError, xerrors.Errorf("auth service: check access rights error: %v", err.Error())
	}

	switch needRoleLevel {
	case domain.AdminUser:
		if role != domain.AdminUser {
			lg.Warnf("insufficient level of rights")
			return domain.NotAuthorizedError, xerrors.Errorf("auth service: check access rights error")
		}
	case domain.ManagerUser:
		if role != domain.AdminUser && role != domain.ManagerUser {
			lg.Warnf("insufficient level of rights")
			return domain.NotAuthorizedError, xerrors.Errorf("auth service: check access rights error")
		}
	case domain.ClientUser:
		if role == domain.GuestUser {
			lg.Warnf("insufficient level of rights")
			return domain.NotAuthorizedError, xerrors.Errorf("auth service: check access rights error")
		}
	}

	return domain.AuthorizeSuccess, nil
}

func (a *authService) ExtractIdFromToken(tokenString string, lg *logrus.Logger) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(pkg.Key), nil
	})

	if err != nil {
		lg.Warnf("parse token string error")
		return domain.ExtractUserIdError, xerrors.Errorf("auth service: extract id error: %v", err.Error())
	}
	if !token.Valid {
		lg.Warnf("parse token string error")
		return domain.ExtractUserIdError, xerrors.Errorf("auth service: extract id error: %v", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		lg.Warnf("bad extract claims of token")
		return domain.ExtractUserIdError, xerrors.Errorf("auth service: extract id error: %v", err.Error())
	}
	id := claims["userID"].(float64)
	intId := int(id)
	if !ok {
		lg.Warnf("bad extract role from claims")
		return domain.ExtractUserIdError, xerrors.Errorf("auth service: extract id error")
	}
	return intId, nil
}
