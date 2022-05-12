//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.8;

type KeyStatus is uint8; // 1=admin, 2=active, 3=canceled

type KeyType is uint8; // 1=secp256k1 2=bls-12-381

import "hardhat/console.sol";

/*

---
# Immutable
# Matches the first entry in pubkeys: the last 20 bytes of the Keccak-256 hash of public key
id: '0x48B7872500c5BefDBb7BBe9dB9070CeEC66bdD4b'

# Immutable, optional
# lower case alphanumeric, [.-_] # immutable, optional
name: 'jin.digitalgreen'

# pubkey:
#  - secp256k1 public key
#  - compressed
#  - hex
#  - prefixed with 0x

# Entries are append-only
# "status" field could be modified
# All 'admin' pubkeys has permission to perform update
publicKeys:
  - type: secp256k1
    key: '0x0233085d6674c5629673e9d0fb01ff9b41c3b2accb1cdb7e94516b0a29c0b399b5'
    status: 'admin' # admin, active, canceled
  - type: secp256k1
    key: '0x02a5ca0d04213b2c044dc636887c0eabf904c31c679f46ad30110f05eb7f093e95'
    status: 'active'
  - type: bls-12-381
    key: 0xxxxxxxxx
    status: 'active'

metadata_url: ipfs/link/containing-public-information


*/


struct Pubkey {
    KeyType keytype;
    KeyStatus status; // 1=admin, 2=active, 3=canceled
    bytes key;
}

contract UserRegistry {


    // string public name = "Dstoolkit-Testing-Token";
    // string public symbol = "DST";

    // // The fixed amount of tokens stored in an unsigned integer type variable.
    // uint256 public totalSupply = 4529;

    // // An address type variable is used to store ethereum accounts.
    // address public owner;

    // // A mapping is a key/value map. Here we store each account balance.
    // mapping(address => uint256) balances;

    // Each user is uniquely identified by an ID, which is a 20 bytes integer
    mapping(address => string) usernames;

    mapping(string => address) users;

    // Each user has many public address
    mapping(address => Pubkey[]) pubkeys;

    // Each user has many public address
    mapping(address => uint8[]) pubkeyType;
    mapping(address => bytes[]) pubkeyByte;

    // // Each user has many address that it could receive funds
    // mapping(address => address[]) addresses;

    /**
     * Contract initialization.
     *
     * The `constructor` is executed only once when the contract is created.
     */
    constructor() {
        // // The totalSupply is assigned to transaction sender, which is the account
        // // that is deploying the contract.
        // balances[msg.sender] = totalSupply;
        // owner = msg.sender;
    }

    /**
     */
    function addPubkey(address id, string memory name, KeyType keytype, KeyStatus status, bytes memory key) public {
        // TODO: require a signed message
        usernames[id] = name;
        users[name] = id;
        // pubkeyByte[id].push(key);
        
        pubkeys[id].push(Pubkey(keytype, status, key));
    }

    function updateKeyStatus(address id, uint8 keypos, KeyStatus status) public {
        // TODO: require proof of private key ownership
        pubkeys[id][keypos].status = status;
    }

    function getKey(address id, uint8 keypos) external view returns (Pubkey memory) {
        // TODO: require proof of private key ownership
        return pubkeys[id][keypos];
    }

    // /**
    //  * Read only function to retrieve the token balance of a given account.
    //  *
    //  * The `view` modifier indicates that it doesn't modify the contract's
    //  * state, which allows us to call it without executing a transaction.
    //  */
    // function balanceOf(address account) external view returns (uint256) {
    //     console.log("address", account);
    //     console.log("address2", account);
    //     console.log("balance", balances[account]);
    //     return balances[account];
    // }
}
