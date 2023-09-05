# Cosmos SDK Escrow Module for DreddSecure

This repository contains the implementation of the Escrow Module based on the Cosmos SDK, which serves as the backend for the DreddSecure application, providing secure escrow services within the Cosmos ecosystem.

## Getting Started

You can set up and run the Escrow Module using the Ignite CLI. There are two ways to get it running: locally in a Linux or macOS environment using Go and Ignite, or using Docker with the provided npm scripts.

### Option 1: Local Setup

1. Install [Go](https://go.dev/doc/install), [Ignite](https://docs.ignite.com/welcome/install), [Hermes](https://hermes.informal.systems/quick-start/installation.html) and [NodeJs](https://nodejs.org/en/download) on your system.
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

5. Configure the Hermes Relayer

   1. Create a config.toml file in the .hermes directory and then paste the following configuration inside of it to create a connection with the Bandchain's Laozi testnet:

   ```
   nano ~/.hermes/config.toml
   ```

   ```
    # The global section has parameters that apply globally to the relayer operation.
    [global]


    # Specify the verbosity for the relayer logging output. Default: 'info'

    # Valid options are 'error', 'warn', 'info', 'debug', 'trace'.

    log_level = 'trace'

    # Specify the mode to be used by the relayer. [Required]

    [mode]

    # Specify the client mode.

    [mode.clients]

    # Whether or not to enable the client workers. [Required]

    enabled = true

    # Whether or not to enable periodic refresh of clients. [Default: true]

    # Note: Even if this is disabled, clients will be refreshed automatically if

    # there is activity on a connection or channel they are involved with.

    refresh = true

    # Whether or not to enable misbehaviour detection for clients. [Default: false]

    misbehaviour = true

    # Specify the connections mode.

    [mode.connections]

    # Whether or not to enable the connection workers for handshake completion. [Required]

    enabled = true

    # Specify the channels mode.

    [mode.channels]

    # Whether or not to enable the channel workers for handshake completion. [Required]

    enabled = true

    # Specify the packets mode.

    [mode.packets]

    # Whether or not to enable the packet workers. [Required]

    enabled = true

    # Parametrize the periodic packet clearing feature.

    # Interval (in number of blocks) at which pending packets

    # should be eagerly cleared. A value of '0' will disable

    # periodic packet clearing. [Default: 100]

    clear_interval = 100

    # Whether or not to clear packets on start. [Default: false]

    clear_on_start = true

    # Toggle the transaction confirmation mechanism.

    # The tx confirmation mechanism periodically queries the `/tx_search` RPC

    # endpoint to check that previously-submitted transactions

    # (to any chain in this config file) have delivered successfully.

    # Experimental feature. Affects telemetry if set to false.

    # [Default: true]

    tx_confirmation = true

    # The REST section defines parameters for Hermes' built-in RESTful API.

    # https://hermes.informal.systems/rest.html

    [rest]

    # Whether or not to enable the REST service. Default: false

    enabled = true

    # Specify the IPv4/6 host over which the built-in HTTP server will serve the RESTful

    # API requests. Default: 127.0.0.1

    host = '127.0.0.1'

    # Specify the port over which the built-in HTTP server will serve the restful API

    # requests. Default: 3000

    port = 3000

    # The telemetry section defines parameters for Hermes' built-in telemetry capabilities.

    # https://hermes.informal.systems/telemetry.html

    [telemetry]

    # Whether or not to enable the telemetry service. Default: false

    enabled = true

    # Specify the IPv4/6 host over which the built-in HTTP server will serve the metrics

    # gathered by the telemetry service. Default: 127.0.0.1

    host = '127.0.0.1'

    # Specify the port over which the built-in HTTP server will serve the metrics gathered

    # by the telemetry service. Default: 3001

    port = 3001

    [[chains]]
    id = 'dreddsecure'
    rpc_addr = 'http://localhost:26657'
    grpc_addr = 'http://localhost:9090'
    event_source = { mode = 'push', url = 'ws://localhost:26657/websocket', batch_delay = '500ms' }
    rpc_timeout = '30s'
    account_prefix = 'cosmos'
    key_name = 'requester'
    store_prefix = 'ibc'
    default_gas = 5000000
    max_gas = 15000000
    gas_price = { price = 0, denom = 'ustake' }
    gas_multiplier = 1.1
    max_msg_num = 20
    max_tx_size = 209715
    clock_drift = '20s'
    max_block_time = '30s'
    trusting_period = '10days'
    trust_threshold = { numerator = '1', denominator = '3' }
    address_type = { derivation = 'cosmos' }

    [[chains]]
    id = 'band-laozi-testnet6'
    rpc_addr = 'https://rpc.laozi-testnet6.bandchain.org:443'
    grpc_addr = 'https://laozi-testnet6.bandchain.org:443'
    event_source = { mode = 'push', url = 'wss://rpc.laozi-testnet6.bandchain.org:443/websocket', batch_delay = '500ms' }
    rpc_timeout = '10s'
    account_prefix = 'band'
    key_name = 'testkey'
    store_prefix = 'ibc'
    default_gas = 100000
    max_gas = 10000000
    gas_price = { price = 0.0025, denom = 'uband' }
    gas_multiplier = 1.1
    max_msg_num = 30
    max_tx_size = 2097152
    clock_drift = '5s'
    max_block_time = '10s'
    trusting_period = '14days'
    trust_threshold = { numerator = '1', denominator = '3' }
    address_type = { derivation = 'cosmos' }
   ```

   2. Create a json file at the root of the project and enter the relayer account mnemonic. It's important to note that you need to have funds in this account in order to send transactions on each chain (you can request BAND testnet tokens [here](https://docs.bandchain.org/develop/api-endpoints#laozi-testnet-6)).

   3. Add Keys to Hermes by following the commands:

      **DreddSecure (local) chain**:

      ```
      hermes keys add --chain dreddsecure --mnemonic-file "<name-of-your-file>.json"
      ```

      **Bandchain Testnet**:

      ```
      hermes keys add --chain band-laozi-testnet6 --mnemonic-file "<name-of-your-file>.json"
      ```

   4. Create clients, connections & channels (you can view your connection on [Bandchain's block explorer](https://laozi-testnet6.cosmoscan.io/relayers)):

   ```
   hermes create channel --a-chain dreddsecure --b-chain band-laozi-testnet6 --a-port escrow --b-port oracle --new-client-connection
   ```

   5. Start the Hermes relayer

   ```
   hermes start
   ```

   6. Test by sending a OracleRequestPacketData to bandchain oracle using the CLI. You should see the packet on the [Bandchain's block explorer](https://laozi-testnet6.cosmoscan.io/relayers).

   ```
   dredd-secured tx escrow send-oracle-request-packet-data escrow <your-channel-id> 401 "[\"BTC\", \"LTC\"]" 16 10 1000uband 10920 178750 --from alice
   ```

### Option 2: Docker Setup

Clone this repository and navigate to the project directory.

1. Run the following commands to build and start the blockchain using Docker:

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

2. Open a bash terminal in your Docker container

```
docker exec -it dredd-secure /bin/bash
```

3. In the bash terminal, follow the steps above to Configure the Hermes relayer (in the [local setup section](#option-1-local-setup)).

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
