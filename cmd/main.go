package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"fqw/internal/blockchain"
	"fqw/internal/service"
	"github.com/ethereum/go-ethereum/common"
)

var (
	mode      = flag.String("mode", "help", "Режим работы: deploy, add-issuer, verify, check, revoke")
	rpcURL    = flag.String("rpc", "http://127.0.0.1:8545", "RPC URL ноды")
	adminKey  = flag.String("admin-key", "", "Приватный ключ администратора")
	issuerKey = flag.String("issuer-key", "", "Приватный ключ issuer'а")

	// Адреса развернутых контрактов
	issuerRegistryAddr       = flag.String("issuer-registry", "", "Адрес IssuerRegistry")
	verificationRegistryAddr = flag.String("verification-registry", "", "Адрес VerificationRegistry")
	nftContractAddr          = flag.String("nft-contract", "", "Адрес NFT контракта")

	// Параметры для операций
	issuerAddress = flag.String("issuer-addr", "", "Адрес issuer'а для добавления/проверки")
	vcHashStr     = flag.String("vc-hash", "", "Хэш VC (hex строка)")
	subjectAddr   = flag.String("subject", "", "Адрес subject'а (владельца документа)")
	tokenID       = flag.String("token-id", "", "ID токена для проверки/отзыва")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	switch *mode {
	case "deploy":
		runDeploy(ctx)
	case "add-issuer":
		runAddIssuer(ctx)
	case "verify":
		runVerify(ctx)
	case "check":
		runCheck(ctx)
	case "revoke":
		runRevoke(ctx)
	case "info":
		runInfo(ctx)
	default:
		printHelp()
	}
}

// Деплой всех контрактов
func runDeploy(ctx context.Context) {
	if *adminKey == "" {
		log.Fatal("Требуется --admin-key")
	}

	cfg := &blockchain.Config{
		RPCURL:           *rpcURL,
		ChainID:          31337, // Hardhat local
		AdminPrivateKey:  *adminKey,
		IssuerPrivateKey: *adminKey, // Временно используем тот же ключ
	}

	client, err := blockchain.NewClient(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer client.Close()

	fmt.Println("🚀 Начинаем деплой контрактов...")
	fmt.Println("RPC:", *rpcURL)
	fmt.Println()

	addresses, err := client.DeployAll(ctx)
	if err != nil {
		log.Fatalf("Ошибка деплоя: %v", err)
	}

	fmt.Println("\n📋 Сохраните эти адреса для дальнейшего использования:")
	fmt.Println("--issuer-registry", addresses.IssuerRegistry.Hex())
	fmt.Println("--verification-registry", addresses.VerificationRegistry.Hex())
	fmt.Println("--nft-contract", addresses.NFT.Hex())
}

// Добавление доверенного issuer'а
func runAddIssuer(ctx context.Context) {
	client := mustConnectToContracts()
	defer client.Close()

	if *issuerAddress == "" {
		log.Fatal("Требуется --issuer-addr")
	}

	issuerSvc := service.NewIssuerService(client)

	addr := common.HexToAddress(*issuerAddress)

	fmt.Printf("➕ Добавляем issuer'а: %s\n", addr.Hex())

	err := issuerSvc.AddTrustedIssuer(ctx, addr)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	// Проверка
	isTrusted, _ := issuerSvc.IsTrustedIssuer(ctx, addr)
	fmt.Printf("✅ Статус доверия: %v\n", isTrusted)
}

// Верификация документа и минт NFT
func runVerify(ctx context.Context) {
	client := mustConnectToContracts()
	defer client.Close()

	if *vcHashStr == "" || *subjectAddr == "" {
		log.Fatal("Требуется --vc-hash и --subject")
	}

	verSvc := service.NewVerificationService(client)

	// Преобразуем хэш в [32]byte
	vcHash := hashFromString(*vcHashStr)
	subject := common.HexToAddress(*subjectAddr)

	fmt.Printf("📝 Регистрируем верификацию...\n")
	fmt.Printf("   VC Hash: 0x%x\n", vcHash)
	fmt.Printf("   Subject: %s\n", subject.Hex())

	tokenId, err := verSvc.RegisterVerification(ctx, vcHash, subject)
	if err != nil {
		log.Fatalf("Ошибка верификации: %v", err)
	}

	fmt.Printf("\n🎫 TokenID: %s\n", tokenId.String())
	fmt.Println("💡 Используйте этот TokenID для проверки статуса")
}

// Проверка статуса документа
func runCheck(ctx context.Context) {
	client := mustConnectToContracts()
	defer client.Close()

	if *tokenID == "" {
		log.Fatal("Требуется --token-id")
	}

	verSvc := service.NewVerificationService(client)
	nftSvc := service.NewNFTService(client)

	tid := new(big.Int)
	tid.SetString(*tokenID, 10)

	fmt.Printf("🔍 Проверяем TokenID: %s\n\n", tid.String())

	// Получаем детали верификации
	details, err := verSvc.GetVerificationDetails(ctx, tid)
	if err != nil {
		log.Fatalf("Ошибка получения данных: %v", err)
	}

	fmt.Println("📄 Информация о верификации:")
	fmt.Printf("   Issuer: %s\n", details.Issuer.Hex())
	fmt.Printf("   Issued At: %s\n", time.Unix(details.IssuedAt.Int64(), 0))
	fmt.Printf("   Revoked: %v\n", details.Revoked)
	fmt.Printf("   Valid: %v\n", details.Valid)

	// Получаем владельца NFT
	owner, err := nftSvc.GetOwnerOf(ctx, tid)
	if err != nil {
		fmt.Printf("\n⚠️  NFT не найден (возможно не заминчен)\n")
	} else {
		fmt.Printf("\n🎨 NFT Owner: %s\n", owner.Hex())

		// Получаем URI
		uri, err := nftSvc.GetTokenURI(ctx, tid)
		if err == nil && uri != "" {
			fmt.Printf("   Token URI: %s\n", uri)
		}
	}
}

// Отзыв верификации
func runRevoke(ctx context.Context) {
	client := mustConnectToContracts()
	defer client.Close()

	if *tokenID == "" {
		log.Fatal("Требуется --token-id")
	}

	verSvc := service.NewVerificationService(client)

	tid := new(big.Int)
	tid.SetString(*tokenID, 10)

	fmt.Printf("❌ Отзываем верификацию TokenID: %s\n", tid.String())

	err := verSvc.RevokeVerification(ctx, tid)
	if err != nil {
		log.Fatalf("Ошибка отзыва: %v", err)
	}
}

// Информация о системе
func runInfo(ctx context.Context) {
	client := mustConnectToContracts()
	defer client.Close()

	issuerSvc := service.NewIssuerService(client)
	verSvc := service.NewVerificationService(client)
	nftSvc := service.NewNFTService(client)

	fmt.Println("ℹ️  Информация о системе\n")

	// Владелец IssuerRegistry
	owner, _ := issuerSvc.GetOwner(ctx)
	fmt.Printf("🔑 Owner IssuerRegistry: %s\n", owner.Hex())

	// Счетчик токенов
	counter, _ := verSvc.GetTokenCounter(ctx)
	fmt.Printf("📊 Всего верификаций: %s\n", counter.String())

	// Информация о NFT коллекции
	name, _ := nftSvc.GetName(ctx)
	symbol, _ := nftSvc.GetSymbol(ctx)
	fmt.Printf("🎨 NFT коллекция: %s (%s)\n", name, symbol)

	// Проверяем issuer'а если указан
	if *issuerAddress != "" {
		addr := common.HexToAddress(*issuerAddress)
		isTrusted, _ := issuerSvc.IsTrustedIssuer(ctx, addr)
		fmt.Printf("\n✓ Issuer %s: доверенный = %v\n", addr.Hex(), isTrusted)
	}
}

// Вспомогательные функции

func mustConnectToContracts() *blockchain.Client {
	if *issuerRegistryAddr == "" || *verificationRegistryAddr == "" || *nftContractAddr == "" {
		log.Fatal("Требуются адреса контрактов: --issuer-registry, --verification-registry, --nft-contract")
	}

	if *adminKey == "" {
		log.Fatal("Требуется --admin-key")
	}

	if *issuerKey == "" {
		*issuerKey = *adminKey // По умолчанию используем admin ключ
	}

	cfg := &blockchain.Config{
		RPCURL:                   *rpcURL,
		ChainID:                  31337,
		IssuerRegistryAddr:       common.HexToAddress(*issuerRegistryAddr),
		VerificationRegistryAddr: common.HexToAddress(*verificationRegistryAddr),
		NFTContractAddr:          common.HexToAddress(*nftContractAddr),
		AdminPrivateKey:          *adminKey,
		IssuerPrivateKey:         *issuerKey,
	}

	client, err := blockchain.NewClient(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}

	return client
}

func hashFromString(s string) [32]byte {
	// Если строка начинается с 0x - это уже хэш
	if len(s) > 2 && s[:2] == "0x" {
		var hash [32]byte
		decoded := common.FromHex(s)
		copy(hash[:], decoded)
		return hash
	}

	// Иначе хэшируем строку
	return sha256.Sum256([]byte(s))
}

func printHelp() {
	fmt.Println(`
🔧 NFT+VC Verification System

Режимы работы:

  deploy          Деплой всех контрактов
  add-issuer      Добавить доверенного issuer'а
  verify          Зарегистрировать верификацию документа
  check           Проверить статус токена
  revoke          Отозвать верификацию
  info            Показать информацию о системе

Примеры использования:

  # 1. Деплой контрактов
  go run cmd/main.go -mode deploy -admin-key YOUR_PRIVATE_KEY

  # 2. Добавить issuer'а
  go run cmd/main.go -mode add-issuer \
    -admin-key YOUR_PRIVATE_KEY \
    -issuer-addr 0x... \
    -issuer-registry 0x... \
    -verification-registry 0x... \
    -nft-contract 0x...

  # 3. Верифицировать документ
  go run cmd/main.go -mode verify \
    -issuer-key YOUR_ISSUER_KEY \
    -vc-hash "document content or 0x..." \
    -subject 0xUSER_ADDRESS \
    -issuer-registry 0x... \
    -verification-registry 0x... \
    -nft-contract 0x...

  # 4. Проверить статус
  go run cmd/main.go -mode check \
    -token-id 1 \
    -admin-key YOUR_KEY \
    -issuer-registry 0x... \
    -verification-registry 0x... \
    -nft-contract 0x...

  # 5. Отозвать верификацию
  go run cmd/main.go -mode revoke \
    -issuer-key YOUR_ISSUER_KEY \
    -token-id 1 \
    -issuer-registry 0x... \
    -verification-registry 0x... \
    -nft-contract 0x...

Флаги:
`)
	flag.PrintDefaults()
}
