package blockchain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"fqw/contracts"
)

type DeployedAddresses struct {
	IssuerRegistry       common.Address
	VerificationRegistry common.Address
	NFT                  common.Address
}

func (c *Client) DeployAll(ctx context.Context) (*DeployedAddresses, error) {
	auth, err := c.getTransactor(c.adminKey)
	if err != nil {
		return nil, err
	}

	addresses := &DeployedAddresses{}

	// 1. Деплой IssuerRegistry
	fmt.Println("Deploying IssuerRegistry...")
	issuerAddr, tx, issuerContract, err := contracts.DeployIssuerRegistry(
		auth,
		c.ethClient,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy IssuerRegistry: %w", err)
	}

	_, err = bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to mine IssuerRegistry: %w", err)
	}

	addresses.IssuerRegistry = issuerAddr
	c.issuerRegistry = issuerContract
	fmt.Printf("✓ IssuerRegistry deployed at: %s\n", issuerAddr.Hex())

	// 2. Деплой VerificationRegistry
	fmt.Println("Deploying VerificationRegistry...")
	verAddr, tx, verContract, err := contracts.DeployVerificationRegistry(
		auth,
		c.ethClient,
		issuerAddr, // передаем адрес IssuerRegistry
	)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy VerificationRegistry: %w", err)
	}

	_, err = bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to mine VerificationRegistry: %w", err)
	}

	addresses.VerificationRegistry = verAddr
	c.verificationRegistry = verContract
	fmt.Printf("✓ VerificationRegistry deployed at: %s\n", verAddr.Hex())

	// 3. Деплой NFT
	fmt.Println("Deploying VerifiedDocumentNFT...")
	nftAddr, tx, nftContract, err := contracts.DeployVerifiedDocumentNFT(
		auth,
		c.ethClient,
		"Verified Documents", // название коллекции
		"VDOC",               // символ
	)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy NFT: %w", err)
	}

	_, err = bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to mine NFT: %w", err)
	}

	addresses.NFT = nftAddr
	c.nft = nftContract
	fmt.Printf("✓ VerifiedDocumentNFT deployed at: %s\n", nftAddr.Hex())

	// 4. Настройка связей между контрактами
	fmt.Println("\nSetting up contract relationships...")

	// NFT -> VerificationRegistry
	tx, err = c.nft.SetVerificationRegistry(auth, verAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to set verification registry in NFT: %w", err)
	}
	_, err = bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, err
	}
	fmt.Println("✓ NFT linked to VerificationRegistry")

	// VerificationRegistry -> NFT
	tx, err = c.verificationRegistry.SetNFTContract(auth, nftAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to set NFT contract in registry: %w", err)
	}
	_, err = bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, err
	}
	fmt.Println("✓ VerificationRegistry linked to NFT")

	fmt.Println("\n🎉 All contracts deployed and configured!")

	return addresses, nil
}
