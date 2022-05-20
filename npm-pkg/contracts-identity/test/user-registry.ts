import chai from "chai";

import { expect } from "chai";

import { ethers } from "hardhat";

import { solidity } from "ethereum-waffle";

import { UserRegistry } from "../typechain";

import { BytesLike, Bytes } from "@ethersproject/bytes";

import * as dotenv from "dotenv";

chai.use(solidity);

// const network = process.env.NETWORK || `localhost`;

// dotenv.config({
//   path: `${__dirname}/envs/${network}.conf`,
// });

// console.log("-----ENV-----");
// console.log("pwd:", __dirname);
// console.log("network:", process.env.NETWORK);
// console.log("url", process.env.URL);
// console.log("--------------");

// const url = process.env.URL || ``;
// const privateKey = process.env.PRIVATE_KEY || `0x${`11`.repeat(32)}`; // Avoid hardhat error

// console.log("network url", url);
// console.log("privateKey", privateKey);

// let wallet = new ethers.Wallet(privateKey);
// let provider = ethers.getDefaultProvider();
// let walletWithProvider = new ethers.Wallet(privateKey, provider);

let userRegistry: UserRegistry;
const pubkeyHex =
  "0x045bb46d799b99b66be40533426d8ec34f3b53f61195ec85cd9443d45551b51aecf0ac19d39b107c97edf66d91bf4a57ed99838a18c943d253664baa0012d9a145";

const userName = "testing-user-001";
const keytype = 1; // admin
const keystatus = 1; // active

async function deployContract(): Promise<UserRegistry> {
  const [signer1] = await ethers.getSigners();
  console.log(`Deploying account:`, signer1.address);
  console.log(`Account balance:`, (await signer1.getBalance()).toString());

  const contractName = "UserRegistry";
  const contractFactory = await ethers.getContractFactory(contractName);
  const contract = await contractFactory.deploy();

  console.log(`Contract address:`, contract.address);

  return contract;
}

async function getPubkey(): Promise<string> {
  const userRegistry = await deployContract();
  const [deployer] = await ethers.getSigners();
  const deployTx = userRegistry.deployTransaction;
  const msg = ethers.utils.RLP.encode(deployTx.data);
  const msgHash = ethers.utils.keccak256(msg);
  const msgBytes = ethers.utils.arrayify(msgHash);
  const expanded = {
    r: deployTx.r,
    s: deployTx.s,
    recoveryParam: 0,
    v: deployTx.v,
  };
  const signature = ethers.utils.joinSignature(expanded);
  const recoveredPubKey = ethers.utils.recoverPublicKey(msgBytes, signature);
  const recoveredAddress = ethers.utils.recoverAddress(msgBytes, signature);

  console.log("deployer addr", deployer.address);
  console.log("hardcoded pubkey", pubkeyHex);
  console.log("recoveredPubKey", recoveredPubKey);
  console.log("recovered addr", recoveredAddress);

  return recoveredPubKey;
}

// TODO: Add more tests

describe(`UserRegistry`, function () {
  before(async function () {
    userRegistry = await deployContract();

    // Add default user
    const [owner] = await ethers.getSigners();
    const pubkeyStr = pubkeyHex;
    await userRegistry.newUser(owner.address, "jinhuang001", 1, 1, pubkeyStr);
  });

  it(`get pubkey`, async function () {
    const pubkey = getPubkey();
  });

  it(`adding the same user twice`, async function () {
    const [signer1, signer2, signer3] = await ethers.getSigners();
    const pubkeyStr = pubkeyHex;
    // Add user for first time
    await userRegistry.newUser(signer3.address, userName, 1, 1, pubkeyStr);
    await expect(
      userRegistry.newUser(signer3.address, userName, 1, 1, pubkeyStr)
    ).to.be.reverted;
  });

  it(`computeAddr`, async function () {
    const addr = await userRegistry.computeAddr(pubkeyHex);
    console.log("computed address", addr);
  });
});
