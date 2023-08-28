# Cosmos SDK Escrow Module for DreddSecure

This repository contains the implementation of the Escrow Module based on the Cosmos SDK, which serves as the backend for the DreddSecure application, providing secure escrow services within the Cosmos ecosystem.

## Getting Started

You can set up and run the Escrow Module using the Ignite CLI. There are two ways to get it running: locally in a Linux or macOS environment using Go and Ignite, or using Docker with the provided npm scripts.

### Option 1: Local Setup

1. Install [Go](https://go.dev/doc/install), [Ignite](https://docs.ignite.com/welcome/install) and [NodeJs](https://nodejs.org/en/download) on your system.
2. Clone this repository and navigate to the project directory.
3. Configure the blockchain by [modifying the `config.yml` file](https://docs.ignite.com/references/config).
4. Run the following command to install dependencies, build, initialize, and start the blockchain:

```
ignite chain serve
```

You should see the following output: 

```
  Blockchain is running
  
  ✔ Added account alice with address cosmos152qptu0s32xgvg2z8t9kxzz9ff3cvmqqcvpwsx and mnemonic:
    exotic oyster foam twist access group midnight garden sentence little art strike
    bubble cover ladder bleak antique point pigeon clinic name goddess fence cable  
  
  ✔ Added account bob with address cosmos1rx2geme8mkgpud9w3ahva4urfq56zsllg647qs and mnemonic:
    potato company tongue fine super club avoid language spring wrist income globe
    increase staff loan adapt vanish shield mind meat tuna video drip swim        
  
  🌍 Tendermint node: http://0.0.0.0:26657
  🌍 Blockchain API: http://0.0.0.0:1317
  🌍 Token faucet: http://0.0.0.0:4500
  
  ⋆ Data directory: /home/user/.dredd-secure
  ⋆ App binary: /home/user/go/bin/dredd-secured
  
  Press the 'q' key to stop serve
```

### Option 2: Docker Setup
Clone this repository and navigate to the project directory.

Run the following commands to build and start the blockchain using Docker:

**Windows Users**
```
npm run build-windows
npm run start
```

**Mac & Linux Users**
```
npm run build
npm run start
```

## Running Tests
To run tests for the Escrow Module, use the following Makefile command:

```
make test
```

## CLI Commands
Once the chain is running, you can test the module with the CLI commands. To explore the cli commands, type 

```
dredd-secured --help
```

**For example**, you can create a new escrow using this command

```
dredd-secured tx escrow create-escrow 100token 1000stake 10token 1588148578 2788148978 "" --from alice
```

And then query the created escrow:
```
dredd-secured query escrow list-escrow
```

To explore all the commands, check out [Ignite CLI Commands](https://docs.ignite.com/references/cli)


## Additional Resources
Here are some additional resources and links that might be helpful:

- [Ignite CLI Documentation](https://docs.ignite.com/)
- [Cosmos SDK Documentation](https://docs.cosmos.network/main)
