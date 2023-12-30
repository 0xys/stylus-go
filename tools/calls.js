/* eslint-disable max-len */
const ethers = require('ethers');

const CALL_MARKER = '0a';
const DELEGATECALL_MARKER = '0b';
const STATICCALL_MARKER = '0c';

const STORE_MARKER = '01';
const LOAD_MARKER = '02';

const usage = 'node calls.js <dstAddress> <callmode: c|d|s> <calledAddress> <mode: s|l> [0x-stored_data]';

const calldataForCall = () => {
  const callModeStr = process.argv[3];
  if (!callModeStr) {
    console.log(usage);
    throw new Error('call mode not specified. c, d or s');
  }

  // this test contract address exposes storage that can be accessed through this test
  let calledContractAddress = process.argv[4];
  if (!calledContractAddress) {
    console.log(usage);
    throw new Error('calledAddress must be set');
  }
  if (!calledContractAddress.startsWith('0x')) {
    console.log(usage);
    throw new Error('calledAddress must be prefixed with 0x');
  }
  calledContractAddress = calledContractAddress.substring(2);

  let callMode = 'call';
  if (callModeStr == 'c') {
    callMode = 'call';
    console.log('callmode: call');
  } else if (callModeStr == 'd') {
    callMode = 'delegatecall';
    console.log('callmode: delegatecall');
  } else if (callModeStr == 's') {
    callMode = 'staticcall';
    console.log('callmode: staticcall');
  } else {
    console.log(usage);
    throw new Error(`unknown call mode ${callModeStr}`);
  }
  switch (callMode) {
    case 'call':
      return `${CALL_MARKER}${calledContractAddress}`;
    case 'delegatecall':
      return `${DELEGATECALL_MARKER}${calledContractAddress}`;
    case 'staticcall':
      return `${STATICCALL_MARKER}${calledContractAddress}`;
    default:
      onsole.log(usage);
      throw new Error(`unknown call mode ${callModeStr}. calldata cannot create`);
  }
};

const calldataForStorage = () => {
  const modeStr = process.argv[5];
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
  switch (mode) {
    case 'store':
      const dataStr = process.argv[6];
      if (!dataStr) {
        console.log(usage);
        throw new Error('stored data not specified');
      }
      let data = dataStr;
      if (dataStr.startsWith('0x')) {
        data = dataStr.substring(2);
      } else {
        console.log(usage);
        throw new Error('data must be a hex string prefixed with 0x');
      }
      console.log(`data: ${data}`);
      return `${STORE_MARKER}${data}`;
    case 'load':
      return `${LOAD_MARKER}`;
    default:
      console.log(usage);
      throw new Error(`unknown mode ${mode}. calldata cannot create.`);
  }
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
  const balanceInWei = await provider.getBalance(wallet.address);
  const balanceInEther = ethers.formatEther(balanceInWei);
  console.log(`Your account addresss is: ${wallet.address}`);
  console.log(`Your account balance is: ${balanceInEther} ETH`);

  const tx = {
    to: destAddr,
    gasPrice: ethers.parseUnits('100', 'mwei'),
    gasLimit: 10000000,
    // 0x<callMode(0)><calledContract(1:21)><mode(21)>[storedData(22:)]
    data: `0x${calldataForCall()}${calldataForStorage()}`,
  };

  const res = await wallet.sendTransaction(tx);
  console.log(res);
  console.log(res.hash);
  const txReceipt = await provider.waitForTransaction(
      res.hash,
  );

  console.log(txReceipt);
})();
