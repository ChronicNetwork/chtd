# Chronic Token
## NOTE: THIS REPO IS DEPRECIATED! IF YOU PLAN ON SPINNING UP A NODE FOR MAIN NETWORK, PLEASE REFER TO THE [PROPER REPO](https://github.com/ChronicNetwork/cht)
**chronic** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

**Note**: Requires [Go 1.17+](https://golang.org/dl/), [starport 0.19.0+](https://starport.com/)

### Dockerized

We provide a docker image to help with test setups. There are two modes to use it

Build: `docker build -t carlosbie/chronic:latest .`  or pull from dockerhub

### Dev server
Using starport
```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

Using docker
```
docker run --rm
-e CHTD_MONIKER=chronic-node-1
-e SEEDS_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/seed-nodes.txt
-e SYNC_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/rpc-nodes.txt
-e GENESIS_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/genesis.json
carlosbie/chronic:latest
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `chtd start` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)

## Options for docker of seed nodes

| Config Option             | Default      | Description           |
| ------------------------ | ------------ | --------------------- |
| CHTD_MONIKER             | chronic-node | Name of Moniker       |
| SEEDS_URL(required)      | none         | url of seeds          |
| SYNC_URL(required)       | none         | url of sync addresses |
| GENESIS_URL(required)    | none         | url of genesis file    |
| DOUBLE_SIGN_CHECK_HEIGHT | 0            | 10 for validator      |
| ADDR_BOOK_STRICT         | 1            | 0 or 1                |
| PEX                      | 1            | 0 or 1                |

