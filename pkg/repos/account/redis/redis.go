package redis

import (
	"context"
	"encoding/json"
	"log/slog"
	"splitmoney/pkg/domain/account"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Id   int
	Addr string
}

type redisClint struct {
	client *redis.Client
}

func New(config RedisConfig) account.Storer {
	cli := redis.NewClient(&redis.Options{
		Addr: config.Addr,
		DB:   config.Id,
	})
	return &redisClint{
		client: cli,
	}
}

func (r *redisClint) CreateAccount(ctx context.Context, acc account.Account) error {
	accInDb := &accountModel{
		ID:        acc.ID,
		Debt:      acc.Debt,
		Lend:      acc.Lend,
		Balance:   acc.Balance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data, err := json.Marshal(accInDb)
	if err != nil {
		slog.Error(err.Error())
		return account.ErrorInvalidAccountData
	}
	r.client.Set(strconv.Itoa(int(acc.ID)), data, 0)
	return nil
}

func (r *redisClint) GetAccount(context.Context, account.AccountID) (account.Account, error) {
	return account.Account{}, nil
}
