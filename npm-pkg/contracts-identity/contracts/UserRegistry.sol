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
    KeyType keytype; // 1=secp256k1, 2=bls-12-381
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

    address[] users;

    // Each user is uniquely identified by an ID, which is a 20 bytes integer
    mapping(address => string) lookupNames;
    mapping(string => address) lookupUsers;

    // Each user has many public address
    mapping(address => Pubkey[]) pubkeys;

    // // Each user has many public address
    // mapping(address => uint8[]) pubkeyType;
    // mapping(address => bytes[]) pubkeyByte;


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

    function newUser(address user, string memory name, KeyType keytype, KeyStatus status, bytes memory key) public returns (string memory) {
        // TODO: require a signed message
        if (pubkeys[user].length >= 1) {
            return "id alreay exist";
        }

        users.push(user);
        lookupNames[user] = name;
        lookupUsers[name] = user;
        pubkeys[user].push(Pubkey(keytype, status, key));
        return "created";
    }

    /**
     */
    function addPubkey(address user, KeyType keytype, KeyStatus status, bytes memory key) public returns (string memory) {
        // TODO: require a signed message
        if (pubkeys[user].length < 1) {
            return "id does not exist";
        }
        
        pubkeys[user].push(Pubkey(keytype, status, key));
        return "added new key";
    }

    function updateKeyStatus(address user, uint8 keypos, KeyStatus status) public {
        // TODO: require proof of private key ownership
        pubkeys[user][keypos].status = status;
    }

    function getName(address user) external view returns (string memory) {
        return lookupNames[user];
    }

    function getUser(string memory name) external view returns (address) {
        return lookupUsers[name];
    }

    function getKeys(address user) external view returns (Pubkey[] memory) {
        return pubkeys[user];
    }

    function getKey(address user, uint8 keypos) external view returns (Pubkey memory) {
        return pubkeys[user][keypos];
    }

    function getKeyLen(address user) external view returns (uint256) {
        return pubkeys[user].length;
    }

    function getAllUsers() external view returns (address[] memory) {
        return users;
    }

    // function getAllIds() external view returns (string[] memory) {
    //     names 
    //     usernames
    //     return usernames;
    // }

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
