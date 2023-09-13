// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const LegacyERC20ETHStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"src/legacy/LegacyERC20ETH.sol:LegacyERC20ETH\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var LegacyERC20ETHStorageLayout = new(solc.StorageLayout)

var LegacyERC20ETHDeployedBin = "0x608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610353578063e78cea9214610307578063ee9a31a21461039957600080fd5b8063ae1f6aaf14610307578063c01e1bd61461032d578063d6c0b2c41461032d57600080fd5b80639dc29fac116100bd5780639dc29fac146102ce578063a457c2d7146102e1578063a9059cbb146102f457600080fd5b806370a082311461029e57806395d89b41146102c657600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461026e57806340c10f191461028157806354fd4d501461029657600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004610ab1565b6103c0565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104b1565b60405161019b9190610b2a565b61018f610213366004610ba4565b610543565b6002545b60405190815260200161019b565b61018f610238366004610bce565b6105d3565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019b565b61018f61027c366004610ba4565b61065e565b61029461028f366004610ba4565b6106e9565b005b6101f861074b565b61021c6102ac366004610c0a565b73ffffffffffffffffffffffffffffffffffffffff163190565b6101f86107ee565b6102946102dc366004610ba4565b6107fd565b61018f6102ef366004610ba4565b61085f565b61018f610302366004610ba4565b6108ea565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c610361366004610c25565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061047957507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104a857507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104c090610c58565b80601f01602080910402602001604051908101604052809291908181526020018280546104ec90610c58565b80156105395780601f1061050e57610100808354040283529160200191610539565b820191906000526020600020905b81548152906001019060200180831161051c57829003601f168201915b5050505050905090565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f4c656761637945524332304554483a20617070726f766520697320646973616260448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526000906084015b60405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f4c656761637945524332304554483a207472616e7366657246726f6d2069732060448201527f64697361626c656400000000000000000000000000000000000000000000000060648201526000906084016105ca565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4c656761637945524332304554483a20696e637265617365416c6c6f77616e6360448201527f652069732064697361626c65640000000000000000000000000000000000000060648201526000906084016105ca565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4c656761637945524332304554483a206d696e742069732064697361626c656460448201526064016105ca565b60606107767f0000000000000000000000000000000000000000000000000000000000000000610974565b61079f7f0000000000000000000000000000000000000000000000000000000000000000610974565b6107c87f0000000000000000000000000000000000000000000000000000000000000000610974565b6040516020016107da93929190610cab565b604051602081830303815290604052905090565b6060600480546104c090610c58565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4c656761637945524332304554483a206275726e2069732064697361626c656460448201526064016105ca565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4c656761637945524332304554483a206465637265617365416c6c6f77616e6360448201527f652069732064697361626c65640000000000000000000000000000000000000060648201526000906084016105ca565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4c656761637945524332304554483a207472616e73666572206973206469736160448201527f626c65640000000000000000000000000000000000000000000000000000000060648201526000906084016105ca565b6060816000036109b757505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156109e157806109cb81610d50565b91506109da9050600a83610db7565b91506109bb565b60008167ffffffffffffffff8111156109fc576109fc610dcb565b6040519080825280601f01601f191660200182016040528015610a26576020820181803683370190505b5090505b8415610aa957610a3b600183610dfa565b9150610a48600a86610e11565b610a53906030610e25565b60f81b818381518110610a6857610a68610e3d565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610aa2600a86610db7565b9450610a2a565b949350505050565b600060208284031215610ac357600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610af357600080fd5b9392505050565b60005b83811015610b15578181015183820152602001610afd565b83811115610b24576000848401525b50505050565b6020815260008251806020840152610b49816040850160208701610afa565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b9f57600080fd5b919050565b60008060408385031215610bb757600080fd5b610bc083610b7b565b946020939093013593505050565b600080600060608486031215610be357600080fd5b610bec84610b7b565b9250610bfa60208501610b7b565b9150604084013590509250925092565b600060208284031215610c1c57600080fd5b610af382610b7b565b60008060408385031215610c3857600080fd5b610c4183610b7b565b9150610c4f60208401610b7b565b90509250929050565b600181811c90821680610c6c57607f821691505b602082108103610ca5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60008451610cbd818460208901610afa565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610cf9816001850160208a01610afa565b60019201918201528351610d14816002840160208801610afa565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d8157610d81610d21565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610dc657610dc6610d88565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610e0c57610e0c610d21565b500390565b600082610e2057610e20610d88565b500690565b60008219821115610e3857610e38610d21565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(LegacyERC20ETHStorageLayoutJSON), LegacyERC20ETHStorageLayout); err != nil {
		panic(err)
	}

	layouts["LegacyERC20ETH"] = LegacyERC20ETHStorageLayout
	deployedBytecodes["LegacyERC20ETH"] = LegacyERC20ETHDeployedBin
}
