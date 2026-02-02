package api

import (
	"context"
	"crypto/sha256"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"fqw/internal/blockchain"
	"fqw/internal/service"
)

type Handler struct {
	client          *blockchain.Client
	issuerSvc       *service.IssuerService
	verificationSvc *service.VerificationService
	nftSvc          *service.NFTService
}

func NewHandler(client *blockchain.Client) *Handler {
	return &Handler{
		client:          client,
		issuerSvc:       service.NewIssuerService(client),
		verificationSvc: service.NewVerificationService(client),
		nftSvc:          service.NewNFTService(client),
	}
}

func (h *Handler) VerifyDocument(c *gin.Context) {
	var req VerifyDocumentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
			Code:    "INVALID_REQUEST",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	vcHash := hashFromString(req.VCHash)
	subject := common.HexToAddress(req.SubjectAddress)

	tokenID, err := h.verificationSvc.RegisterVerification(ctx, vcHash, subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, VerifyDocumentResponse{
			Success: false,
			Error:   err.Error(),
			Message: "Failed to register verification",
		})
		return
	}

	c.JSON(http.StatusOK, VerifyDocumentResponse{
		Success:   true,
		TokenID:   tokenID.String(),
		NFTMinted: true,
		Message:   "Document verified and NFT minted successfully",
	})
}

func (h *Handler) CheckVerification(c *gin.Context) {
	tokenIDStr := c.Param("tokenId")
	vcHashStr := c.Query("vc_hash")

	tokenID := new(big.Int)
	_, ok := tokenID.SetString(tokenIDStr, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid token ID",
			Code:    "INVALID_TOKEN_ID",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var result *service.VerificationResult
	var err error

	if vcHashStr != "" {
		vcHash := hashFromString(vcHashStr)
		result, err = h.verificationSvc.VerifyDocument(ctx, tokenID, vcHash)
	} else {
		result, err = h.verificationSvc.GetVerificationDetails(ctx, tokenID)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, CheckVerificationResponse{
			Success: false,
			Error:   err.Error(),
			Message: "Failed to check verification",
		})
		return
	}

	owner, err := h.nftSvc.GetOwnerOf(ctx, tokenID)
	nftExists := (err == nil && owner != common.Address{})

	response := CheckVerificationResponse{
		Success:   true,
		TokenID:   tokenID.String(),
		Issuer:    result.Issuer.Hex(),
		IssuedAt:  time.Unix(result.IssuedAt.Int64(), 0),
		Revoked:   result.Revoked,
		Valid:     result.Valid,
		NFTExists: nftExists,
	}

	if nftExists {
		response.NFTOwner = owner.Hex()
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) RevokeVerification(c *gin.Context) {
	tokenIDStr := c.Param("tokenId")

	tokenID := new(big.Int)
	_, ok := tokenID.SetString(tokenIDStr, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid token ID",
			Code:    "INVALID_TOKEN_ID",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	err := h.verificationSvc.RevokeVerification(ctx, tokenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RevokeVerificationResponse{
			Success: false,
			TokenID: tokenID.String(),
			Error:   err.Error(),
			Message: "Failed to revoke verification",
		})
		return
	}

	c.JSON(http.StatusOK, RevokeVerificationResponse{
		Success: true,
		TokenID: tokenID.String(),
		Message: "Verification revoked successfully",
	})
}

func (h *Handler) GetSystemInfo(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	owner, _ := h.issuerSvc.GetOwner(ctx)
	counter, _ := h.verificationSvc.GetTokenCounter(ctx)
	name, _ := h.nftSvc.GetName(ctx)
	symbol, _ := h.nftSvc.GetSymbol(ctx)

	c.JSON(http.StatusOK, SystemInfoResponse{
		Success:              true,
		IssuerRegistryAddr:   h.client.IssuerRegistryAddr().Hex(),
		VerificationRegistry: h.client.VerificationRegistryAddr().Hex(),
		NFTContractAddr:      h.client.NFTContractAddr().Hex(),
		Owner:                owner.Hex(),
		TotalVerifications:   counter.String(),
		NFTCollectionName:    name,
		NFTCollectionSymbol:  symbol,
	})
}

func (h *Handler) AddIssuer(c *gin.Context) {
	var req AddIssuerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid request: " + err.Error(),
			Code:    "INVALID_REQUEST",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	issuerAddr := common.HexToAddress(req.IssuerAddress)

	err := h.issuerSvc.AddTrustedIssuer(ctx, issuerAddr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AddIssuerResponse{
			Success:       false,
			IssuerAddress: issuerAddr.Hex(),
			Error:         err.Error(),
			Message:       "Failed to add issuer",
		})
		return
	}

	isTrusted, _ := h.issuerSvc.IsTrustedIssuer(ctx, issuerAddr)

	c.JSON(http.StatusOK, AddIssuerResponse{
		Success:       true,
		IssuerAddress: issuerAddr.Hex(),
		IsTrusted:     isTrusted,
		Message:       "Issuer added successfully",
	})
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Unix(),
	})
}

func hashFromString(s string) [32]byte {
	if len(s) > 2 && s[:2] == "0x" {
		var hash [32]byte
		decoded := common.FromHex(s)
		copy(hash[:], decoded)
		return hash
	}
	return sha256.Sum256([]byte(s))
}
