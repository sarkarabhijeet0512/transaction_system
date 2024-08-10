package transaction

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Service struct {
	conf *viper.Viper
	log  *logrus.Logger
	Repo Repository
}

// NewService returns a user service object.
func NewService(conf *viper.Viper, log *logrus.Logger, Repo Repository) *Service {
	return &Service{conf: conf, log: log, Repo: Repo}
}

func (s *Service) CreateTransaction(ctx context.Context, tx *Transaction) error {
	return s.Repo.createTransaction(ctx, tx)
}

func (s *Service) GetTransactionByID(ctx context.Context, id int64) (*Transaction, error) {
	return s.Repo.getTransactionByID(ctx, id)
}

func (s *Service) GetTransactionsByType(ctx context.Context, txType string) ([]int64, error) {
	return s.Repo.getTransactionsByType(ctx, txType)
}

func (s *Service) GetSumByTransactionID(ctx context.Context, id int64) (float64, error) {
	return s.Repo.getSumByTransactionID(ctx, id)
}
