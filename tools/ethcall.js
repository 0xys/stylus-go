/* eslint-disable max-len */
const ethers = require('ethers');

(async () => {
  const url = 'http://localhost:8547';
  const destAddr = process.argv[2];
  if (!destAddr) {
    console.log('no destination address');
    return;
  }
  const provider = new ethers.JsonRpcProvider(url);
  const response = await provider.send('eth_call', [
    {
      'from': null,
      'to': destAddr,
      'data':
        '0x70a082310000000000000000000000006E0d01A76C3Cf4288372a29124A26D4353EE51BE',
    },
    'latest',
  ]);

  console.log(response);
})();
