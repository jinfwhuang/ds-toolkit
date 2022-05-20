import chai, { expect } from "chai";
import * as dotenv from "dotenv";
import { solidity } from "ethereum-waffle";
import { ethers } from "hardhat";
import { getPubkey } from "../src";
import { UserRegistry } from "../typechain";

chai.use(solidity);

const network = process.env.NETWORK || `hardhat`;
dotenv.config({
  path: `${__dirname}/envs/${network}.conf`,
});
const url = process.env.URL || ``;
const privateKey1 = process.env.PRIVATE_KEY1 || `0x${`11`.repeat(32)}`; // void hardhat error
const privateKey2 = process.env.PRIVATE_KEY2 || `0x${`11`.repeat(32)}`;
const privateKey3 = process.env.PRIVATE_KEY3 || `0x${`11`.repeat(32)}`;

let userRegistry: UserRegistry;

const userName = "testing-user-001";
const keytype = 1; // admin
const keystatus = 1; // active
const pubkey = getPubkey(privateKey1);

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

describe(`UserRegistry`, function () {
  before(async function () {
    userRegistry = await deployContract(); // Singleton for the whole test suite

    // Add a default user
    const [deployer] = await ethers.getSigners();
    const pubkey = getPubkey(privateKey1);
    await userRegistry.newUser(
      deployer.address,
      userName,
      keytype,
      keystatus,
      pubkey
    );
  });

  it(`adding the same user twice`, async function () {
    let wallet = ethers.Wallet.createRandom();
    wallet = wallet.connect(ethers.getDefaultProvider());
    // Add user for first time
    await userRegistry.newUser(
      wallet.address,
      userName,
      keytype,
      keystatus,
      pubkey
    );
    // Add user for the second time
    await expect(
      userRegistry.newUser(wallet.address, userName, keytype, keystatus, pubkey)
    ).to.be.reverted;
  });

  it(`computeAddr`, async function () {
    const addr = await userRegistry.computeAddr(pubkey);
    console.log("computed address", addr);
  });

  it.only(`add more pubkey`, async function () {
    const key = new ethers.utils.SigningKey(privateKey1);
    let wallet = new ethers.Wallet(key.privateKey);
    wallet = wallet.connect(ethers.getDefaultProvider());

    const user = wallet.address;
    const pubkey = key.publicKey;

    // msgToKeccak := encodePacked(addr.Bytes(), []byte{1}, []byte{1}, pubkeyBytes);
    // msgToSign := crypto.Keccak256(msgToKeccak)
    // sig, err := crypto.Sign(msgToSign, privkey)
    const msgToKeccak = ethers.utils.solidityPack(
      ["address", "uint8", "uint8", "bytes"],
      [user, keytype, keystatus, pubkey]
    );
    const msgToSign = ethers.utils.solidityKeccak256(
      ["address", "uint8", "uint8", "bytes"],
      [user, keytype, keystatus, pubkey]
    );
    const sig = key.signDigest(msgToSign);
    const sigFlat = ethers.utils.joinSignature(sig);

    console.log("msgToKeccak=", msgToKeccak);
    console.log("msgToSign=", msgToSign);

    await userRegistry.addPubkey(user, keytype, keystatus, pubkey, sigFlat);
    await userRegistry.addPubkey(user, keytype, keystatus, pubkey, sigFlat);

    // user: string,
    // keytype: BigNumberish,
    // keystatus: BigNumberish,
    // pubkey: BytesLike,
    // sig: BytesLike,
    // overrides?: CallOverrides

    // address user,
    // uint8 keytype,
    // uint8 keystatus,
    // bytes memory pubkey,
    // bytes memory sig

  });
});
