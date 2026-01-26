// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./IssuerRegistry.sol";

/**
 * @title IVerifiedDocumentNFT
 * @notice Интерфейс для вызова mint из VerificationRegistry
 */
interface IVerifiedDocumentNFT {
    function mint(address to, uint256 tokenId) external;
}

/**
 * @title VerificationRegistry
 * @notice Реестр фактов верификации VC и управления их статусом
 * @dev Контракт фиксирует результат верификации и автоматически минтит NFT
 */
contract VerificationRegistry {
    /// @dev Структура результата верификации
    struct Verification {
        bytes32 vcHash;     // Хэш Verifiable Credential
        address issuer;     // Кто выполнил верификацию
        uint256 issuedAt;   // Время верификации
        bool revoked;       // Статус отзыва
    }

    /// @dev Реестр верификаций: tokenId -> Verification
    mapping(uint256 => Verification) private verifications;

    /// @dev Счётчик токенов
    uint256 public tokenCounter;

    /// @dev Контракт реестра доверенных issuer'ов
    IssuerRegistry public issuerRegistry;

    /// @dev Адрес NFT-контракта
    address public nftContract;

    /// @dev События
    event VerificationRegistered(
        uint256 indexed tokenId,
        address indexed issuer,
        bytes32 vcHash
    );
    event VerificationRevoked(uint256 indexed tokenId);
    event NFTMinted(uint256 indexed tokenId, address indexed owner);

    /// @dev Модификатор: только доверенный issuer
    modifier onlyTrustedIssuer() {
        require(
            issuerRegistry.isTrustedIssuer(msg.sender),
            "VerificationRegistry: not trusted issuer"
        );
        _;
    }

    /**
     * @notice Инициализация
     * @param _issuerRegistry Адрес IssuerRegistry
     */
    constructor(address _issuerRegistry) {
        require(_issuerRegistry != address(0), "VerificationRegistry: zero address");
        issuerRegistry = IssuerRegistry(_issuerRegistry);
    }

    /**
     * @notice Установка адреса NFT-контракта
     * @dev Вызывается один раз при деплое системы
     */
    function setNFTContract(address _nftContract) external {
        require(nftContract == address(0), "VerificationRegistry: NFT already set");
        require(_nftContract != address(0), "VerificationRegistry: zero address");
        nftContract = _nftContract;
    }

    /**
     * @notice Регистрация факта верификации с автоматическим минтом NFT
     * @param vcHash Хэш VC
     * @param subject Владелец NFT (Subject)
     */
    function registerVerification(
        bytes32 vcHash,
        address subject
    ) external onlyTrustedIssuer returns (uint256) {
        require(vcHash != bytes32(0), "VerificationRegistry: empty hash");
        require(subject != address(0), "VerificationRegistry: zero subject");
        require(nftContract != address(0), "VerificationRegistry: NFT not set");

        uint256 tokenId = ++tokenCounter;

        verifications[tokenId] = Verification({
            vcHash: vcHash,
            issuer: msg.sender,
            issuedAt: block.timestamp,
            revoked: false
        });

        emit VerificationRegistered(tokenId, msg.sender, vcHash);

        // Автоматически минтим NFT
        IVerifiedDocumentNFT(nftContract).mint(subject, tokenId);
        emit NFTMinted(tokenId, subject);

        return tokenId;
    }

    /**
     * @notice Отзыв верификации
     * @param tokenId Идентификатор NFT
     */
    function revokeVerification(uint256 tokenId) external {
        Verification storage v = verifications[tokenId];
        require(v.issuer == msg.sender, "VerificationRegistry: not issuer");
        require(!v.revoked, "VerificationRegistry: already revoked");

        v.revoked = true;
        emit VerificationRevoked(tokenId);
    }

    /**
     * @notice Проверка подлинности VC
     * @param tokenId Идентификатор NFT
     * @param vcHash Хэш проверяемого VC
     */
    function verify(
        uint256 tokenId,
        bytes32 vcHash
    ) external view returns (
        bool valid,
        address issuer,
        bool revoked,
        uint256 issuedAt
    ) {
        Verification memory v = verifications[tokenId];
        valid = (v.vcHash == vcHash && !v.revoked);
        issuer = v.issuer;
        revoked = v.revoked;
        issuedAt = v.issuedAt;
    }
}