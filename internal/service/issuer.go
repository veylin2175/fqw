package service

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"fqw/internal/blockchain"
)

type IssuerService struct {
	client *blockchain.Client
}

func NewIssuerService(client *blockchain.Client) *IssuerService {
	return &IssuerService{client: client}
}

// AddTrustedIssuer добавляет доверенную организацию
func (s *IssuerService) AddTrustedIssuer(ctx context.Context, issuerAddress common.Address) error {
	auth, err := s.client.GetAdminTransactor()
	if err != nil {
		return fmt.Errorf("failed to get admin transactor: %w", err)
	}

	tx, err := s.client.IssuerRegistry().AddIssuer(auth, issuerAddress)
	if err != nil {
		return fmt.Errorf("failed to add issuer: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, s.client.EthClient(), tx)
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction reverted")
	}

	fmt.Printf("✓ Issuer %s added successfully (tx: %s)\n",
		issuerAddress.Hex(), tx.Hash().Hex())

	return nil
}

// RemoveTrustedIssuer удаляет доверенную организацию
func (s *IssuerService) RemoveTrustedIssuer(ctx context.Context, issuerAddress common.Address) error {
	auth, err := s.client.GetAdminTransactor()
	if err != nil {
		return fmt.Errorf("failed to get admin transactor: %w", err)
	}

	tx, err := s.client.IssuerRegistry().RemoveIssuer(auth, issuerAddress)
	if err != nil {
		return fmt.Errorf("failed to remove issuer: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, s.client.EthClient(), tx)
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction reverted")
	}

	fmt.Printf("✓ Issuer %s removed successfully (tx: %s)\n",
		issuerAddress.Hex(), tx.Hash().Hex())

	return nil
}

// IsTrustedIssuer проверяет, является ли адрес доверенным
func (s *IssuerService) IsTrustedIssuer(ctx context.Context, issuerAddress common.Address) (bool, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	isTrusted, err := s.client.IssuerRegistry().IsTrustedIssuer(callOpts, issuerAddress)
	if err != nil {
		return false, fmt.Errorf("failed to check issuer: %w", err)
	}

	return isTrusted, nil
}

// GetOwner возвращает адрес владельца контракта
func (s *IssuerService) GetOwner(ctx context.Context) (common.Address, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	owner, err := s.client.IssuerRegistry().Owner(callOpts)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get owner: %w", err)
	}

	return owner, nil
}

// TransferOwnership передает права владельца
func (s *IssuerService) TransferOwnership(ctx context.Context, newOwner common.Address) error {
	auth, err := s.client.GetAdminTransactor()
	if err != nil {
		return fmt.Errorf("failed to get admin transactor: %w", err)
	}

	tx, err := s.client.IssuerRegistry().TransferOwnership(auth, newOwner)
	if err != nil {
		return fmt.Errorf("failed to transfer ownership: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, s.client.EthClient(), tx)
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction reverted")
	}

	fmt.Printf("✓ Ownership transferred to %s (tx: %s)\n",
		newOwner.Hex(), tx.Hash().Hex())

	return nil
}
