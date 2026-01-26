// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VerifiedDocumentNFTMetaData contains all meta data concerning the VerifiedDocumentNFT contract.
var VerifiedDocumentNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"VerifiedNFTMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setVerificationRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051612eae380380612eae833981810160405281019061003191906102ec565b338282815f90816100429190610583565b5080600190816100529190610583565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100c5575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100bc9190610691565b60405180910390fd5b6100d4816100dc60201b60201c565b5050506106aa565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101fe826101b8565b810181811067ffffffffffffffff8211171561021d5761021c6101c8565b5b80604052505050565b5f61022f61019f565b905061023b82826101f5565b919050565b5f67ffffffffffffffff82111561025a576102596101c8565b5b610263826101b8565b9050602081019050919050565b8281835e5f83830152505050565b5f61029061028b84610240565b610226565b9050828152602081018484840111156102ac576102ab6101b4565b5b6102b7848285610270565b509392505050565b5f82601f8301126102d3576102d26101b0565b5b81516102e384826020860161027e565b91505092915050565b5f5f60408385031215610302576103016101a8565b5b5f83015167ffffffffffffffff81111561031f5761031e6101ac565b5b61032b858286016102bf565b925050602083015167ffffffffffffffff81111561034c5761034b6101ac565b5b610358858286016102bf565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103b057607f821691505b6020821081036103c3576103c261036c565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104257fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103ea565b61042f86836103ea565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61047361046e61046984610447565b610450565b610447565b9050919050565b5f819050919050565b61048c83610459565b6104a06104988261047a565b8484546103f6565b825550505050565b5f5f905090565b6104b76104a8565b6104c2818484610483565b505050565b5f5b828110156104e8576104dd5f8284016104af565b6001810190506104c9565b505050565b601f82111561053b578282111561053a57610507816103c9565b610510836103db565b610519856103db565b6020861015610526575f90505b808301610535828403826104c7565b505050505b5b505050565b5f82821c905092915050565b5f61055b5f1984600802610540565b1980831691505092915050565b5f610573838361054c565b9150826002028217905092915050565b61058c82610362565b67ffffffffffffffff8111156105a5576105a46101c8565b5b6105af8254610399565b6105ba8282856104ed565b5f60209050601f8311600181146105eb575f84156105d9578287015190505b6105e38582610568565b86555061064a565b601f1984166105f9866103c9565b5f5b82811015610620578489015182556001820191506020850194506020810190506105fb565b8683101561063d5784890151610639601f89168261054c565b8355505b6001600288020188555050505b505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61067b82610652565b9050919050565b61068b81610671565b82525050565b5f6020820190506106a45f830184610682565b92915050565b6127f7806106b75f395ff3fe608060405234801561000f575f5ffd5b506004361061012a575f3560e01c8063715018a6116100ab578063b88d4fde1161006f578063b88d4fde14610316578063c87b56dd14610332578063e985e9c514610362578063efa4f8d414610392578063f2fde38b146103b05761012a565b8063715018a61461029857806387664f91146102a25780638da5cb5b146102be57806395d89b41146102dc578063a22cb465146102fa5761012a565b806340c10f19116100f257806340c10f19146101e457806342842e0e1461020057806355f804b31461021c5780636352211e1461023857806370a08231146102685761012a565b806301ffc9a71461012e57806306fdde031461015e578063081812fc1461017c578063095ea7b3146101ac57806323b872dd146101c8575b5f5ffd5b61014860048036038101906101439190611c14565b6103cc565b6040516101559190611c59565b60405180910390f35b6101666104ad565b6040516101739190611ce2565b60405180910390f35b61019660048036038101906101919190611d35565b61053c565b6040516101a39190611d9f565b60405180910390f35b6101c660048036038101906101c19190611de2565b610557565b005b6101e260048036038101906101dd9190611e20565b61056d565b005b6101fe60048036038101906101f99190611de2565b61066c565b005b61021a60048036038101906102159190611e20565b6107bb565b005b61023660048036038101906102319190611ed1565b6107da565b005b610252600480360381019061024d9190611d35565b6107f8565b60405161025f9190611d9f565b60405180910390f35b610282600480360381019061027d9190611f1c565b610809565b60405161028f9190611f56565b60405180910390f35b6102a06108bf565b005b6102bc60048036038101906102b79190611f1c565b6108d2565b005b6102c6610a1a565b6040516102d39190611d9f565b60405180910390f35b6102e4610a42565b6040516102f19190611ce2565b60405180910390f35b610314600480360381019061030f9190611f99565b610ad2565b005b610330600480360381019061032b91906120ff565b610ae8565b005b61034c60048036038101906103479190611d35565b610b0d565b6040516103599190611ce2565b60405180910390f35b61037c6004803603810190610377919061217f565b610b73565b6040516103899190611c59565b60405180910390f35b61039a610c01565b6040516103a79190611d9f565b60405180910390f35b6103ca60048036038101906103c59190611f1c565b610c26565b005b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061049657507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806104a657506104a582610caa565b5b9050919050565b60605f80546104bb906121ea565b80601f01602080910402602001604051908101604052809291908181526020018280546104e7906121ea565b80156105325780601f1061050957610100808354040283529160200191610532565b820191905f5260205f20905b81548152906001019060200180831161051557829003601f168201915b5050505050905090565b5f61054682610d13565b5061055082610d99565b9050919050565b6105698282610564610dd2565b610dd9565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105dd575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016105d49190611d9f565b60405180910390fd5b5f6105f083836105eb610dd2565b610deb565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610666578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161065d9392919061221a565b60405180910390fd5b50505050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f2906122bf565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610769576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076090612327565b60405180910390fd5b6107738282610ff6565b8173ffffffffffffffffffffffffffffffffffffffff16817f759831dc39e06177efad640f26e510fec0996d44554975f64801fcfe85e7585060405160405180910390a35050565b6107d583838360405180602001604052805f815250610ae8565b505050565b6107e2611013565b8181600891826107f3929190612500565b505050565b5f61080282610d13565b9050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361087a575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016108719190611d9f565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6108c7611013565b6108d05f61109a565b565b6108da611013565b5f73ffffffffffffffffffffffffffffffffffffffff1660075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610969576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096090612617565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ce90612327565b60405180910390fd5b8060075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060018054610a51906121ea565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7d906121ea565b8015610ac85780601f10610a9f57610100808354040283529160200191610ac8565b820191905f5260205f20905b815481529060010190602001808311610aab57829003601f168201915b5050505050905090565b610ae4610add610dd2565b838361115d565b5050565b610af384848461056d565b610b07610afe610dd2565b858585856112c6565b50505050565b6060610b1882610d13565b505f610b22611472565b90505f815111610b405760405180602001604052805f815250610b6b565b80610b4a84611502565b604051602001610b5b92919061266f565b6040516020818303038152906040525b915050919050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610c2e611013565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c9e575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610c959190611d9f565b60405180910390fd5b610ca78161109a565b50565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f5f610d1e836115cc565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d9057826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610d879190611f56565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610de68383836001611605565b505050565b5f5f610df6846115cc565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610e3757610e368184866117c4565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ec257610e765f855f5f611605565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610f4157600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b61100f828260405180602001604052805f815250611887565b5050565b61101b610dd2565b73ffffffffffffffffffffffffffffffffffffffff16611039610a1a565b73ffffffffffffffffffffffffffffffffffffffff16146110985761105c610dd2565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161108f9190611d9f565b60405180910390fd5b565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036111cd57816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016111c49190611d9f565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516112b99190611c59565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b111561146b578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b815260040161132494939291906126e4565b6020604051808303815f875af192505050801561135f57506040513d601f19601f8201168201806040525081019061135c9190612742565b60015b6113e0573d805f811461138d576040519150601f19603f3d011682016040523d82523d5f602084013e611392565b606091505b505f8151036113d857836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016113cf9190611d9f565b60405180910390fd5b805160208201fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461146957836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016114609190611d9f565b60405180910390fd5b505b5050505050565b606060088054611481906121ea565b80601f01602080910402602001604051908101604052809291908181526020018280546114ad906121ea565b80156114f85780601f106114cf576101008083540402835291602001916114f8565b820191905f5260205f20905b8154815290600101906020018083116114db57829003601f168201915b5050505050905090565b60605f6001611510846118aa565b0190505f8167ffffffffffffffff81111561152e5761152d611fdb565b5b6040519080825280601f01601f1916602001820160405280156115605781602001600182028036833780820191505090505b5090505f82602083010190505b6001156115c1578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816115b6576115b561276d565b5b0494505f850361156d575b819350505050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061163d57505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1561176f575f61164c84610d13565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156116b657508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156116c957506116c78184610b73565b155b1561170b57826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016117029190611d9f565b60405180910390fd5b811561176d57838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6117cf8383836119fb565b611882575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361184357806040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161183a9190611f56565b60405180910390fd5b81816040517f177e802f00000000000000000000000000000000000000000000000000000000815260040161187992919061279a565b60405180910390fd5b505050565b6118918383611abb565b6118a561189c610dd2565b5f8585856112c6565b505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611906577a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083816118fc576118fb61276d565b5b0492506040810190505b6d04ee2d6d415b85acef81000000008310611943576d04ee2d6d415b85acef810000000083816119395761193861276d565b5b0492506020810190505b662386f26fc10000831061197257662386f26fc1000083816119685761196761276d565b5b0492506010810190505b6305f5e100831061199b576305f5e10083816119915761199061276d565b5b0492506008810190505b61271083106119c05761271083816119b6576119b561276d565b5b0492506004810190505b606483106119e357606483816119d9576119d861276d565b5b0492506002810190505b600a83106119f2576001810190505b80915050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015611ab257508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480611a735750611a728484610b73565b5b80611ab157508273ffffffffffffffffffffffffffffffffffffffff16611a9983610d99565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611b2b575f6040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401611b229190611d9f565b60405180910390fd5b5f611b3783835f610deb565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611ba9575f6040517f73c6ac6e000000000000000000000000000000000000000000000000000000008152600401611ba09190611d9f565b60405180910390fd5b505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611bf381611bbf565b8114611bfd575f5ffd5b50565b5f81359050611c0e81611bea565b92915050565b5f60208284031215611c2957611c28611bb7565b5b5f611c3684828501611c00565b91505092915050565b5f8115159050919050565b611c5381611c3f565b82525050565b5f602082019050611c6c5f830184611c4a565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611cb482611c72565b611cbe8185611c7c565b9350611cce818560208601611c8c565b611cd781611c9a565b840191505092915050565b5f6020820190508181035f830152611cfa8184611caa565b905092915050565b5f819050919050565b611d1481611d02565b8114611d1e575f5ffd5b50565b5f81359050611d2f81611d0b565b92915050565b5f60208284031215611d4a57611d49611bb7565b5b5f611d5784828501611d21565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611d8982611d60565b9050919050565b611d9981611d7f565b82525050565b5f602082019050611db25f830184611d90565b92915050565b611dc181611d7f565b8114611dcb575f5ffd5b50565b5f81359050611ddc81611db8565b92915050565b5f5f60408385031215611df857611df7611bb7565b5b5f611e0585828601611dce565b9250506020611e1685828601611d21565b9150509250929050565b5f5f5f60608486031215611e3757611e36611bb7565b5b5f611e4486828701611dce565b9350506020611e5586828701611dce565b9250506040611e6686828701611d21565b9150509250925092565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112611e9157611e90611e70565b5b8235905067ffffffffffffffff811115611eae57611ead611e74565b5b602083019150836001820283011115611eca57611ec9611e78565b5b9250929050565b5f5f60208385031215611ee757611ee6611bb7565b5b5f83013567ffffffffffffffff811115611f0457611f03611bbb565b5b611f1085828601611e7c565b92509250509250929050565b5f60208284031215611f3157611f30611bb7565b5b5f611f3e84828501611dce565b91505092915050565b611f5081611d02565b82525050565b5f602082019050611f695f830184611f47565b92915050565b611f7881611c3f565b8114611f82575f5ffd5b50565b5f81359050611f9381611f6f565b92915050565b5f5f60408385031215611faf57611fae611bb7565b5b5f611fbc85828601611dce565b9250506020611fcd85828601611f85565b9150509250929050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61201182611c9a565b810181811067ffffffffffffffff821117156120305761202f611fdb565b5b80604052505050565b5f612042611bae565b905061204e8282612008565b919050565b5f67ffffffffffffffff82111561206d5761206c611fdb565b5b61207682611c9a565b9050602081019050919050565b828183375f83830152505050565b5f6120a361209e84612053565b612039565b9050828152602081018484840111156120bf576120be611fd7565b5b6120ca848285612083565b509392505050565b5f82601f8301126120e6576120e5611e70565b5b81356120f6848260208601612091565b91505092915050565b5f5f5f5f6080858703121561211757612116611bb7565b5b5f61212487828801611dce565b945050602061213587828801611dce565b935050604061214687828801611d21565b925050606085013567ffffffffffffffff81111561216757612166611bbb565b5b612173878288016120d2565b91505092959194509250565b5f5f6040838503121561219557612194611bb7565b5b5f6121a285828601611dce565b92505060206121b385828601611dce565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061220157607f821691505b602082108103612214576122136121bd565b5b50919050565b5f60608201905061222d5f830186611d90565b61223a6020830185611f47565b6122476040830184611d90565b949350505050565b7f5665726966696564446f63756d656e744e46543a206e6f7420566572696669635f8201527f6174696f6e526567697374727900000000000000000000000000000000000000602082015250565b5f6122a9602d83611c7c565b91506122b48261224f565b604082019050919050565b5f6020820190508181035f8301526122d68161229d565b9050919050565b7f5a65726f206164647265737300000000000000000000000000000000000000005f82015250565b5f612311600c83611c7c565b915061231c826122dd565b602082019050919050565b5f6020820190508181035f83015261233e81612305565b9050919050565b5f82905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026123ab7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612370565b6123b58683612370565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6123f06123eb6123e684611d02565b6123cd565b611d02565b9050919050565b5f819050919050565b612409836123d6565b61241d612415826123f7565b84845461237c565b825550505050565b5f5f905090565b612434612425565b61243f818484612400565b505050565b5f5b828110156124655761245a5f82840161242c565b600181019050612446565b505050565b601f8211156124b857828211156124b7576124848161234f565b61248d83612361565b61249685612361565b60208610156124a3575f90505b8083016124b282840382612444565b505050505b5b505050565b5f82821c905092915050565b5f6124d85f19846008026124bd565b1980831691505092915050565b5f6124f083836124c9565b9150826002028217905092915050565b61250a8383612345565b67ffffffffffffffff81111561252357612522611fdb565b5b61252d82546121ea565b61253882828561246a565b5f601f831160018114612565575f8415612553578287013590505b61255d85826124e5565b8655506125c4565b601f1984166125738661234f565b5f5b8281101561259a57848901358255600182019150602085019450602081019050612575565b868310156125b757848901356125b3601f8916826124c9565b8355505b6001600288020188555050505b50505050505050565b7f526567697374727920616c7265616479207365740000000000000000000000005f82015250565b5f612601601483611c7c565b915061260c826125cd565b602082019050919050565b5f6020820190508181035f83015261262e816125f5565b9050919050565b5f81905092915050565b5f61264982611c72565b6126538185612635565b9350612663818560208601611c8c565b80840191505092915050565b5f61267a828561263f565b9150612686828461263f565b91508190509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f6126b682612692565b6126c0818561269c565b93506126d0818560208601611c8c565b6126d981611c9a565b840191505092915050565b5f6080820190506126f75f830187611d90565b6127046020830186611d90565b6127116040830185611f47565b818103606083015261272381846126ac565b905095945050505050565b5f8151905061273c81611bea565b92915050565b5f6020828403121561275757612756611bb7565b5b5f6127648482850161272e565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6040820190506127ad5f830185611d90565b6127ba6020830184611f47565b939250505056fea2646970667358221220cd0ef3b15e226cbfcf8e63b7770fba4a68d1a1bd873001cede9b1e80e67337f064736f6c63430008210033",
}

// VerifiedDocumentNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifiedDocumentNFTMetaData.ABI instead.
var VerifiedDocumentNFTABI = VerifiedDocumentNFTMetaData.ABI

// VerifiedDocumentNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifiedDocumentNFTMetaData.Bin instead.
var VerifiedDocumentNFTBin = VerifiedDocumentNFTMetaData.Bin

// DeployVerifiedDocumentNFT deploys a new Ethereum contract, binding an instance of VerifiedDocumentNFT to it.
func DeployVerifiedDocumentNFT(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *VerifiedDocumentNFT, error) {
	parsed, err := VerifiedDocumentNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifiedDocumentNFTBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifiedDocumentNFT{VerifiedDocumentNFTCaller: VerifiedDocumentNFTCaller{contract: contract}, VerifiedDocumentNFTTransactor: VerifiedDocumentNFTTransactor{contract: contract}, VerifiedDocumentNFTFilterer: VerifiedDocumentNFTFilterer{contract: contract}}, nil
}

// VerifiedDocumentNFT is an auto generated Go binding around an Ethereum contract.
type VerifiedDocumentNFT struct {
	VerifiedDocumentNFTCaller     // Read-only binding to the contract
	VerifiedDocumentNFTTransactor // Write-only binding to the contract
	VerifiedDocumentNFTFilterer   // Log filterer for contract events
}

// VerifiedDocumentNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifiedDocumentNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiedDocumentNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifiedDocumentNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiedDocumentNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifiedDocumentNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiedDocumentNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifiedDocumentNFTSession struct {
	Contract     *VerifiedDocumentNFT // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VerifiedDocumentNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifiedDocumentNFTCallerSession struct {
	Contract *VerifiedDocumentNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VerifiedDocumentNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifiedDocumentNFTTransactorSession struct {
	Contract     *VerifiedDocumentNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VerifiedDocumentNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifiedDocumentNFTRaw struct {
	Contract *VerifiedDocumentNFT // Generic contract binding to access the raw methods on
}

// VerifiedDocumentNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifiedDocumentNFTCallerRaw struct {
	Contract *VerifiedDocumentNFTCaller // Generic read-only contract binding to access the raw methods on
}

// VerifiedDocumentNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifiedDocumentNFTTransactorRaw struct {
	Contract *VerifiedDocumentNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifiedDocumentNFT creates a new instance of VerifiedDocumentNFT, bound to a specific deployed contract.
func NewVerifiedDocumentNFT(address common.Address, backend bind.ContractBackend) (*VerifiedDocumentNFT, error) {
	contract, err := bindVerifiedDocumentNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFT{VerifiedDocumentNFTCaller: VerifiedDocumentNFTCaller{contract: contract}, VerifiedDocumentNFTTransactor: VerifiedDocumentNFTTransactor{contract: contract}, VerifiedDocumentNFTFilterer: VerifiedDocumentNFTFilterer{contract: contract}}, nil
}

// NewVerifiedDocumentNFTCaller creates a new read-only instance of VerifiedDocumentNFT, bound to a specific deployed contract.
func NewVerifiedDocumentNFTCaller(address common.Address, caller bind.ContractCaller) (*VerifiedDocumentNFTCaller, error) {
	contract, err := bindVerifiedDocumentNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFTCaller{contract: contract}, nil
}

// NewVerifiedDocumentNFTTransactor creates a new write-only instance of VerifiedDocumentNFT, bound to a specific deployed contract.
func NewVerifiedDocumentNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifiedDocumentNFTTransactor, error) {
	contract, err := bindVerifiedDocumentNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFTTransactor{contract: contract}, nil
}

// NewVerifiedDocumentNFTFilterer creates a new log filterer instance of VerifiedDocumentNFT, bound to a specific deployed contract.
func NewVerifiedDocumentNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifiedDocumentNFTFilterer, error) {
	contract, err := bindVerifiedDocumentNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFTFilterer{contract: contract}, nil
}

// bindVerifiedDocumentNFT binds a generic wrapper to an already deployed contract.
func bindVerifiedDocumentNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifiedDocumentNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifiedDocumentNFT *VerifiedDocumentNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifiedDocumentNFT.Contract.VerifiedDocumentNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifiedDocumentNFT *VerifiedDocumentNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.VerifiedDocumentNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifiedDocumentNFT *VerifiedDocumentNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.VerifiedDocumentNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifiedDocumentNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VerifiedDocumentNFT.Contract.BalanceOf(&_VerifiedDocumentNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VerifiedDocumentNFT.Contract.BalanceOf(&_VerifiedDocumentNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _VerifiedDocumentNFT.Contract.GetApproved(&_VerifiedDocumentNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _VerifiedDocumentNFT.Contract.GetApproved(&_VerifiedDocumentNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _VerifiedDocumentNFT.Contract.IsApprovedForAll(&_VerifiedDocumentNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _VerifiedDocumentNFT.Contract.IsApprovedForAll(&_VerifiedDocumentNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) Name() (string, error) {
	return _VerifiedDocumentNFT.Contract.Name(&_VerifiedDocumentNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) Name() (string, error) {
	return _VerifiedDocumentNFT.Contract.Name(&_VerifiedDocumentNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) Owner() (common.Address, error) {
	return _VerifiedDocumentNFT.Contract.Owner(&_VerifiedDocumentNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) Owner() (common.Address, error) {
	return _VerifiedDocumentNFT.Contract.Owner(&_VerifiedDocumentNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _VerifiedDocumentNFT.Contract.OwnerOf(&_VerifiedDocumentNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _VerifiedDocumentNFT.Contract.OwnerOf(&_VerifiedDocumentNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VerifiedDocumentNFT.Contract.SupportsInterface(&_VerifiedDocumentNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VerifiedDocumentNFT.Contract.SupportsInterface(&_VerifiedDocumentNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) Symbol() (string, error) {
	return _VerifiedDocumentNFT.Contract.Symbol(&_VerifiedDocumentNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) Symbol() (string, error) {
	return _VerifiedDocumentNFT.Contract.Symbol(&_VerifiedDocumentNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _VerifiedDocumentNFT.Contract.TokenURI(&_VerifiedDocumentNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _VerifiedDocumentNFT.Contract.TokenURI(&_VerifiedDocumentNFT.CallOpts, tokenId)
}

// VerificationRegistry is a free data retrieval call binding the contract method 0xefa4f8d4.
//
// Solidity: function verificationRegistry() view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCaller) VerificationRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifiedDocumentNFT.contract.Call(opts, &out, "verificationRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerificationRegistry is a free data retrieval call binding the contract method 0xefa4f8d4.
//
// Solidity: function verificationRegistry() view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) VerificationRegistry() (common.Address, error) {
	return _VerifiedDocumentNFT.Contract.VerificationRegistry(&_VerifiedDocumentNFT.CallOpts)
}

// VerificationRegistry is a free data retrieval call binding the contract method 0xefa4f8d4.
//
// Solidity: function verificationRegistry() view returns(address)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTCallerSession) VerificationRegistry() (common.Address, error) {
	return _VerifiedDocumentNFT.Contract.VerificationRegistry(&_VerifiedDocumentNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.Approve(&_VerifiedDocumentNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.Approve(&_VerifiedDocumentNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.Mint(&_VerifiedDocumentNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.Mint(&_VerifiedDocumentNFT.TransactOpts, to, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.RenounceOwnership(&_VerifiedDocumentNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.RenounceOwnership(&_VerifiedDocumentNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SafeTransferFrom(&_VerifiedDocumentNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SafeTransferFrom(&_VerifiedDocumentNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SafeTransferFrom0(&_VerifiedDocumentNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SafeTransferFrom0(&_VerifiedDocumentNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SetApprovalForAll(&_VerifiedDocumentNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SetApprovalForAll(&_VerifiedDocumentNFT.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) SetBaseURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "setBaseURI", uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SetBaseURI(&_VerifiedDocumentNFT.TransactOpts, uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string uri) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) SetBaseURI(uri string) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SetBaseURI(&_VerifiedDocumentNFT.TransactOpts, uri)
}

// SetVerificationRegistry is a paid mutator transaction binding the contract method 0x87664f91.
//
// Solidity: function setVerificationRegistry(address _registry) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) SetVerificationRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "setVerificationRegistry", _registry)
}

// SetVerificationRegistry is a paid mutator transaction binding the contract method 0x87664f91.
//
// Solidity: function setVerificationRegistry(address _registry) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) SetVerificationRegistry(_registry common.Address) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SetVerificationRegistry(&_VerifiedDocumentNFT.TransactOpts, _registry)
}

// SetVerificationRegistry is a paid mutator transaction binding the contract method 0x87664f91.
//
// Solidity: function setVerificationRegistry(address _registry) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) SetVerificationRegistry(_registry common.Address) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.SetVerificationRegistry(&_VerifiedDocumentNFT.TransactOpts, _registry)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.TransferFrom(&_VerifiedDocumentNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.TransferFrom(&_VerifiedDocumentNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.TransferOwnership(&_VerifiedDocumentNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifiedDocumentNFT *VerifiedDocumentNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifiedDocumentNFT.Contract.TransferOwnership(&_VerifiedDocumentNFT.TransactOpts, newOwner)
}

// VerifiedDocumentNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTApprovalIterator struct {
	Event *VerifiedDocumentNFTApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VerifiedDocumentNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiedDocumentNFTApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VerifiedDocumentNFTApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VerifiedDocumentNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiedDocumentNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiedDocumentNFTApproval represents a Approval event raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*VerifiedDocumentNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFTApprovalIterator{contract: _VerifiedDocumentNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VerifiedDocumentNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiedDocumentNFTApproval)
				if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) ParseApproval(log types.Log) (*VerifiedDocumentNFTApproval, error) {
	event := new(VerifiedDocumentNFTApproval)
	if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiedDocumentNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTApprovalForAllIterator struct {
	Event *VerifiedDocumentNFTApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VerifiedDocumentNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiedDocumentNFTApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VerifiedDocumentNFTApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VerifiedDocumentNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiedDocumentNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiedDocumentNFTApprovalForAll represents a ApprovalForAll event raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*VerifiedDocumentNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFTApprovalForAllIterator{contract: _VerifiedDocumentNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VerifiedDocumentNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiedDocumentNFTApprovalForAll)
				if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) ParseApprovalForAll(log types.Log) (*VerifiedDocumentNFTApprovalForAll, error) {
	event := new(VerifiedDocumentNFTApprovalForAll)
	if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiedDocumentNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTOwnershipTransferredIterator struct {
	Event *VerifiedDocumentNFTOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VerifiedDocumentNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiedDocumentNFTOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VerifiedDocumentNFTOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VerifiedDocumentNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiedDocumentNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiedDocumentNFTOwnershipTransferred represents a OwnershipTransferred event raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifiedDocumentNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFTOwnershipTransferredIterator{contract: _VerifiedDocumentNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifiedDocumentNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiedDocumentNFTOwnershipTransferred)
				if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) ParseOwnershipTransferred(log types.Log) (*VerifiedDocumentNFTOwnershipTransferred, error) {
	event := new(VerifiedDocumentNFTOwnershipTransferred)
	if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiedDocumentNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTTransferIterator struct {
	Event *VerifiedDocumentNFTTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VerifiedDocumentNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiedDocumentNFTTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VerifiedDocumentNFTTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VerifiedDocumentNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiedDocumentNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiedDocumentNFTTransfer represents a Transfer event raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*VerifiedDocumentNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFTTransferIterator{contract: _VerifiedDocumentNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VerifiedDocumentNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiedDocumentNFTTransfer)
				if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) ParseTransfer(log types.Log) (*VerifiedDocumentNFTTransfer, error) {
	event := new(VerifiedDocumentNFTTransfer)
	if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiedDocumentNFTVerifiedNFTMintedIterator is returned from FilterVerifiedNFTMinted and is used to iterate over the raw logs and unpacked data for VerifiedNFTMinted events raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTVerifiedNFTMintedIterator struct {
	Event *VerifiedDocumentNFTVerifiedNFTMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VerifiedDocumentNFTVerifiedNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiedDocumentNFTVerifiedNFTMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VerifiedDocumentNFTVerifiedNFTMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VerifiedDocumentNFTVerifiedNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiedDocumentNFTVerifiedNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiedDocumentNFTVerifiedNFTMinted represents a VerifiedNFTMinted event raised by the VerifiedDocumentNFT contract.
type VerifiedDocumentNFTVerifiedNFTMinted struct {
	TokenId *big.Int
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVerifiedNFTMinted is a free log retrieval operation binding the contract event 0x759831dc39e06177efad640f26e510fec0996d44554975f64801fcfe85e75850.
//
// Solidity: event VerifiedNFTMinted(uint256 indexed tokenId, address indexed owner)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) FilterVerifiedNFTMinted(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*VerifiedDocumentNFTVerifiedNFTMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.FilterLogs(opts, "VerifiedNFTMinted", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VerifiedDocumentNFTVerifiedNFTMintedIterator{contract: _VerifiedDocumentNFT.contract, event: "VerifiedNFTMinted", logs: logs, sub: sub}, nil
}

// WatchVerifiedNFTMinted is a free log subscription operation binding the contract event 0x759831dc39e06177efad640f26e510fec0996d44554975f64801fcfe85e75850.
//
// Solidity: event VerifiedNFTMinted(uint256 indexed tokenId, address indexed owner)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) WatchVerifiedNFTMinted(opts *bind.WatchOpts, sink chan<- *VerifiedDocumentNFTVerifiedNFTMinted, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VerifiedDocumentNFT.contract.WatchLogs(opts, "VerifiedNFTMinted", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiedDocumentNFTVerifiedNFTMinted)
				if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "VerifiedNFTMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifiedNFTMinted is a log parse operation binding the contract event 0x759831dc39e06177efad640f26e510fec0996d44554975f64801fcfe85e75850.
//
// Solidity: event VerifiedNFTMinted(uint256 indexed tokenId, address indexed owner)
func (_VerifiedDocumentNFT *VerifiedDocumentNFTFilterer) ParseVerifiedNFTMinted(log types.Log) (*VerifiedDocumentNFTVerifiedNFTMinted, error) {
	event := new(VerifiedDocumentNFTVerifiedNFTMinted)
	if err := _VerifiedDocumentNFT.contract.UnpackLog(event, "VerifiedNFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
