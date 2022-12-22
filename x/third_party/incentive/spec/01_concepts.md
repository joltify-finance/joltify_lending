<!--
order: 1
-->

# Concepts

This module presents an implementation of user incentives that are controlled by governance. When users take a certain action, for example opening a CDP, they become eligible for rewards. Rewards are **opt in** meaning that users must submit a message before the claim deadline to claim their rewards. The goals and background of this module were subject of a previous Kava governance proposal, which can be found [here](https://ipfs.io/ipfs/QmSYedssC3nyQacDJmNcREtgmTPyaMx2JX7RNkMdAVkdkr/user-growth-fund-proposal.pdf)

## General Reward Distribution

Rewards target various user activity. For example, usdx borrowed from bnb CDPs, btcb supplied to the jolt money market, or owned shares in a swap jolt/usdx pool.

Each second, the rewards accumulate at a rate set in the params, eg 100 ujolt per second. These are then distributed to all users ratably based on their percentage involvement in the rewarded activity. For example if a user holds 1% of all funds deposited to the jolt/usdt pool. They will receive 1% of the total rewards each second.

The quantity tracking a user's involvement is referred to as "source shares" in the code. And the total across all users the "total source shares". The quotient then gives their percentage involvement, eg if a user borrowed 10,000 usdx, and there is 100,000 usdx borrowed by all users, then they will get 10% of rewards.

## Efficiency

Paying out rewards to every user every block would be slow and lead to long block times. Instead rewards are calculated much less frequently.

Every block a global tracker adds up total rewards paid out per unit of user involvement. A user's specific reward can then be calculated as needed based on their current source shares.

Users' rewards must be updated whenever their source shares change. This happens through hooks into other modules that run before deposits/borrows/supplies etc.
