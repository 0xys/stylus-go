const ethers = require('ethers');

const addr = '0x1f1aebcA347A03218bEDEe7173D150db1d3c5388';
const coder = new ethers.AbiCoder();

/**
 *
 * @param {string} sig e.g. 'transfer(address,uint256)'
 * @return string
 */
const funcSignature = (sig) => {
  return ethers.id(sig).substring(0, 10);
};

const call = async (provider) => {
  const returnData = await provider.call({
    to: '0xAddress',
    data: ethers.concat([
      funcSignature('transfer(address,uint256)'),
      coder.encode(['address', 'uint256'], ['0x1f1aebcA347A03218bEDEe7173D150db1d3c5388', '0x11']),
    ]),
  });
  return returnData;
};


(async () => {
  const privateKey = process.env.ETH_PRIVATE_KEY;
  const nodeType = process.env.NODETYPE;
  let url = 'http://localhost:8547';
  if (nodeType == 'remote') {
    url = 'https://stylus-testnet.arbitrum.io/rpc';
  }
  const destAddr = process.argv[2];
  if (!destAddr) {
    console.log('no destination address');
    return;
  }

  const provider = new ethers.JsonRpcProvider(url);
  const wallet = new ethers.Wallet(privateKey, provider);
  const balanceInWei = await provider.getBalance(addr);
  const balanceInEther = ethers.formatEther(balanceInWei);
  console.log(`Your account balance is: ${balanceInEther} ETH`);

  // const a = ethers.concat([
  //     funcSignature('transfer(address,uint256)'),
  //     coder.encode(['address', 'uint256'], ['0x1f1aebcA347A03218bEDEe7173D150db1d3c5388', '0x11'])
  // ])

  const tx = {
    to: destAddr,
    gasPrice: ethers.parseUnits('100', 'mwei'),
    gasLimit: 2000000,
    data: '0xa4b000000000000000000073657175656e636572a4b000000000000000000073657175656e636572',
  };

  const res = await wallet.sendTransaction(tx);
  console.log(res);
  console.log(res.hash);
})();
