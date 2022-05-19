// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { ethers } from "hardhat";

export async function deployContract() {
  const [deployer] = await ethers.getSigners();
  console.log(`Deploying contracts with the account:`, deployer.address);
  console.log(`Account balance:`, (await deployer.getBalance()).toString());

  const contractName = "UserRegistry";
  const Token = await ethers.getContractFactory(contractName);
  const token = await Token.deploy();

  console.log(`Token address:`, token.address);
}

deployContract()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
