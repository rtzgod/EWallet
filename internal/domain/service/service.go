package service

import "github.com/rtzgod/EWallet/internal/repository"

type Wallet interface {
}

type Transaction interface {
}

type Service struct {
	Wallet
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Wallet: NewWalletService(repos.Wallet),
	}
}
