package transaction

import (
	"context"
	"transaction_system/utils"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Repository interface {
	createTransaction(ctx context.Context, tx *Transaction) error
	getTransactionByID(ctx context.Context, id int64) (*Transaction, error)
	getTransactionsByType(ctx context.Context, txType string) ([]int64, error)
	getSumByTransactionID(ctx context.Context, id int64) (*Sum, error)
}

// NewRepositoryIn is function param struct of func `NewRepository`
type NewRepositoryIn struct {
	fx.In

	Log *logrus.Logger
	DB  *pg.DB `name:"transactiondb"`
}

// PGRepo is postgres implementation
type dbRepo struct {
	log *logrus.Logger
	db  *pg.DB
}

// NewDBRepository returns a new persistence layer object which can be used for
// CRUD on db
func NewDBRepository(i NewRepositoryIn) (Repo Repository, err error) {
	Repo = &dbRepo{
		log: i.Log,
		db:  i.DB,
	}

	return
}
func (r *dbRepo) createTransaction(ctx context.Context, tx *Transaction) error {
	utils.SetGenericFieldValue(tx)
	_, err := r.db.ModelContext(ctx, tx).
		Returning("*").OnConflict("(transaction_id) DO UPDATE").
		Insert()
	//INSERT INTO transactions (id, amount, type, parent_id) VALUES (id, amount, type, parent_id)`, &tx)
	return err
}

func (r *dbRepo) getTransactionByID(ctx context.Context, id int64) (*Transaction, error) {
	var tx Transaction
	err := r.db.ModelContext(ctx, &tx).Where("transaction_id = ?", id).Select()
	// SELECT * FROM transactions WHERE id = id
	if err != nil && err == pg.ErrNoRows {
		err = nil
	}
	return &tx, err
}

func (r *dbRepo) getTransactionsByType(ctx context.Context, txType string) ([]int64, error) {
	var ids []int64
	err := r.db.ModelContext(ctx, &Transaction{}).Column("transaction_id").Where("type = ?", txType).Select(&ids)
	//SELECT id FROM transactions WHERE type = txtype

	return ids, err
}

func (r *dbRepo) getSumByTransactionID(ctx context.Context, id int64) (sum *Sum, err error) {
	sum = &Sum{}

	_, err = r.db.QueryContext(ctx, sum, `
		WITH RECURSIVE tx_tree AS (
			SELECT transaction_id, amount FROM transactions WHERE transaction_id = ?
			UNION ALL
			SELECT t.transaction_id, t.amount FROM transactions t
			JOIN tx_tree tt ON t.parent_id = tt.transaction_id
		)
		SELECT SUM(amount) FROM tx_tree
	`, id)
	return sum, err
}
