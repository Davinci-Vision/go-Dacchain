# go-Dacchain

## Davinci Project

DPOS / BFT + DAI consensus minimizes the risk of hard fork. Parallel AI technology enables fast transaction processing by dividing transactions of different tokens, which not allow to be affected by each other. It provides multi-asset issuance, standardized token issuance procedures, and main chain-level processing speed and scalability.


### Decentralized blockchain operated by AI

Since Bitcoin began actively trading, the biggest technical challenge has been to overcome the limitations of high mining costs and transaction speeds. Various methods have been proposed to address this problem, but the value of decentralization was inevitably compromised for rapid transaction processing. Even in the case of D-Pos method, which is developing by solving speed problems and producing multiple derived technologies, there is a problem that a small number of BPs get to determine the operation direction of blockchain. By considering these problems, the method does not meet the originally intended purpose of developing value-neutral currency through the perfect decentralization. Nevertheless, it is also true that in order to cope with the growing number of dApps and new business types, it requires constant technical upgrades by someone.


#### Structure
![img](https://davinci.vision/static/img/tech/tech_02_1.png)

Build DAI on the existing D-Pos based Mainnet to manage transactions and learn user usage patterns. For this purpose, a total of two AI engines are installed.

1. Client Analysis Engine : Client Analysis Engine: Analyze and monitor the propensity of each participant and node to increase the operational efficiency of the blockchain. Dynamically involved in the generation and removal of side-chains.

2. AI Cube Learning Engine : AI Cube Learning Engine: Analyze and learn patterns of each transaction. The learned data will be used for future technical upgrades of the chain.


#### Common Console Commond Example
```

dac.blockNumber

dac.getBlock(blockHashOrBlockNumber)

dac.accounts

dac.getTransaction(transactionHash)

personal.newAccount(passphrase)

personal.sendTransaction({from:'affress',to:'address',value:web3.toWei(100,'dac'),action:0}, "password")

admin.startRPC("0.0.0.0", 8545)

admin.stopRPC()
```
