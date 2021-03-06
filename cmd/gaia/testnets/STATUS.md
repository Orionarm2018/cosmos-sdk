# TESTNET STATUS

## *June 10, 2018, 8:30 EST* - Gaia-6001 consensus failure

- Validator unbonding and revocation activity caused a consensus failure
- There is a bug in the staking module that must be fixed
- The team is taking its time to look into this and release a fix following a
  proper protocol for hotfix upgrades to the testnet
- Please stay tuned!

## *June 9, 2018, 14:00 EST* - New Release

- Released gaia
  [v0.18.0](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.18.0) with
  update for Tendermint
  [v0.20.0](https://github.com/tendermint/tendermint/releases/tag/v0.20.0)
- Includes bug fix for declaring candidacy from the command line

## *June 8, 2018, 23:30 EST* - Gaia-6001 is making blocks

- +2/3 of the voting power is finally online for Gaia-6001 and it is making
  blocks!
- This is a momentous achievement - a successful asynchronous decentralized
  testnet launch
- Congrats everyone!

## *June 8, 2018, 12:00 EST* - New Testnet Gaia-6001

- After some confusion around testnet deployment and a contention testnet
  hardfork, a new genesis file and network was released for `gaia-6001`

## *June 7, 2018, 9:00 EST* - New Testnet Gaia-6000

- Released a new `genesis.json` file for `gaia-6000`
- Initial validators include those that were most active in
  the gaia-5001 testnet
- Join the network via gaia `v0.18.0-rc0`

## *June 5, 2018, 21:00 EST* - New Release

- Released gaia
  [v0.17.5](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.17.5) 
  with update for Tendermint
  [v0.19.9](https://github.com/tendermint/tendermint/releases/tag/v0.19.9)
- Fixes many bugs!
    - evidence gossipping 
    - mempool deadlock
    - WAL panic
    - memory leak
- Please update to this to put a stop to the rampant invalid evidence gossiping
  :)

## *May 31, 2018, 14:00 EST* - New Release

- Released gaia
  [v0.17.4](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.17.4) with update for Tendermint v0.19.7
- Fixes a WAL bug and some more
- Please update to this if you have trouble restarting a node

## *May 31, 2018, 2:00 EST* - Testnet Halt

- A validator equivocated last week and Evidence is being rampantly gossipped
- Peers that can't process the evidence (either too far behind or too far ahead) are disconnecting from the peers that
  sent it, causing high peer turn-over
- The high peer turn-over may be causing a memory-leak, resulting in some nodes
  crashing and the testnet halting
- We need to fix some issues in the EvidenceReactor to address this and also
  investigate the possible memory-leak

## *May 29, 2018* - New Release

- Released v0.17.3 with update for Tendermint v0.19.6
- Fixes fast-sync bug
- Please update to this to sync with the testnet
