package blockchain

import "github.com/ethereum/go-ethereum/common"

type Config struct {
	RPCURL  string
	ChainID int64

	IssuerRegistryAddr       common.Address
	VerificationRegistryAddr common.Address
	NFTContractAddr          common.Address

	AdminPrivateKey  string
	IssuerPrivateKey string
}

func DefaultTestConfig() *Config {
	return &Config{
		RPCURL:  "https://sepolia.infura.io/v3/YOUR_INFURA_KEY",
		ChainID: 11155111,
	}
}

func LocalConfig() *Config {
	return &Config{
		RPCURL:  "http://127.0.0.1:8545",
		ChainID: 31337,
	}
}
