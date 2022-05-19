// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.8;

// dev only
import "hardhat/console.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "solidity-bytes-utils/contracts/BytesLib.sol";

struct Pubkey {
    uint8 keytype; // 1=secp256k1, 2=bls-12-381
    uint8 keystatus; // 1=admin, 2=active, 3=canceled
    bytes key;
}

contract UserRegistry {

    using ECDSA for bytes32;
    using BytesLib for bytes;

    // Each user is uniquely identified by an ID, which is a 20 bytes integer
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

        // // Debugging only
        // // https://github.com/NomicFoundation/hardhat/blob/master/packages/hardhat-core/console.sol
        // console.log("concat message");
        // console.logBytes(concatMsg);
        // console.log("hashToSign");
        // console.logBytes32(hashToSign);
        // console.log("recovering address");
        // console.logAddress(hashToSign.recover(signature));

        return hashToSign.recover(signature) == user;
    }

    function getImpliedAddr(address user, uint8 keypos) public view returns (address) {
        Pubkey storage pubkey = pubkeys[user][keypos];
        assert(uint8(1) == pubkey.keytype); // TODO: Expand support beyond 1=secp256k1

        return computeAddr(pubkey.key);
    }

    function computeAddr(bytes memory pubkey) public pure returns (address) {
        // The first byte indicates that it is an uncompressed point
        // See: section 4.3.6 of ANSI X9.62.
        bytes32 _hash = keccak256(pubkey.slice(1, 64));

        // The last 20 bytes of the keccak256 hash
        bytes memory addr = abi.encodePacked(_hash).slice(12, 20);

        return address(bytes20(addr));
    }


    // Anyone could create a user without verification.
    // Having access to the privatekey that matches "pubkey" is equivalent to owning the user identity.
    function newUser(address user, string memory name, uint8 keytype, uint8 keykeystatus, bytes memory pubkey) public {
        require(pubkeys[user].length == 0, "id already exist");

        users.push(user);
        lookupNames[user] = name;
        lookupUsers[name] = user;
        userNonce[user] = 0;
        pubkeys[user].push(Pubkey(keytype, keykeystatus, pubkey));

        console.log(user);
        console.log(name);
        console.log(keytype);
        console.log(keykeystatus);
        console.logBytes(pubkey);
    }

    /**
      TODO: Allow other kinds of user managements
        1. Update keytype
        2. Update keystatus
        3. Use different keypos to perform verification
     */
    function addPubkey(address user, uint8 keytype, uint8 keystatus, bytes memory pubkey, bytes memory sig) public {
        console.log(user);
        console.log(pubkeys[user].length);
        require(pubkeys[user].length > 0, "id does not exist");

        // Get signature pubkey
        bytes memory msgToKeccak = abi.encodePacked(user, keytype, keystatus, pubkey);
        bytes32 msgToSign = keccak256(msgToKeccak);
        address signAddr = msgToSign.recover(sig);

        console.logBytes(msgToKeccak);
        console.logBytes32(msgToSign);
        console.log(signAddr);

        // Get existing pubkey
        Pubkey storage oPubkey = pubkeys[user][0]; // TODO: add support for all keys
        require(uint8(1) == oPubkey.keytype); // TODO: Expand support beyond 1=secp256k1
        require(uint8(1) == oPubkey.keystatus); // Only admin key can modify
        address allowedAddr = computeAddr(oPubkey.key);

        // Proof of private key ownership
        require(signAddr == allowedAddr, "sig is not valid");
        console.log("verify signature");

        pubkeys[user].push(Pubkey(keytype, keystatus, pubkey));
    }

    // function updateKeystatus(address user, uint8 keypos, uint8 keystatus) public {
    // }

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
