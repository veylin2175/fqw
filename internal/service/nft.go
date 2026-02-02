package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"fqw/internal/blockchain"
)

type NFTService struct {
	client *blockchain.Client
}

func NewNFTService(client *blockchain.Client) *NFTService {
	return &NFTService{client: client}
}

func (s *NFTService) MintNFT(
	ctx context.Context,
	to common.Address,
	tokenId *big.Int,
) error {
	auth, err := s.client.GetAdminTransactor()
	if err != nil {
		return fmt.Errorf("failed to get admin transactor: %w", err)
	}

	tx, err := s.client.NFT().Mint(auth, to, tokenId)
	if err != nil {
		return fmt.Errorf("failed to mint NFT: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, s.client.EthClient(), tx)
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction reverted")
	}

	fmt.Printf("NFT minted (tokenId: %s, owner: %s, tx: %s)\n",
		tokenId.String(), to.Hex(), tx.Hash().Hex())

	return nil
}

func (s *NFTService) GetOwnerOf(ctx context.Context, tokenId *big.Int) (common.Address, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	owner, err := s.client.NFT().OwnerOf(callOpts, tokenId)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get owner: %w", err)
	}

	return owner, nil
}

func (s *NFTService) GetTokenURI(ctx context.Context, tokenId *big.Int) (string, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	uri, err := s.client.NFT().TokenURI(callOpts, tokenId)
	if err != nil {
		return "", fmt.Errorf("failed to get token URI: %w", err)
	}

	return uri, nil
}

func (s *NFTService) SetBaseURI(ctx context.Context, baseURI string) error {
	auth, err := s.client.GetAdminTransactor()
	if err != nil {
		return fmt.Errorf("failed to get admin transactor: %w", err)
	}

	tx, err := s.client.NFT().SetBaseURI(auth, baseURI)
	if err != nil {
		return fmt.Errorf("failed to set base URI: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, s.client.EthClient(), tx)
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction reverted")
	}

	fmt.Printf("Base URI set to: %s (tx: %s)\n", baseURI, tx.Hash().Hex())

	return nil
}

func (s *NFTService) GetBalanceOf(ctx context.Context, owner common.Address) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	balance, err := s.client.NFT().BalanceOf(callOpts, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance, nil
}

func (s *NFTService) GetName(ctx context.Context) (string, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	name, err := s.client.NFT().Name(callOpts)
	if err != nil {
		return "", fmt.Errorf("failed to get name: %w", err)
	}

	return name, nil
}

func (s *NFTService) GetSymbol(ctx context.Context) (string, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	symbol, err := s.client.NFT().Symbol(callOpts)
	if err != nil {
		return "", fmt.Errorf("failed to get symbol: %w", err)
	}

	return symbol, nil
}
