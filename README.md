# boringDAO-bridge

## Usage

### 1. Initialize configuration

Use below command to initialize the bridge, it will create .bridge directory with a configuration file in your home directory by default:
```shell
bridge init
```

You can also specify another place to put the .bridge directory by --repo option:
```shell
bridge --repo <repo> init
```

### 2. Edit the configuration file

Open the bridge.yaml file in .bridge directory, and edit below items:
#### [eth]
1. **addrs**: specify the eth rpc address this bridge will connect, you can configure more than one
1. **minConfirms**: minimal blocks to confirm a transaction, a number not less than **15** is recommended 
1. **crossLockContract**: specify the cross lock contract address deployed in eth network
1. **token**: specify the eth token contract address which will be cross over to bsc network
#### [bsc]
1. **addrs**: specify the bsc rpc address this bridge will connect, you can configure more than one
1. **minConfirms**: minimal blocks to confirm a transaction, a number not less than **15** is recommended
1. **bridgeContract**: specify the bridge contract address deployed in bsc network

### 3. Start the bridge

Use below command to start the bridge, it will use the configuration in the default home directory:
```shell
bridge start
```

Or if the bridge was initialized with --repo option, then you should use --repo to specify the repository as well:
```shell
bridge --repo <repo> start
```