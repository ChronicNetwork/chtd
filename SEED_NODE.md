# Seed Nodes for Chronic Network

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

### Validator Node Configuration

| Config Option             | Default      |
| ------------------------ | ------------ |
| DOUBLE_SIGN_CHECK_HEIGHT | 10           |
| ADDR_BOOK_STRICT         | 0            |
| PEX                      | 0            |


### Sentry Node Configuration

| Config Option             | Default      |
| ------------------------ | ------------ |
| ADDR_BOOK_STRICT         | 0            |
| PEX                      | 1            |

## Net

### Mainnet
```
- SEEDS_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/mainnet/seed-nodes.txt
- SYNC_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/mainnet/rpc-nodes.txt
- GENESIS_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/mainnet/genesis.json
```

### Testnet(Beta)
```
- SEEDS_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/testnet-beta/seed-nodes.txt
- SYNC_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/testnet-beta/rpc-nodes.txt
- GENESIS_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/testnet-beta/genesis.json
```

### Testnet(Athena)
```
- SEEDS_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/testnet-athena/seed-nodes.txt
- SYNC_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/testnet-athena/rpc-nodes.txt
- GENESIS_URL=https://raw.githubusercontent.com/ChronicNetwork/net/main/testnet-athena/genesis.json
```

