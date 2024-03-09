package redis

import (
	"context"
	"encoding/json"
	"log/slog"
	"splitmoney/pkg/domain/account"
	"splitmoney/pkg/domain/payment"
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

func (r *redisClint) CreateAccount(_ context.Context, acc account.Account) error {
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
	r.client.Set(strconv.Itoa(int(acc.ID)), string(data), 0)
	return nil
}

func (r *redisClint) GetAccount(_ context.Context, accID account.AccountID) (account.Account, error) {
	var acc account.Account
	cmd := r.client.Get(strconv.Itoa(int(accID)))
	if cmd.Err() != nil {
		slog.Error(cmd.Err().Error())
		return acc, account.ErrorAccountNotFound
	}
	data := cmd.Val()
	var accInDb accountModel
	err := json.Unmarshal([]byte(data), &accInDb)
	if err != nil {
		slog.Error(err.Error())
		// TODO how to report downstream errors
		return acc, err
	}

	acc.ID = accInDb.ID
	acc.Debt = accInDb.Debt
	acc.Lend = accInDb.Lend
	acc.Balance = accInDb.Balance

	return acc, nil
}

func (r *redisClint) CreateTx(_ context.Context, txn payment.Transaction) error {
	tym := time.Now()
	day := tym.Day()
	month := int(tym.Month())
	year := tym.Year()
	hour := tym.Hour()
	min := tym.Minute()
	sec := tym.Second()

	txnId := ((year * 1000000000) + (month * 100000000) + (day * 1000000) + (hour * 10000) + (min * 100) + sec)
	paymentInDb := &paymentModel{
		FromAccount:   txn.FromAccount,
		ToAccount:     txn.ToAccount,
		Ammount:       txn.Amount,
		TransactionId: txnId,
		CreatedAt:     tym,
	}
	data, error := json.Marshal(paymentInDb)

	if error != nil {
		slog.Error(error.Error())
		return payment.ErrorTransactionFailed
	}

	r.client.Set(strconv.Itoa(txnId), string(data), 0)

	return nil
}

func (r *redisClint) GetTxs(_ context.Context, accId account.AccountID) ([]payment.Transaction, error) {
	var txns []payment.Transaction

	return txns, nil
}
