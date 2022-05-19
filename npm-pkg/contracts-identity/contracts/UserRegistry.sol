// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.8;


// dev only
import "hardhat/console.sol";

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
// import "solidity-bytes-utils/BytesLib.sol";
import "solidity-bytes-utils/contracts/BytesLib.sol";


// using ECDSA for bytes32;



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
    uint8 keytype; // 1=secp256k1, 2=bls-12-381
    uint8 status; // 1=admin, 2=active, 3=canceled
    bytes key;
}

contract UserRegistry {

    using ECDSA for bytes32;
    using BytesLib for bytes;

    // Each user is uniquely identified by an ID, which is a 20 bytes integer
    // - 
    address[] users; // Append only; address is immutale
    mapping(string => address) lookupUsers;  // Name is immutable

    mapping(address => string) lookupNames;
    mapping(address => uint16) userNonce;

    // Each user has many public address
    mapping(address => Pubkey[]) pubkeys;

    /**
     * Contract initialization.
     */
    constructor() {
    }

    
    function getUserNonce(address user) public view returns (uint16) {
        return userNonce[user];
    }

    function updateUserNonce(address user) private {
        userNonce[user] += 1;
    }

    function rndHash() public view returns(bytes32) {
        return keccak256(abi.encodePacked(block.number));
    }

    function ethSignedHash(bytes32 messageHash) public pure returns(bytes32) {
        return messageHash.toEthSignedMessageHash();
    }

    function recover(bytes32 hash, bytes memory signature) public pure returns(address) {
        return hash.recover(signature);
    }

    function verifyUser(address user, bytes32 msgHash, bytes memory signature) public view returns (bool isValid) {
        uint16 nonce = userNonce[user];
        bytes memory concatMsg = abi.encodePacked(nonce, msgHash);
        bytes32 hashToSign = keccak256(concatMsg);

        // Debugging only
        // https://github.com/NomicFoundation/hardhat/blob/master/packages/hardhat-core/console.sol
        console.log("concat message");
        console.logBytes(concatMsg);
        console.log("hashToSign");
        console.logBytes32(hashToSign);
        console.log("recovering address");
        console.logAddress(hashToSign.recover(signature));

        return hashToSign.recover(signature) == user;
    }

    function getImpliedAddr(address user, uint8 keypos) public view returns (address) {
        Pubkey storage pubkey = pubkeys[user][keypos];
        assert(uint8(1) == uint8(1)); // TODO: Expand support beyond 1=secp256k1

        return computeAddr(pubkey.key);
    }

    function computeAddr(bytes memory key) public pure returns (address) {
        // The first byte indicates that it is an uncompressed point
        // See: section 4.3.6 of ANSI X9.62.
        bytes32 _hash = keccak256(key.slice(1, 64));
        bytes memory addr = abi.encodePacked(_hash).slice(12, 20);
        address a = address(bytes20(addr));
        return a;
    }


    function newUser(address user, string memory name, uint8 keytype, uint8 keystatus, bytes memory key) public {
        // TODO: require proof of private key ownership
        require(pubkeys[user].length == 0, "id already exist");
        // if (pubkeys[user].length >= 1) {
        //     console.log("id already exist");
        //     // return "id alreay exist";
        //     throw;
        // }

        users.push(user);
        lookupNames[user] = name;
        lookupUsers[name] = user;
        userNonce[user] = 0;
        pubkeys[user].push(Pubkey(keytype, keystatus, key));

        console.log("user");
        console.log(name);
        console.log(keytype);
        console.log(keystatus);
        console.logBytes(key);
    }



    /**
     */
    function addPubkey(address user, uint8 keytype, uint8 status, bytes memory key) public returns (string memory) {
        // TODO: require proof of private key ownership
        if (pubkeys[user].length < 1) {
            return "id does not exist";
        }
        
        pubkeys[user].push(Pubkey(keytype, status, key));
        return "added new key";
    }

    function updateKeyStatus(address user, uint8 keypos, uint8 status) public {
        // TODO: require proof of private key ownership
        pubkeys[user][keypos].status = status;
    }

    function getName(address user) external view returns (string memory) {
        return lookupNames[user];
    }

    function getUser(string memory name) external view returns (address) {
        return lookupUsers[name];
    }

    function getLenKeys(address user) external view returns (uint256) {
        return pubkeys[user].length;
    }

    function getKeys(address user) external view returns (Pubkey[] memory) {
        return pubkeys[user];
    }

    function getKey(address user, uint8 keypos) external view returns (Pubkey memory) {
        return pubkeys[user][keypos];
    }

    function getAllUsers() external view returns (address[] memory) {
        return users;
    }

}
