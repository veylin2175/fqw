// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title VerifiedDocumentNFT
 * @notice NFT-токен как криптографический якорь верификации VC
 * @dev Контракт не хранит данные документа и не выполняет проверок
 */
contract VerifiedDocumentNFT is ERC721, Ownable(msg.sender) {

    /// @dev Адрес VerificationRegistry
    address public verificationRegistry;

    /// @dev Базовый URI для metadata (опционально)
    string private baseTokenURI;

    /// @dev Событие выпуска NFT
    event VerifiedNFTMinted(
        uint256 indexed tokenId,
        address indexed owner
    );

    /// @dev Модификатор: только VerificationRegistry
    modifier onlyVerificationRegistry() {
        require(
            msg.sender == verificationRegistry,
            "VerifiedDocumentNFT: not VerificationRegistry"
        );
        _;
    }

    /**
     * @notice Инициализация NFT-контракта
     * @param _name Название коллекции
     * @param _symbol Символ
     */
    constructor(
        string memory _name,
        string memory _symbol
    ) ERC721(_name, _symbol) {}

    /**
     * @notice Установка адреса VerificationRegistry
     * @dev Выполняется один раз при деплое системы
     */
    function setVerificationRegistry(address _registry) external onlyOwner {
        require(verificationRegistry == address(0), "Registry already set");
        require(_registry != address(0), "Zero address");
        verificationRegistry = _registry;
    }

    /**
     * @notice Выпуск NFT после успешной верификации
     * @param to Владелец NFT (Subject)
     * @param tokenId Идентификатор, полученный из VerificationRegistry
     */
    function mint(
        address to,
        uint256 tokenId
    ) external onlyVerificationRegistry {

        require(to != address(0), "Zero address");
        _safeMint(to, tokenId);

        emit VerifiedNFTMinted(tokenId, to);
    }

    /**
     * @notice Установка baseURI для metadata
     */
    function setBaseURI(string calldata uri) external onlyOwner {
        baseTokenURI = uri;
    }

    /**
     * @dev Переопределение baseURI
     */
    function _baseURI() internal view override returns (string memory) {
        return baseTokenURI;
    }
}
