import { expect } from "chai";
import { ethers } from "hardhat";

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

describe(`Token contract`, function () {
  it(`Deployment should assign the total supply of tokens to the owner`, async function () {
    const Token = await ethers.getContractFactory(`Token`);
    const hardhatToken = await Token.deploy();

    const [owner] = await ethers.getSigners();
    const ownerBalance = await hardhatToken.balanceOf(owner.address);
    const tokenTotalSupply = await hardhatToken.totalSupply()

    expect(tokenTotalSupply).to.equal(ownerBalance);
  });
});
