# boringDAO-bridge

## Usage

### 1. Initialize configuration

Use below command to initialize the bridge, it will create .bridge directory with a configuration file in your home directory by default:
```shell
bridge init
```

You can also specify another place to put the configuration file by --repo option:
```shell
bridge --repo <repo> init
```

### 2. Edit the configuration file

Open the bridge.yaml file in .bridge directory, and edit below items:

#### [center]
```
  name = "matic_test"
  addrs = [
    "https://rpc-mumbai.maticvigil.com"
  ]
  chainID = 80001 # main 137, testnet 80001
  minConfirms = 1 # min confirms
  gasLimit = 1500000 # gas Limit
  gasFeeRate = 1.8 # newGasPrice = gasPrice * gasFeeRate
  centerContract = "0xCbb3f8292CbE44efb8A059eA2406bD36F5ab1652" # twoway center contract
  [center.index]
    421611 = 1 # specify egde‘s chainId with cross out index
```
#### [[edges]]
```
  isFilter = false # whether to filter events by block height, if not then [edge.index] is valid
  name = "arbi_test"
  addrs = [
    "https://rinkeby.arbitrum.io/rpc" # rpc host
  ]
  chainID = 421611
  minConfirms = 1 # min confirms
  gasLimit = 1500000 # gas Limit
  gasFeeRate = 1.2  # newGasPrice = gasPrice * gasFeeRate
  edgeContract = "0x3caf319515A122E98Da710304f31F237a9b304C7" # twoway edge contract
  depositedHeight = 0  # Deposit Event‘s block height
  crossOutedHeight = 0  # CrossOut Event’s block height
  [edge.index]
    421611 = 1 # specify egde‘s chainId with cross out index
```
### 3. Start the bridge

Use below command to start the bridge, it will use the configuration in the default home directory:
```shell
bridge start
```

Or if the bridge was initialized with --repo option, then you should use --repo to specify the repository as well:
```shell
bridge --repo <repo> start
```
