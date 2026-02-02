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
	fmt.Printf("IssuerRegistry deployed at: %s\n", issuerAddr.Hex())

	fmt.Println("Deploying VerificationRegistry...")
	verAddr, tx, verContract, err := contracts.DeployVerificationRegistry(
		auth,
		c.ethClient,
		issuerAddr,
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
	fmt.Printf("VerificationRegistry deployed at: %s\n", verAddr.Hex())

	fmt.Println("Deploying VerifiedDocumentNFT...")
	nftAddr, tx, nftContract, err := contracts.DeployVerifiedDocumentNFT(
		auth,
		c.ethClient,
		"Verified Documents",
		"VDOC",
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
	fmt.Printf("VerifiedDocumentNFT deployed at: %s\n", nftAddr.Hex())

	fmt.Println("\nSetting up contract relationships...")

	tx, err = c.nft.SetVerificationRegistry(auth, verAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to set verification registry in NFT: %w", err)
	}
	_, err = bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, err
	}
	fmt.Println("NFT linked to VerificationRegistry")

	tx, err = c.verificationRegistry.SetNFTContract(auth, nftAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to set NFT contract in registry: %w", err)
	}
	_, err = bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, err
	}
	fmt.Println("VerificationRegistry linked to NFT")

	fmt.Println("\nAll contracts deployed and configured!")

	c.issuerRegistryAddr = addresses.IssuerRegistry
	c.verificationRegistryAddr = addresses.VerificationRegistry
	c.nftContractAddr = addresses.NFT

	return addresses, nil
}
