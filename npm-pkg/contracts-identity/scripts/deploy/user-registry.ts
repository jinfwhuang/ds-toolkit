// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { ethers } from "hardhat";
import { UserRegistry } from "typechain/UserRegistry";

const pubkeyHex =
//0x045bb46d799b99b66be40533426d8ec34f3b53f61195ec85cd9443d45551b51aecf0ac19d39b107c97edf66d91bf4a57ed99838a18c943d253664baa0012d9a145
 "0x040c1ca15b1ee87e5c493b85d4f2db6b13bc3aadb61f7af5b84ad30451074ad500b95b745a6600326d91bd4323da514b4b81d5d76f0973b66d6cf8e3b131525d41"

async function deployContract() {
  const [deployer] = await ethers.getSigners();
  console.log(`Deploying account:`, deployer.address);
  console.log(`Account balance:`, (await deployer.getBalance()).toString());

  const contractName = "UserRegistry";
  const contractFactory = await ethers.getContractFactory(contractName);
  const contract = await contractFactory.deploy();

  console.log(`Contract address:`, contract.address);

  const pubkeyStr = pubkeyHex;
  await contract.newUser(deployer.address, "jinhuang001", 1, 1, pubkeyStr);

//   return contract;
}

deployContract()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
