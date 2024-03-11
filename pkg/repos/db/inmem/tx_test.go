package inmem

import (
	"context"
	"splitmoney/pkg/domain/account"
	"splitmoney/pkg/domain/payment"
	"testing"
)

func TestCreateTx(t *testing.T) {
	store := &txStore{
		fromTransactions: make(map[account.AccountID]Transactions),
		toTransactions:   make(map[account.AccountID]Transactions),
		txs:              make(map[account.AccountID]Transactions),
	}

	tx1 := payment.Transaction{
		FromAccount: 10,
		ToAccount:   20,
		Amount:      100,
	}
	tx2 := payment.Transaction{
		FromAccount: 20,
		ToAccount:   10,
		Amount:      200,
	}

	_ = store.CreateTx(context.TODO(), tx1)
	_ = store.CreateTx(context.TODO(), tx2)

	if store.fromTransactions == nil || len(store.fromTransactions) != 2 {
		t.Error("failed")
	}

	if len(store.fromTransactions[10]) != 1 {
		t.Errorf("failed: %v", len(store.fromTransactions[10]))
	}

	if store.toTransactions == nil || len(store.toTransactions) != 2 {
		t.Error("failed")
	}

}
