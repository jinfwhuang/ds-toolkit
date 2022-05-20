import { Wallet } from "@ethersproject/wallet";
import { ethers } from "hardhat";

import * as dotenv from "dotenv";
dotenv.config({
  path: `${__dirname}/envs/localhost.conf`,
});

console.log("-----ENV-----");
console.log("pwd:", __dirname);
console.log("network:", process.env.NETWORK);
console.log("url", process.env.URL);
console.log("--------------");

const url = process.env.URL || ``;
const privateKey1 = process.env.PRIVATE_KEY1 || `0x${`11`.repeat(32)}`; // void hardhat error
const privateKey2 = process.env.PRIVATE_KEY2 || `0x${`11`.repeat(32)}`;
const privateKey3 = process.env.PRIVATE_KEY3 || `0x${`11`.repeat(32)}`;


let acc001 = new Wallet(privateKey1);
acc001 = acc001.connect(ethers.provider);

// const balanceInEth = ethers.utils.formatEther(balancae); // wei to ether
console.log(`wei: ${ethers.utils.parseEther("1")}`);
console.log(`eth: ${ethers.utils.formatEther("1"+"0".repeat(18))}`); // 10^18
console.log(`eth: ${ethers.utils.formatEther(ethers.utils.parseEther("1"))}`);

export async function setup() {
  const [accWithMoney, ] = await ethers.getSigners();
  console.log(`Default addr:`, accWithMoney.address);
  console.log(`Balance     :`, (await accWithMoney.getBalance()).toString());

  const resp = await accWithMoney.sendTransaction({
    to: acc001.address,
    value: ethers.utils.parseEther("999"),
  });

  await ethers.provider.getBalance(acc001.address).then((balance) => {
    console.log("deployer addr", resp.to);
    const balanceInEth = ethers.utils.formatEther(balance); // wei to ether
    console.log(`balance: ${balance} wei`);
    console.log(`balance: ${balanceInEth} ETH`);
  });
}

// Run this once for a new hardhat local network
setup()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
