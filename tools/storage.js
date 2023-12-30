/* eslint-disable max-len */
const ethers = require('ethers');

const STORE_MARKER = '01';
const LOAD_MARKER = '02';
const LOAD2_MARKER = '03';

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
  const modeStr = process.argv[3];
  if (!modeStr) {
    console.log('mode not specified. s or l');
    return;
  }
  let mode = 'load';
  if (modeStr == 's') {
    mode = 'store';
    console.log('mode: store');
  } else if (modeStr == 'l') {
    mode = 'load';
    console.log('mode: load');
  } else if (modeStr == 'l2') {
    mode = 'load2';
    console.log('mode: load2');
  } else {
    console.log(`unknown mode ${modeStr}`);
    return;
  }

  const provider = new ethers.JsonRpcProvider(url);
  const wallet = new ethers.Wallet(privateKey, provider);
  const balanceInWei = await provider.getBalance(wallet.address);
  const balanceInEther = ethers.formatEther(balanceInWei);
  console.log(`Your account addresss is: ${wallet.address}`);
  console.log(`Your account balance is: ${balanceInEther} ETH`);

  if (mode == 'store') {
    const dataStr = process.argv[4];
    if (!dataStr) {
      console.log('stored data not specified');
      return;
    }
    let data = dataStr;
    if (dataStr.startsWith('0x')) {
      data = dataStr.substring(2);
    } else {
      console.log('data must be a hex string prefixed with 0x');
      return;
    }
    console.log(`data: ${data}`);

    const tx = {
      to: destAddr,
      gasPrice: ethers.parseUnits('100', 'mwei'),
      gasLimit: 2000000,
      data: `0x${STORE_MARKER}${data}`,
    };

    const res = await wallet.sendTransaction(tx);
    console.log(res);
    console.log(res.hash);
    const txReceipt = await provider.waitForTransaction(
        res.hash,
    );

    console.log(txReceipt);
  } else if (mode == 'load2') {
    const tx = {
      to: destAddr,
      gasPrice: ethers.parseUnits('100', 'mwei'),
      gasLimit: 2000000,
      data: `0x${LOAD2_MARKER}`,
    };

    const res = await wallet.sendTransaction(tx);
    console.log(res);
    console.log(res.hash);
    const txReceipt = await provider.waitForTransaction(
        res.hash,
    );

    console.log(txReceipt);
  } else {
    const response = await provider.send('eth_call', [
      {
        'from': null,
        'to': destAddr,
        'data':
            `0x${LOAD_MARKER}`,
      },
      'latest',
    ]);

    console.log(response);
  }
})();
