# RXP


βοΈFor issue disclosure, check out [SECURITY.md](./SECURITY.md) βοΈ

π For release procedures, check out [RELEASES.md](./RELEASES.md). π

**RXP** is an open source platform for inter-operable smart contracts which automatically execute, control or document a procedure of events and actions 
according to the terms of the contract or agreement to be valid and usable across multiple sovereign networks.

RXP is a **sovereign public blockchain** in the Cosmos ecosystem. It aims to provide a sandbox environment for the deployment 
of such inter-operable smart contracts. The network serves as a **decentralized, permissionless, and censorship resistant** zone 
for developers to efficiently and securely launch application specific smart contracts. Developers can leverage proven frameworks 
and compile code in various languages, like **Rust & Go** with the potential future addition of C and C++.
Battle-tested contract modules such as CosmWasm will allow for decentralized applications (dApps) to be compiled as robust and secure multi-chain smart contracts.
EVM support and additional specialized modules can be introduced after genesis subject to onchain governance.

At the heart of Cosmos is the Inter Blockchain Communication Protocol (IBC), which sets the table for an inter-operable base layer 0 
that can be used to transfer data packets across thousands of independent networks supporting IBC. 
Naturally, the next evolutionary milestone is to enable cross-network smart contracts.

The RXP blockchain is built using the **Cosmos SDK framework**. 
Cosmos SDK is a generalized framework that simplifies the process of building secure blockchain applications on top of Tendermint BFT. 
It is based on two major principles: Modularity and capabilities-based security.

Agreement on the network is reached via **Tendermint BFT consensus**.

Tendermint BFT is a solution that packages the networking and consensus layers of a blockchain into a generic engine, 
allowing developers to focus on application development as opposed to the complex underlying protocol. 
As a result, Tendermint saves hundreds of hours of development time.

RXP originates from a **community driven initiative**, prompted by developers, validators and delegators in the Cosmos ecosystem.
The common vision is to preserve the neutrality, performance, and reliability of the Cosmos Hub. This is achieved by offloading smart contract deployment to a dedicated sister Hub. 
RXP plans to make an early connection to the Cosmos Hub enabling IBC transfers, cross-chain smart contracts and shared security.

A decentralized launch of the RXP main-net was enabled by a large set of independent validators from across the blockchain space.
$RXP, the native asset, has many functions like securing the RXP Hub and serving as a work token to give access to on-chain governance voting rights 
and to provide utility in the deployment and execution of smart contracts.


**What differentiates RXP from other Smart Contract networks?**

βͺοΈ Interoperable smart contracts

βͺοΈ First mover advantage

βͺοΈ Open source

βͺοΈ Permissionless 

βͺοΈ Modular

βͺοΈ Wasm + (EVM)

βͺοΈ Compilation in multiple languages Rust & Go (C,C++)

βͺοΈ Highly scalable

βͺοΈ Ease of use/developer ergonomics

βͺοΈ Free & fair asset distribution (100% to staked atom only)

βͺοΈ Onchain governance

βͺοΈ Balanced governance (Zero top heavy control) 

βͺοΈ Grass roots community

βͺοΈ Decentralized

**Tokenomics & reward shedule** (updated on 24.07.2021)

βͺοΈ **Ticker**: RXP

βͺοΈ **Decimals**: 6

βͺοΈ **Unit**: uRXP

βͺοΈ **Supply**: Snapshot of Cosmoshub-3 at 06:00 PM UTC on Feb 18th 2021

βͺοΈ **Rewards**: Fixed yearly reward schedule (Reward model below)

βͺοΈ **Community pool tax**: 5% of block rewards


β Circulating supply at genesis 64.903.243 $RXP (64.9 Million)

β Max supply (After year 12): 185.562.268 RXP (185.5 Million)


**Supply Breakdown**

βͺοΈ Stakedrop: 30.663.193 $RXP

βͺοΈ Community Pool: 20.000.000 $RXP

βͺοΈ Development Reserve (Multi-sig): 11.866.708 $RXP

βͺοΈ Smart Contract Challenges: 2.373.341 $RXP


**Genesis Distribution**

A 1:1 stakedrop is allocated to $ATOM stakers, giving the $RXP genesis supply to staked $ATOM balances that had their assets bonded 
at the time of the Stargate snapshot on Feb. 18th 6:00 PM UTC. 
Addresses that qualify are included in the RXP genesis block at launch. 
All exchange validators & their delegators are excluded from the genesis allocation. Additionally any unbonded ATOM at the time of the snapshot is excluded.
A whale cap was voted in by the community, effectively hard-capping $ATOM accounts at 50 thousand $ATOM in order to ensure a less top heavy allocation.
Approx 10% of the supply difference is allocated to the development reserve (multi-sig) address for the funding of core-development efforts. The remaining 90% of the excess supply to be allocated in the following ways (20 million $RXP community pool, Smart contract competition 2.373.341,66 million to be managed/distributed by the multi-sig committee. The remaining difference will not be included in the genesis file ie. burned)



**Reward Schedule**

βͺοΈ Phase 1: Fixed inflation 40% 

New RXP in year 1 = (+25.961.297)

Supply after year 1 = 90.864.540 $RXP


βͺοΈ Phase 2: Fixed inflation 20% 

New RXP in year 2 = (+18.172.908)

Supply after year 2 = 109.037.449 RXP


βͺοΈ Phase 3: Fixed inflation 10% 

New RXP in year 3= (+10.903.744)

Supply after year 3 = 119.941.194 RXP


Once the inflation reaches 10% it gradually reduces on a fixed 1% basis each year.


βͺοΈ Phase 4 = Fixed 9% (+10.794.707) Supply = 130.735.901 RXP

βͺοΈ Phase 5 = Fixed 8% (+10.458.872) Supply = 141.194.773 RXP

βͺοΈ Phase 6 = Fixed 7% (+9.883.634) Supply = 151.078.407 RXP

βͺοΈ Phase 7 = Fixed 6% (+9.064.704) Supply = 160.143.112  RXP

βͺοΈ Phase 8 = Fixed 5% (+8.007.155) Supply = 168.150.267  RXP

βͺοΈ Phase 9 = Fixed 4% (+6.726.010) Supply = 174.876.278 RXP

βͺοΈ Phase 10 = Fixed 3% (+5.246.288) Supply = 180.122.566 RXP

βͺοΈ Phase 11 = Fixed 2% (+3.602.451) Supply = 183.725.018 RXP

βͺοΈ Phase 12 = Fixed 1% (+1.837.250) Supply = 185.562.268 RXP 

RXP MAX SUPPLY (185.5 Million)

After year 12 the inflation reward schedule ends. 
Network incentives would primarily come from smart contract usage and regular tx fees generated on the network.

**RXP** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

If you have [Docker](https://www.docker.com/) installed, then you can run a local node with a single command.

```bash
STAKE_TOKEN=urxpx UNSAFE_CORS=true docker-compose up
```

## Learn more

- [RXP](https://rxpchain.com)
- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Telegram](https://t.me/RXPNetwork)
- [Discord](https://discord.gg/QcWPfK4gJ2)

