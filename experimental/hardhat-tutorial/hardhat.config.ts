// import * as dotenv from "dotenv";
// dotenv.config();

import { HardhatUserConfig, task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";
import "hardhat-gas-reporter";
import "solidity-coverage";

const network = process.env.NETWORK || "ropsten";

require("dotenv").config({
  path: __dirname + "/envs/" + network + ".env",
});
console.log("-----ENV-----");
console.log("pwd", __dirname);
console.log("network", process.env.NETWORK);
console.log("url", process.env.URL);
console.log("--------------");

const enableGasReport = !!process.env.ENABLE_GAS_REPORT;
const url = process.env.URL || "";
const privateKey = process.env.PRIVATE_KEY || "0x" + "11".repeat(32); // this is to avoid hardhat error
const deploy = process.env.DEPLOY_DIRECTORY || "deploy";

const config: HardhatUserConfig = {
  solidity: "0.8.4",
  defaultNetwork: network,
  networks: {
    mainnet: {
      chainId: 1,
      url: url,
      accounts: [privateKey],
    },
    ropsten: {
      url: url,
      accounts: [privateKey],
    },
    rinkeby: {
      url: url,
      accounts: [privateKey],
    },
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
};

// https://hardhat.org/guides/create-task.html
task("fast", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log("address", account.address);
    // console.log(account.provider);

    // Look up the balance
    if (account.provider) {
      let balance = await account.provider.getBalance(account.address);
      console.log("balance", balance);
    }
  }
});


// https://hardhat.org/config/
export default config;
