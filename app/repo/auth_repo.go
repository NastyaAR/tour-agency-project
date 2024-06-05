package repo

import (
	"app/domain"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"time"
)

type redisAuthRepo struct {
	db               *redis.Client
	expTimeInSeconds time.Duration
}

func NewRedisAuthRepo(clnt *redis.Client, eTime time.Duration) domain.IAuthRepo {
	return &redisAuthRepo{db: clnt, expTimeInSeconds: eTime}
}

func (r *redisAuthRepo) GetToken(c context.Context, login string, lg *logrus.Logger) (string, error) {
	rStr := r.db.Get(c, login)
	if rStr.Err() == redis.Nil {
		lg.Warnf("redis auth repo: gettoken error: %v", rStr.Err)
		return "", xerrors.Errorf("redis auth repo: gettoken error: %v", rStr.Err())
	}

	token := rStr.String()
	return token, nil
}

func (r *redisAuthRepo) AddToken(c context.Context, login string, token string, lg *logrus.Logger) error {
	err := r.db.Set(c, login, token, r.expTimeInSeconds)
	if err.Err() != nil {
		lg.Warnf("redis auth repo: addtoken error: %v", err.Err())
		return xerrors.Errorf("redis auth repo: addtoken error: %v", err.Err())
	}
	return nil
}
