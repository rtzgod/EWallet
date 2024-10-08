package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rtzgod/ewallet-api/internal/domain/entity"
	"github.com/rtzgod/ewallet-api/internal/repository/postgres"
)

type Wallet interface {
	Create(id string) (entity.Wallet, error)
	GetById(id string) (entity.Wallet, error)
	Update(senderId, receiverId string, amount float64) error
}

type Transaction interface {
	Create(senderId, receiverId string, amount float64) error
	GetAllById(id string) ([]entity.Transaction, error)
}

type Repository struct {
	Wallet
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Wallet:      postgres.NewWalletPostgres(db),
		Transaction: postgres.NewTransactionPostgres(db),
	}
}
