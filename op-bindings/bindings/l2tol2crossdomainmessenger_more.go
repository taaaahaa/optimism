// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2ToL2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2ToL2CrossDomainMessenger.sol:L2ToL2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1001,\"contract\":\"src/L2/L2ToL2CrossDomainMessenger.sol:L2ToL2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint240\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"}}}"

var L2ToL2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2ToL2CrossDomainMessengerDeployedBin = "0x6080604052600436106100ad575f3560e01c80637056f41f11610066578063b1b1b2091161004c578063b1b1b2091461021a578063b1f35f2c14610258578063ecc704281461028b575f80fd5b80637056f41f146101d45780638fe5a280146101e7575f80fd5b806338ffde181161009657806338ffde18146101085780633f827a5a1461015957806354fd4d501461017f575f80fd5b80631ecd26f2146100b157806324794462146100c6575b5f80fd5b6100c46100bf366004610a3a565b6102bf565b005b3480156100d1575f80fd5b507f711dfa3259c842fffc17d6e1f1e0fc5927756133a2345ca56b4cb8178589fee75c5b6040519081526020015b60405180910390f35b348015610113575f80fd5b5060405173ffffffffffffffffffffffffffffffffffffffff7fb83444d07072b122e2e72a669ce32857d892345c19856f4e7142d06a167ab3f35c1681526020016100ff565b348015610164575f80fd5b5061016c5f81565b60405161ffff90911681526020016100ff565b34801561018a575f80fd5b506101c76040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100ff9190610ba3565b6100c46101e2366004610bbc565b6107fe565b3480156101f2575f80fd5b506100f57f711dfa3259c842fffc17d6e1f1e0fc5927756133a2345ca56b4cb8178589fee781565b348015610225575f80fd5b50610248610234366004610c3e565b5f6020819052908152604090205460ff1681565b60405190151581526020016100ff565b348015610263575f80fd5b506100f57fb83444d07072b122e2e72a669ce32857d892345c19856f4e7142d06a167ab3f381565b348015610296575f80fd5b506001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff166100f5565b3373420000000000000000000000000000000000002214610367576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f4c32546f4c3243726f7373446f6d61696e4d657373656e6765723a2073656e6460448201527f6572206e6f742043726f73734c32496e626f780000000000000000000000000060648201526084015b60405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff1673420000000000000000000000000000000000002273ffffffffffffffffffffffffffffffffffffffff1663938b5f326040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103db573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103ff9190610c55565b73ffffffffffffffffffffffffffffffffffffffff16146104c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f4c32546f4c3243726f7373446f6d61696e4d657373656e6765723a2043726f7360448201527f734c32496e626f78206f726967696e206e6f74207468697320636f6e7472616360648201527f7400000000000000000000000000000000000000000000000000000000000000608482015260a40161035e565b468614610557576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c32546f4c3243726f7373446f6d61696e4d657373656e6765723a206465737460448201527f696e6174696f6e206e6f74207468697320636861696e00000000000000000000606482015260840161035e565b7fffffffffffffffffffffffffbdffffffffffffffffffffffffffffffffffffde73ffffffffffffffffffffffffffffffffffffffff83160161061c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603b60248201527f4c32546f4c3243726f7373446f6d61696e4d657373656e6765723a2043726f7360448201527f734c32496e626f782063616e6e6f742063616c6c20697473656c660000000000606482015260840161035e565b5f86868686868660405160200161063896959493929190610c70565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f8181529283905291205490915060ff161561070c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f4c32546f4c3243726f7373446f6d61696e4d657373656e6765723a206d65737360448201527f61676520616c72656164792072656c6179656400000000000000000000000000606482015260840161035e565b5f867f711dfa3259c842fffc17d6e1f1e0fc5927756133a2345ca56b4cb8178589fee75d847fb83444d07072b122e2e72a669ce32857d892345c19856f4e7142d06a167ab3f35d5f8084516020860134885af1905080156107c9575f8281526020819052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26107f4565b60405182907f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f905f90a25b5050505050505050565b46840361088d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c32546f4c3243726f7373446f6d61696e4d657373656e6765723a2063616e6e60448201527f6f742073656e64206d65737361676520746f2073656c66000000000000000000606482015260840161035e565b5f84466108b96001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b338787876040516024016108d39796959493929190610cc6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1ecd26f20000000000000000000000000000000000000000000000000000000017905251909150610958908290610ba3565b60405180910390a0600180547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905f61098f83610d51565b91906101000a8154817dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff160217905550505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610a0a575f80fd5b50565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f805f805f8060c08789031215610a4f575f80fd5b8635955060208701359450604087013593506060870135610a6f816109e9565b92506080870135610a7f816109e9565b915060a087013567ffffffffffffffff80821115610a9b575f80fd5b818901915089601f830112610aae575f80fd5b813581811115610ac057610ac0610a0d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610b0657610b06610a0d565b816040528281528c6020848701011115610b1e575f80fd5b826020860160208301375f6020848301015280955050505050509295509295509295565b5f81518084525f5b81811015610b6657602081850181015186830182015201610b4a565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610bb56020830184610b42565b9392505050565b5f805f8060608587031215610bcf575f80fd5b843593506020850135610be1816109e9565b9250604085013567ffffffffffffffff80821115610bfd575f80fd5b818701915087601f830112610c10575f80fd5b813581811115610c1e575f80fd5b886020828501011115610c2f575f80fd5b95989497505060200194505050565b5f60208284031215610c4e575f80fd5b5035919050565b5f60208284031215610c65575f80fd5b8151610bb5816109e9565b8681528560208201528460408201525f73ffffffffffffffffffffffffffffffffffffffff808616606084015280851660808401525060c060a0830152610cba60c0830184610b42565b98975050505050505050565b8781528660208201528560408201525f73ffffffffffffffffffffffffffffffffffffffff808716606084015280861660808401525060c060a08301528260c0830152828460e08401375f60e0848401015260e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f850116830101905098975050505050505050565b5f7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808316818103610da8577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b600101939250505056fea164736f6c6343000818000a"


func init() {
	if err := json.Unmarshal([]byte(L2ToL2CrossDomainMessengerStorageLayoutJSON), L2ToL2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ToL2CrossDomainMessenger"] = L2ToL2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2ToL2CrossDomainMessenger"] = L2ToL2CrossDomainMessengerDeployedBin
	immutableReferences["L2ToL2CrossDomainMessenger"] = false
}
