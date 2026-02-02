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

	issuerRegistry       *contracts.IssuerRegistry
	verificationRegistry *contracts.VerificationRegistry
	nft                  *contracts.VerifiedDocumentNFT

	issuerRegistryAddr       common.Address
	verificationRegistryAddr common.Address
	nftContractAddr          common.Address

	adminKey  *ecdsa.PrivateKey
	issuerKey *ecdsa.PrivateKey
}

func NewClient(cfg *Config) (*Client, error) {
	ethClient, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum node: %w", err)
	}

	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

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

	if cfg.IssuerRegistryAddr != (common.Address{}) {
		if err = client.ConnectToContracts(cfg); err != nil {
			return nil, fmt.Errorf("failed to connect to contracts: %w", err)
		}
	}

	return client, nil
}

func (c *Client) ConnectToContracts(cfg *Config) error {
	var err error

	c.issuerRegistryAddr = cfg.IssuerRegistryAddr
	c.verificationRegistryAddr = cfg.VerificationRegistryAddr
	c.nftContractAddr = cfg.NFTContractAddr

	c.issuerRegistry, err = contracts.NewIssuerRegistry(
		cfg.IssuerRegistryAddr,
		c.ethClient,
	)
	if err != nil {
		return fmt.Errorf("failed to load IssuerRegistry: %w", err)
	}

	c.verificationRegistry, err = contracts.NewVerificationRegistry(
		cfg.VerificationRegistryAddr,
		c.ethClient,
	)
	if err != nil {
		return fmt.Errorf("failed to load VerificationRegistry: %w", err)
	}

	c.nft, err = contracts.NewVerifiedDocumentNFT(
		cfg.NFTContractAddr,
		c.ethClient,
	)
	if err != nil {
		return fmt.Errorf("failed to load VerifiedDocumentNFT: %w", err)
	}

	return nil
}

func (c *Client) getTransactor(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

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

func (c *Client) GetAdminTransactor() (*bind.TransactOpts, error) {
	return c.getTransactor(c.adminKey)
}

func (c *Client) GetIssuerTransactor() (*bind.TransactOpts, error) {
	return c.getTransactor(c.issuerKey)
}

func (c *Client) EthClient() *ethclient.Client {
	return c.ethClient
}

func (c *Client) IssuerRegistryAddr() common.Address {
	return c.issuerRegistryAddr
}

func (c *Client) VerificationRegistryAddr() common.Address {
	return c.verificationRegistryAddr
}

func (c *Client) NFTContractAddr() common.Address {
	return c.nftContractAddr
}
