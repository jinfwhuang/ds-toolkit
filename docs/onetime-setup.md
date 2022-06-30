## One-time Setup Tips

### Running ```npx hardhat node ``` in [contracts-identity](/npm-pkg/contracts-identity/):
- If you run into ```npm ERR! Unsupported URL Type "workspace": workspace:* ``` error
    - Run ```yarn add --dev typescript``` instead of ```npm install --save-dev typescript ```
    - To run ```npm install --save-dev ts-node```, replace ```"workspace:*"``` with ```"file:*"``` in [my-controls](/experimental/my-controls/package.json) and [contracts-identity](/npm-pkg/contracts-identity/package.json)