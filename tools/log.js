const ethers = require('ethers');

(async () => {
  const url = 'http://localhost:8547';
  const destAddr = process.argv[2];
  if (!destAddr) {
    console.log('no destination address');
    return;
  }
  const provider = new ethers.JsonRpcProvider(url);
  const blockNum = await provider.getBlockNumber();
  const logs = await provider.getLogs({
    fromBlock: blockNum-5,
    toBlock: blockNum,

  });

  const found = logs.filter((x) => x.address == destAddr);
  for (let i =0; i<found.length; i++) {
    console.log('-', i);
    console.log('  - txhash:', found[i].transactionHash);
    console.log('  - data:', found[i].data);
    console.log('  - topics:', found[i].topics);
    console.log('');
  }
})();
