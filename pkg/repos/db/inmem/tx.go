package inmem

import (
	"context"
	"slices"
	"splitmoney/pkg/domain/account"
	"splitmoney/pkg/domain/payment"
	"time"
)

type txInDB struct {
	FromAccount account.AccountID
	ToAccount   account.AccountID
	Amount      int
	CreatedAt   time.Time
}

type Transactions []*txInDB

type txStore struct {
	fromTransactions map[account.AccountID]Transactions
	toTransactions   map[account.AccountID]Transactions
	txs              map[account.AccountID]Transactions
}

func NewTxStore() payment.Storer {
	return &txStore{
		fromTransactions: make(map[account.AccountID]Transactions),
		toTransactions:   make(map[account.AccountID]Transactions),
		txs:              make(map[account.AccountID]Transactions),
	}
}

func (s *txStore) CreateTx(_ context.Context, tx payment.Transaction) error {
	td := txInDB{
		FromAccount: tx.FromAccount,
		ToAccount:   tx.ToAccount,
		Amount:      tx.Amount,
		CreatedAt:   time.Now(),
	}
	s.fromTransactions[tx.FromAccount] = append(s.fromTransactions[tx.FromAccount], &td)
	s.toTransactions[tx.ToAccount] = append(s.toTransactions[tx.ToAccount], &td)

	s.txs[tx.FromAccount] = append(s.txs[tx.FromAccount], &td)
	s.txs[tx.ToAccount] = append(s.txs[tx.ToAccount], &td)
	return nil
}

func (s *txStore) GetTxs(_ context.Context, accID account.AccountID) ([]payment.Transaction, error) {

	if _, ok := s.txs[accID]; !ok {
		return nil, nil
	}

	txs := make([]payment.Transaction, len(s.txs[accID]))

	for i, td := range s.txs[accID] {
		txs[i] = payment.Transaction{
			FromAccount: td.FromAccount,
			ToAccount:   td.ToAccount,
			Amount:      td.Amount,
		}
	}
	return txs, nil
}

func (s *txStore) GetTxsInDuration(
	_ context.Context,
	accID account.AccountID,
	from time.Time,
	to time.Time,
) (
	[]payment.Transaction,
	error,
) {
	if _, ok := s.txs[accID]; !ok {
		return nil, nil
	}

	accTxs := s.txs[accID]
	bSearchCmp := func(tid *txInDB, t time.Time) int {
		if tid.CreatedAt == t {
			return 0
		}
		if tid.CreatedAt.Before(t) {
			return -1
		}
		return 1
	}

	startIndex, _ := slices.BinarySearchFunc[Transactions, *txInDB, time.Time](
		accTxs, from, bSearchCmp,
	)
	endIndex, ok := slices.BinarySearchFunc[Transactions, *txInDB, time.Time](
		accTxs, to, bSearchCmp,
	)

	if !ok {
		endIndex--
	}

	len := endIndex - startIndex + 1
	txs := make([]payment.Transaction, len)

	txsInDB := s.txs[accID][startIndex : endIndex+1]
	for i, td := range txsInDB {
		txs[i] = payment.Transaction{
			FromAccount: td.FromAccount,
			ToAccount:   td.ToAccount,
			Amount:      td.Amount,
		}
	}
	return txs, nil
}
