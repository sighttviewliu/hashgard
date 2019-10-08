# What is Hashgard?

`hashgard` is the name of the Cosmos SDK application for the Cosmos Hub. It comes with 2 main entrypoints:

- `hashgard`: The Hashgard Daemon, runs a full-node of the `hashgard` application.
- `hashgardcli`: The Hashgard command-line interface, which enables interaction with a Hashgard full-node.

`hashgard` is built on the Cosmos SDK using the following modules:

- `x/auth`: Accounts and signatures.
- `x/bank`: Token transfers.
- `x/staking`: Staking logic.
- `x/mint`: Inflation logic.
- `x/distribution`: Fee distribution logic.
- `x/slashing`: Slashing logic.
- `x/gov`: Governance logic.
- `x/ibc`: Inter-blockchain transfers.
- `x/params`: Handles app-level parameters.
+ `x/contract` : Smart Contract

Next, learn how to [install Hashgard](./installation.md).
