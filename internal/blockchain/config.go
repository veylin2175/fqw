package blockchain

import "github.com/ethereum/go-ethereum/common"

type Config struct {
	RPCURL  string
	ChainID int64

	// Адреса развернутых контрактов
	IssuerRegistryAddr       common.Address
	VerificationRegistryAddr common.Address
	NFTContractAddr          common.Address

	// Приватные ключи (в проде использовать env или KMS)
	AdminPrivateKey  string
	IssuerPrivateKey string
}

// Для тестовой сети (например, Sepolia)
func DefaultTestConfig() *Config {
	return &Config{
		RPCURL:  "https://sepolia.infura.io/v3/YOUR_INFURA_KEY",
		ChainID: 11155111, // Sepolia
		// Адреса заполнятся после деплоя
	}
}

// Для локальной ноды (Hardhat/Ganache)
func LocalConfig() *Config {
	return &Config{
		RPCURL:  "http://127.0.0.1:8545",
		ChainID: 31337, // Hardhat default
	}
}
