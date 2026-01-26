// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IssuerRegistry
 * @author Veylin
 * @notice Реестр доверенных верифицирующих организаций (Issuer)
 * @dev Контракт отвечает только за управление доверием
 */
contract IssuerRegistry {

    /// @dev Администратор системы (System Authority)
    address public owner;

    /// @dev Реестр доверенных issuer'ов
    mapping(address => bool) private trustedIssuers;

    /// @dev События для аудита
    event IssuerAdded(address indexed issuer);
    event IssuerRemoved(address indexed issuer);
    event OwnershipTransferred(address indexed oldOwner, address indexed newOwner);

    /// @dev Модификатор доступа администратора
    modifier onlyOwner() {
        require(msg.sender == owner, "IssuerRegistry: not owner");
        _;
    }

    /// @notice Инициализация контракта
    constructor() {
        owner = msg.sender;
    }

    /**
     * @notice Добавить доверенную организацию
     * @param issuer Адрес организации
     */
    function addIssuer(address issuer) external onlyOwner {
        require(issuer != address(0), "IssuerRegistry: zero address");
        require(!trustedIssuers[issuer], "IssuerRegistry: already trusted");

        trustedIssuers[issuer] = true;
        emit IssuerAdded(issuer);
    }

    /**
     * @notice Удалить доверенную организацию
     * @param issuer Адрес организации
     */
    function removeIssuer(address issuer) external onlyOwner {
        require(trustedIssuers[issuer], "IssuerRegistry: not trusted");

        trustedIssuers[issuer] = false;
        emit IssuerRemoved(issuer);
    }

    /**
     * @notice Проверка, является ли адрес доверенным issuer'ом
     * @param issuer Адрес для проверки
     */
    function isTrustedIssuer(address issuer) external view returns (bool) {
        return trustedIssuers[issuer];
    }

    /**
     * @notice Передача прав администратора
     * @param newOwner Новый администратор
     */
    function transferOwnership(address newOwner) external onlyOwner {
        require(newOwner != address(0), "IssuerRegistry: zero address");

        address oldOwner = owner;
        owner = newOwner;

        emit OwnershipTransferred(oldOwner, newOwner);
    }
}
