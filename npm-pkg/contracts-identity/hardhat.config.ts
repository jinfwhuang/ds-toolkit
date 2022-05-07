import { HardhatUserConfig, task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";
import "hardhat-gas-reporter";
import "solidity-coverage";

const network = process.env.NETWORK || `ropsten`;

// import "dotenv";
// import 'dotenv/config'
import * as dotenv from "dotenv";

import "hardhat-abi-exporter";

dotenv.config({
  path: `${__dirname}/envs/${network}.conf`,
});

console.log('-----ENV-----');
console.log('pwd:', __dirname);
console.log('network:', process.env.NETWORK);
console.log('url', process.env.URL);
console.log('--------------');

const url = process.env.URL || ``;
const privateKey = process.env.PRIVATE_KEY || `0x${`11`.repeat(32)}`; // this is to avoid hardhat error

const config: HardhatUserConfig = {
  solidity: `0.8.4`,
  defaultNetwork: network,
  networks: {
    mainnet: {
      // chainId: 1,
      url,
      accounts: [privateKey],
    },
    ropsten: {
      url,
      accounts: [privateKey],
    },
    rinkeby: {
      url,
      accounts: [privateKey],
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
    path: './artifacts/abi',
    clear: true,
    flat: true,
    // only: [':ERC20$'],
    spacing: 2
  }  
};

// https://hardhat.org/guides/create-task.html
// @ts-ignore
task('fast', 'Prints the list of accounts', async (taskArgs, hre) => {
  console.log("task arguments:", taskArgs);
  const accounts = await hre.ethers.getSigners();

  /* eslint-disable no-await-in-loop */
  for (const account of accounts) {
    console.log(`address:`, account.address);
    // console.log(account.provider);

    // Look up the balance
    if (account.provider) {
      const balance = await account.provider.getBalance(account.address);
      console.log('balance:', balance);
    }
  }
});

// https://hardhat.org/config/
export default config;

// export {}
