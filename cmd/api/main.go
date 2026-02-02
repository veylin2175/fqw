package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"fqw/internal/api"
	"fqw/internal/blockchain"
)

var (
	port   = flag.String("port", "8080", "HTTP server port")
	rpcURL = flag.String("rpc", "http://127.0.0.1:8545", "Blockchain RPC URL")

	adminKey  = flag.String("admin-key", "", "Admin private key")
	issuerKey = flag.String("issuer-key", "", "Issuer private key")

	issuerRegistryAddr       = flag.String("issuer-registry", "", "IssuerRegistry address")
	verificationRegistryAddr = flag.String("verification-registry", "", "VerificationRegistry address")
	nftContractAddr          = flag.String("nft-contract", "", "NFT contract address")

	apiKey = flag.String("api-key", "dev-secret-key", "API key for admin endpoints")
)

func main() {
	flag.Parse()

	if *adminKey == "" {
		log.Fatal("--admin-key is required")
	}

	if *issuerRegistryAddr == "" || *verificationRegistryAddr == "" || *nftContractAddr == "" {
		log.Fatal("Contract addresses are required: --issuer-registry, --verification-registry, --nft-contract")
	}

	if *issuerKey == "" {
		*issuerKey = *adminKey
	}

	bcConfig := &blockchain.Config{
		RPCURL:                   *rpcURL,
		ChainID:                  31337,
		IssuerRegistryAddr:       common.HexToAddress(*issuerRegistryAddr),
		VerificationRegistryAddr: common.HexToAddress(*verificationRegistryAddr),
		NFTContractAddr:          common.HexToAddress(*nftContractAddr),
		AdminPrivateKey:          *adminKey,
		IssuerPrivateKey:         *issuerKey,
	}

	client, err := blockchain.NewClient(bcConfig)
	if err != nil {
		log.Fatalf("Failed to connect to blockchain: %v", err)
	}
	defer client.Close()

	log.Println("Connected to blockchain")
	log.Printf("   IssuerRegistry: %s\n", bcConfig.IssuerRegistryAddr.Hex())
	log.Printf("   VerificationRegistry: %s\n", bcConfig.VerificationRegistryAddr.Hex())
	log.Printf("   NFT Contract: %s\n", bcConfig.NFTContractAddr.Hex())

	routerConfig := &api.RouterConfig{
		Client: client,
		APIKey: *apiKey,
	}

	router := api.SetupRouter(routerConfig)

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: router,
	}

	go func() {
		log.Printf("API Server starting on port %s\n", *port)
		log.Println("API Documentation:")
		log.Println("   POST   /api/v1/verify              - Verify document")
		log.Println("   GET    /api/v1/verify/:tokenId     - Check verification")
		log.Println("   DELETE /api/v1/verify/:tokenId     - Revoke verification")
		log.Println("   GET    /api/v1/system/info         - System info")
		log.Println("   POST   /api/v1/admin/issuer        - Add issuer (requires API key)")
		log.Println()

		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server stopped gracefully")
}
