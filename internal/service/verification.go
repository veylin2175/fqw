package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"fqw/internal/blockchain"
)

type VerificationService struct {
	client *blockchain.Client
}

type VerificationResult struct {
	Valid    bool
	Issuer   common.Address
	Revoked  bool
	IssuedAt *big.Int
}

func NewVerificationService(client *blockchain.Client) *VerificationService {
	return &VerificationService{client: client}
}

func (s *VerificationService) RegisterVerification(
	ctx context.Context,
	vcHash [32]byte,
	subjectAddress common.Address,
) (*big.Int, error) {

	auth, err := s.client.GetIssuerTransactor()
	if err != nil {
		return nil, fmt.Errorf("failed to get issuer transactor: %w", err)
	}

	// Вызов registerVerification
	tx, err := s.client.VerificationRegistry().RegisterVerification(
		auth,
		vcHash,
		subjectAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register verification: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, s.client.EthClient(), tx)
	if err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}

	if receipt.Status != 1 {
		return nil, fmt.Errorf("transaction reverted")
	}

	for _, log := range receipt.Logs {
		event, err := s.client.VerificationRegistry().ParseVerificationRegistered(*log)
		if err == nil {
			fmt.Printf("Verification registered (tokenId: %s, tx: %s)\n",
				event.TokenId.String(), tx.Hash().Hex())
			return event.TokenId, nil
		}
	}

	return nil, fmt.Errorf("failed to parse VerificationRegistered event")
}

func (s *VerificationService) RevokeVerification(
	ctx context.Context,
	tokenId *big.Int,
) error {

	auth, err := s.client.GetIssuerTransactor()
	if err != nil {
		return fmt.Errorf("failed to get issuer transactor: %w", err)
	}

	tx, err := s.client.VerificationRegistry().RevokeVerification(auth, tokenId)
	if err != nil {
		return fmt.Errorf("failed to revoke verification: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, s.client.EthClient(), tx)
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction reverted")
	}

	fmt.Printf("Verification revoked (tokenId: %s, tx: %s)\n",
		tokenId.String(), tx.Hash().Hex())

	return nil
}

func (s *VerificationService) VerifyDocument(
	ctx context.Context,
	tokenId *big.Int,
	vcHash [32]byte,
) (*VerificationResult, error) {

	callOpts := &bind.CallOpts{Context: ctx}

	result, err := s.client.VerificationRegistry().Verify(
		callOpts,
		tokenId,
		vcHash,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to verify: %w", err)
	}

	return &VerificationResult{
		Valid:    result.Valid,
		Issuer:   result.Issuer,
		Revoked:  result.Revoked,
		IssuedAt: result.IssuedAt,
	}, nil
}

func (s *VerificationService) GetTokenCounter(ctx context.Context) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	counter, err := s.client.VerificationRegistry().TokenCounter(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get token counter: %w", err)
	}

	return counter, nil
}

func (s *VerificationService) GetVerificationDetails(
	ctx context.Context,
	tokenId *big.Int,
) (*VerificationResult, error) {

	return s.VerifyDocument(ctx, tokenId, [32]byte{})
}
