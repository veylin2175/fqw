package api

import "time"

type VerifyDocumentRequest struct {
	VCHash         string `json:"vc_hash" binding:"required"`
	SubjectAddress string `json:"subject_address" binding:"required"`
}

type VerifyDocumentResponse struct {
	Success   bool   `json:"success"`
	TokenID   string `json:"token_id,omitempty"`
	TxHash    string `json:"tx_hash,omitempty"`
	NFTMinted bool   `json:"nft_minted"`
	Message   string `json:"message,omitempty"`
	Error     string `json:"error,omitempty"`
}

type CheckVerificationRequest struct {
	TokenID string `json:"token_id" binding:"required"`
	VCHash  string `json:"vc_hash,omitempty"`
}

type CheckVerificationResponse struct {
	Success   bool      `json:"success"`
	TokenID   string    `json:"token_id"`
	Issuer    string    `json:"issuer"`
	IssuedAt  time.Time `json:"issued_at"`
	Revoked   bool      `json:"revoked"`
	Valid     bool      `json:"valid"`
	NFTOwner  string    `json:"nft_owner,omitempty"`
	NFTExists bool      `json:"nft_exists"`
	Message   string    `json:"message,omitempty"`
	Error     string    `json:"error,omitempty"`
}

type RevokeVerificationRequest struct {
	TokenID string `json:"token_id" binding:"required"`
}

type RevokeVerificationResponse struct {
	Success bool   `json:"success"`
	TokenID string `json:"token_id"`
	TxHash  string `json:"tx_hash,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type SystemInfoResponse struct {
	Success              bool   `json:"success"`
	IssuerRegistryAddr   string `json:"issuer_registry_address"`
	VerificationRegistry string `json:"verification_registry_address"`
	NFTContractAddr      string `json:"nft_contract_address"`
	Owner                string `json:"owner"`
	TotalVerifications   string `json:"total_verifications"`
	NFTCollectionName    string `json:"nft_collection_name"`
	NFTCollectionSymbol  string `json:"nft_collection_symbol"`
}

type AddIssuerRequest struct {
	IssuerAddress string `json:"issuer_address" binding:"required"`
}

type AddIssuerResponse struct {
	Success       bool   `json:"success"`
	IssuerAddress string `json:"issuer_address"`
	TxHash        string `json:"tx_hash,omitempty"`
	IsTrusted     bool   `json:"is_trusted"`
	Message       string `json:"message,omitempty"`
	Error         string `json:"error,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    string `json:"code,omitempty"`
}
