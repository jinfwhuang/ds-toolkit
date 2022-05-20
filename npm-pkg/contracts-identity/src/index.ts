import { ethers } from "hardhat";

function getPubkey(privkey: string): string {
  const key = new ethers.utils.SigningKey(privkey);
  return key.publicKey;
}

export { getPubkey };
