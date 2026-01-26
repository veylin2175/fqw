package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"fqw/contracts"
)

type Client struct {
	ethClient *ethclient.Client
	chainID   *big.Int

	// Контракты
	issuerRegistry       *contracts.IssuerRegistry
	verificationRegistry *contracts.VerificationRegistry
	nft                  *contracts.VerifiedDocumentNFT

	// Приватные ключи
	adminKey  *ecdsa.PrivateKey
	issuerKey *ecdsa.PrivateKey
}

func NewClient(cfg *Config) (*Client, error) {
	// Подключение к ноде
	ethClient, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum node: %w", err)
	}

	// Проверка подключения
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Загрузка приватных ключей
	adminKey, err := crypto.HexToECDSA(cfg.AdminPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid admin private key: %w", err)
	}

	issuerKey, err := crypto.HexToECDSA(cfg.IssuerPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid issuer private key: %w", err)
	}

	client := &Client{
		ethClient: ethClient,
		chainID:   chainID,
		adminKey:  adminKey,
		issuerKey: issuerKey,
	}

	// Подключение к контрактам (если адреса указаны)
	if cfg.IssuerRegistryAddr != (common.Address{}) {
		if err = client.ConnectToContracts(cfg); err != nil {
			return nil, fmt.Errorf("failed to connect to contracts: %w", err)
		}
	}

	return client, nil
}

func (c *Client) ConnectToContracts(cfg *Config) error {
	var err error

	// IssuerRegistry
	c.issuerRegistry, err = contracts.NewIssuerRegistry(
		cfg.IssuerRegistryAddr,
		c.ethClient,
	)
	if err != nil {
		return fmt.Errorf("failed to load IssuerRegistry: %w", err)
	}

	// VerificationRegistry
	c.verificationRegistry, err = contracts.NewVerificationRegistry(
		cfg.VerificationRegistryAddr,
		c.ethClient,
	)
	if err != nil {
		return fmt.Errorf("failed to load VerificationRegistry: %w", err)
	}

	// NFT
	c.nft, err = contracts.NewVerifiedDocumentNFT(
		cfg.NFTContractAddr,
		c.ethClient,
	)
	if err != nil {
		return fmt.Errorf("failed to load VerifiedDocumentNFT: %w", err)
	}

	return nil
}

// Создание transactor для отправки транзакций
func (c *Client) getTransactor(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)
	if err != nil {
		return nil, err
	}

	// Можно настроить gas limit, gas price и т.д.
	// auth.GasLimit = 3000000

	return auth, nil
}

// Геттеры для контрактов
func (c *Client) IssuerRegistry() *contracts.IssuerRegistry {
	return c.issuerRegistry
}

func (c *Client) VerificationRegistry() *contracts.VerificationRegistry {
	return c.verificationRegistry
}

func (c *Client) NFT() *contracts.VerifiedDocumentNFT {
	return c.nft
}

func (c *Client) Close() {
	if c.ethClient != nil {
		c.ethClient.Close()
	}
}

// GetAdminTransactor возвращает transactor для администратора
func (c *Client) GetAdminTransactor() (*bind.TransactOpts, error) {
	return c.getTransactor(c.adminKey)
}

// GetIssuerTransactor возвращает transactor для issuer'а
func (c *Client) GetIssuerTransactor() (*bind.TransactOpts, error) {
	return c.getTransactor(c.issuerKey)
}

// EthClient возвращает ethclient для прямого доступа
func (c *Client) EthClient() *ethclient.Client {
	return c.ethClient
}
