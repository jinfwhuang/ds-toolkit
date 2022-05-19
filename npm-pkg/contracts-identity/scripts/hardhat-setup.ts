import { Wallet } from "@ethersproject/wallet";
import { ethers } from "hardhat";

const hardhatPrivkeyHex =
  "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";

const deployerAddr = "0xaABcEa31ac2c76B5d11ad579d26A671D4F20171B";

export async function setup() {
  let wallet = new Wallet(hardhatPrivkeyHex);
  wallet = wallet.connect(ethers.provider);

  console.log(`Default addr:`, wallet.address);
  console.log(`Account balance:`, (await wallet.getBalance()).toString());

  const resp = await wallet.sendTransaction({
    to: deployerAddr,
    value: ethers.utils.parseEther("999"),
  });

  await ethers.provider.getBalance(deployerAddr).then((balance) => {
    // convert a currency unit from wei to ether
    console.log("deployer addr", resp.to);
    const balanceInEth = ethers.utils.formatEther(balance);
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
