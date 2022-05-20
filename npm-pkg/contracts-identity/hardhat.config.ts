import { HardhatUserConfig, task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";
import "hardhat-gas-reporter";
import "solidity-coverage";

import {
  HardhatNetworkAccountsUserConfig,
  HardhatNetworkAccountUserConfig,
} from "hardhat/types";

// const network = process.env.NETWORK || `ropsten`;
const network = process.env.NETWORK || `hardhat`;

import "hardhat-abi-exporter";

import * as dotenv from "dotenv";
dotenv.config({
  path: `${__dirname}/envs/${network}.conf`,
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

const hardhatAccounts: HardhatNetworkAccountUserConfig[] = [
  {
    privateKey: privateKey1,
    balance: "19999900" + "0".repeat(18),
  },
  {
    privateKey: privateKey2,
    balance: "29999900" + "0".repeat(18),
  },
  {
    privateKey: privateKey3,
    balance: "39999900" + "0".repeat(18),
  },
];

const config: HardhatUserConfig = {
  solidity: `0.8.8`,
  defaultNetwork: network,
  networks: {
    mainnet: {
      // chainId: 1,
      url,
      accounts: [privateKey1, privateKey1, privateKey3],
    },
    hardhat: {
      chainId: 1337,
      accounts: hardhatAccounts,
    },
    localhost: {
      chainId: 1337,
      url,
      accounts: [privateKey1, privateKey1, privateKey3],
    },
    ropsten: {
      url,
      accounts: [privateKey1, privateKey1, privateKey3],
    },
    rinkeby: {
      url,
      accounts: [privateKey1, privateKey1, privateKey3],
    },
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: `USD`,
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
  abiExporter: {
    path: "./artifacts/abi",
    clear: true,
    flat: true,
    // only: [':ERC20$'],
    spacing: 2,
  },
};

// https://hardhat.org/guides/create-task.html
// @ts-ignore
task("fast", "Prints the list of accounts", async (taskArgs, hre) => {
  console.log("task arguments:", taskArgs);
  const accounts = await hre.ethers.getSigners();

  /* eslint-disable no-await-in-loop */
  for (const account of accounts) {
    console.log(`address:`, account.address);
    // console.log(account.provider);

    // Look up the balance
    if (account.provider) {
      const balance = await account.provider.getBalance(account.address);
      console.log("balance:", balance);
    }
  }
});

// https://hardhat.org/config/
export default config;

// export {}
