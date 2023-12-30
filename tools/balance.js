/* eslint-disable max-len */
const ethers = require('ethers');

(async () => {
  const privateKey = process.env.ETH_PRIVATE_KEY;
  const nodeType = process.env.NODETYPE;
  let url = 'http://localhost:8547';
  if (nodeType == 'remote') {
    url = 'https://stylus-testnet.arbitrum.io/rpc';
  }

  const provider = new ethers.JsonRpcProvider(url);
  const wallet = new ethers.Wallet(privateKey, provider);
  const balanceInWei = await provider.getBalance(wallet.address);
  const balanceInEther = ethers.formatEther(balanceInWei);
  console.log(`Your account addresss is: ${wallet.address}`);
  console.log(`Your account balance is: ${balanceInEther} ETH`);
})();
