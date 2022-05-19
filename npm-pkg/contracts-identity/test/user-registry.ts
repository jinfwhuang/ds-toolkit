import chai from "chai";

import { expect } from "chai";

import { ethers } from "hardhat";

import { solidity } from "ethereum-waffle";

import { UserRegistry } from "../typechain";

import { BytesLike, Bytes } from "@ethersproject/bytes";
import { AbiCoder } from "@ethersproject/abi";
import { abort } from "process";

chai.use(solidity);

// describe(`Greeter`, function () {
//   it(`Should return the new greeting once it's changed`, async function () {
//     const Greeter = await ethers.getContractFactory(`Greeter`);
//     const greeter = await Greeter.deploy(`Hello, world!`);
//     await greeter.deployed();

//     expect(await greeter.greet()).to.equal(`Hello, world!`);

//     const setGreetingTx = await greeter.setGreeting(`Hola, mundo!`);

//     // wait until the transaction is mined
//     await setGreetingTx.wait();

//     expect(await greeter.greet()).to.equal(`Hola, mundo!`);
//   });
// });

// describe(`User Registry`, function () {

//     it(`Deployment should assign the total supply of tokens to the owner`, async function () {
//       const Token = await ethers.getContractFactory(`Token`);
//       const hardhatToken = await Token.deploy();

//       const [owner] = await ethers.getSigners();
//       const ownerBalance = await hardhatToken.balanceOf(owner.address);
//       const tokenTotalSupply = await hardhatToken.totalSupply();

//       expect(tokenTotalSupply).to.equal(ownerBalance);
//     });
//   });

let userRegistry: UserRegistry;
const pubkeyHex = "0x045bb46d799b99b66be40533426d8ec34f3b53f61195ec85cd9443d45551b51aecf0ac19d39b107c97edf66d91bf4a57ed99838a18c943d253664baa0012d9a145";


async function deployContract(): Promise<UserRegistry> {
  const [deployer] = await ethers.getSigners();
  console.log(`Deploying account:`, deployer.address);
  console.log(`Account balance:`, (await deployer.getBalance()).toString());

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
      v: deployTx.v
    };
    const signature = ethers.utils.joinSignature(expanded);
    const recoveredPubKey = ethers.utils.recoverPublicKey(
      msgBytes,
      signature
    );
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
    const pubkeyStr = pubkeyHex
    await userRegistry.newUser(
      owner.address,
      "jinhuang001",
      1,
      1,
      pubkeyStr
    );    
  });

  it(`get pubkey`, async function () {
    const pubkey = getPubkey();
  });

//   it(`adding the same user twice`, async function () {
//     const [deployer] = await ethers.getSigners();

//     const pubkeyStr = pubkeyHex
//     const tx01 = await userRegistry.newUser(
//       deployer.address,
//       "jinhuang001",
//       1,
//       1,
//       pubkeyStr
//     );
//     await expect(
//       userRegistry.newUser(deployer.address, "jinhuang001", 1, 1, pubkeyStr)
//     ).to.be.reverted;
//   });


  it(`computeAddr`, async function () {
    const addr = await userRegistry.computeAddr(
      pubkeyHex
    );
    console.log("computed address", addr)
  });


});
