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
    mapping(string => address) lookupUsers; // Name is immutable

    mapping(address => string) lookupNames;

    // Each user has many public address
    mapping(address => Pubkey[]) pubkeys;

    /**
     * Contract initialization.
     */
    constructor() {}

    function rndHash() public view returns (bytes32) {
        return keccak256(abi.encodePacked(block.number));
    }

    function ethSignedHash(bytes32 messageHash) public pure returns (bytes32) {
        return messageHash.toEthSignedMessageHash();
    }

    function recover(bytes32 hash, bytes memory signature)
        public
        pure
        returns (address)
    {
        return hash.recover(signature);
    }

    function getImpliedAddr(address user, uint8 keypos)
        public
        view
        returns (address)
    {
        Pubkey storage pubkey = pubkeys[user][keypos];
        require(uint8(1) == pubkey.keytype, "keytype support not implemented"); // TODO: Expand support beyond 1=secp256k1

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
    function newUser(
        address user,
        string memory name,
        uint8 keytype,
        uint8 keystatus,
        bytes memory pubkey
    ) public {
        require(pubkeys[user].length == 0, "id already exist");

        users.push(user);
        lookupNames[user] = name;
        lookupUsers[name] = user;
        pubkeys[user].push(Pubkey(keytype, keystatus, pubkey));
    }

    /**
      TODO: Allow other kinds of user managements
        1. Update keytype
        2. Update keystatus
        3. Use different keypos to perform verification

      TODO:
        1. How to prevent signature reuse
     */
    function addPubkey(
        address user,
        uint8 keytype,
        uint8 keystatus,
        bytes memory pubkey,
        bytes memory sig
    ) public {
        require(pubkeys[user].length > 0, "id does not exist");

        // Get signature pubkey
        bytes memory msgToKeccak = abi.encodePacked(
            user,
            keytype,
            keystatus,
            pubkey
        );
        bytes32 msgToSign = keccak256(msgToKeccak);
        address signAddr = msgToSign.recover(sig);

        // console.log("----");
        // console.log("msgToKeccak");
        // console.logBytes(msgToKeccak);
        // console.log("msgToSign");
        // console.logBytes32(msgToSign);
        // console.log("signAddr");
        // console.log(signAddr);

        // Get existing pubkey
        Pubkey storage oPubkey = pubkeys[user][0]; // TODO: add support for all keys
        require(uint8(1) == oPubkey.keytype); // TODO: Expand support beyond 1=secp256k1
        require(uint8(1) == oPubkey.keystatus); // Only admin key can modify
        address allowedAddr = computeAddr(oPubkey.key);

        // Proof of private key ownership
        require(signAddr == allowedAddr, "sig is not valid");

        pubkeys[user].push(Pubkey(keytype, keystatus, pubkey));
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

    function getKey(address user, uint8 keypos)
        external
        view
        returns (Pubkey memory)
    {
        return pubkeys[user][keypos];
    }

    function getAllUsers() external view returns (address[] memory) {
        return users;
    }
}
