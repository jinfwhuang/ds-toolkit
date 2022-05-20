import chai from "chai";

import { ethers } from "hardhat";

import { solidity } from "ethereum-waffle";

import { getPubkey } from "../src";

import { expect } from "chai";

import { UserRegistry } from "../typechain";

import { BytesLike, Bytes } from "@ethersproject/bytes";

import { Signer } from "@ethersproject/abstract-signer";

import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";

import { Wallet } from "@ethersproject/wallet";

import * as dotenv from "dotenv";

chai.use(solidity);

const network = process.env.NETWORK || `hardhat`;
dotenv.config({
  path: `${__dirname}/envs/${network}.conf`,
});
const url = process.env.URL || ``;
const privateKey1 = process.env.PRIVATE_KEY1 || `0x${`11`.repeat(32)}`; // void hardhat error
const privateKey2 = process.env.PRIVATE_KEY2 || `0x${`11`.repeat(32)}`;
const privateKey3 = process.env.PRIVATE_KEY3 || `0x${`11`.repeat(32)}`;

describe(`Examples of key operations`, function () {
  it(`get signer from private key`, async function () {
    let privateKey =
      "0x0123456789012345678901234567890123456789012345678901234567890123";
    let signer = new ethers.Wallet(privateKey);

    // Connect a wallet to mainnet
    let provider = ethers.getDefaultProvider();
    let signerConnected = new ethers.Wallet(privateKey, provider);

    console.log("privateKey", privateKey);
    console.log("address", signer.address);
  });

  it(`using signer`, async function () {
    const [signer1] = await ethers.getSigners();
    console.log("address", signer1.address);
  });

  it(`pubkey and address`, async function () {
    const pubkey = getPubkey(privateKey1);
    console.log("pubkey", pubkey);
  });
});
