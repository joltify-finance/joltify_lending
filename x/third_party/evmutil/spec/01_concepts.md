<!--
order: 1
-->

# Concepts

## EVM Gas Denom

In order to use the EVM and be compatible with existing clients, the gas denom used by the EVM must be in 18 decimals. Since `ujolt` has 6 decimals of precision, it cannot be used as the EVM gas denom directly.

To use the Jolt token on the EVM, the evmutil module provides an `EvmBankKeeper` that is responsible for the conversion of `ujolt` and `ajolt`. A user's excess `ajolt` balance is stored in the `x/evmutil` store, while its `ujolt` balance remains in the cosmos-sdk `x/bank` module.

## `EvmBankKeeper` Overview

The `EvmBankKeeper` provides access to an account's total `ajolt` balance and the ability to transfer, mint, and burn `ajolt`. If anything other than the `ajolt` denom is requested, the `EvmBankKeeper` will panic.

This keeper implements the `x/evm` module's `BankKeeper` interface to enable the usage of `ajolt` denom on the EVM.

### `x/evm` Parameter Requirements

Since the EVM denom `ajolt` is required to use the `EvmBankKeeper`, it is necessary to set the `EVMDenom` param of the `x/evm` module to `ajolt`.

### Balance Calculation of `ajolt`

The `ajolt` balance of an account is derived from an account's **spendable** `ujolt` balance times 10^12 (to derive its `ajolt` equivalent), plus the account's excess `ajolt` balance that can be accessed via the module `Keeper`.

### `ajolt` <> `ujolt` Conversion

When an account does not have sufficient `ajolt` to cover a transfer or burn, the `EvmBankKeeper` will try to swap 1 `ujolt` to its equivalent `ajolt` amount. It does this by transferring 1 `ujolt` from the sender to the `x/evmutil` module account, then adding the equivalent `ajolt` amount to the sender's balance in the module state.

In reverse, if an account has enough `ajolt` balance for one or more `ujolt`, the excess `ajolt` balance will be converted to `ujolt`. This is done by removing the excess `ajolt` balance in the module store, then transferring the equivalent `ujolt` coins from the `x/evmutil` module account to the target account.

The swap logic ensures that all `ajolt` is backed by the equivalent `ujolt` balance stored in the module account.

## ERC20 token <> sdk.Coin Conversion

`x/evmutil` facilitates moving assets between Jolt's EVM and Cosmos co-chains. This must be handled differently depending on which co-chain to which the asset it native. The messages controlling these flows involve two accounts:
1. The _initiator_ who sends coins from their co-chain
2. The _receiver_ who receives coins on the other co-chain

When converting assets from the EVM to the Cosmos co-chain, the initiator is an 0x EVM address and the receiver is a `jolt1` Bech32 address.

When converting assets from the Cosmos co-chain to the EVM, the initiator is a `jolt1` Bech32 address and the receiver is an 0x EVM address.

### Cosmos-Native Assets

`sdk.Coin`s native to the Cosmos co-chain can be converted to an ERC-20 representing the coin in the EVM. This works by transferring the coin from the initiator to `x/evmutil`'s module account and then minting an ERC-20 token to the receiver. Converting back, the initiator's ERC-20 representation of the coin is burned and the original Cosmos-native asset is transferred to the receiver.

Cosmos-native asset converstion is done through the use of the `MsgConvertCosmosCoinToERC20` & `MsgConvertCosmosCoinFromERC20` messages (see **[Messages](03_messages.md)**).

Only Cosmos co-chain denominations that are in the `AllowedCosmosDenoms` param (see **[Params](05_params.md)**) can be converted via these messages.

`AllowedCosmosDenoms` can be altered through governance.

The ERC20 contracts are deployed and managed by x/evmutil. The contract is deployed on first convert of the coin. Once deployed, the addresses of the contracts can be queried via the `DeployedCosmosCoinContracts` query (`deployed_cosmos_coin_contracts` endpoint).

If a denom is removed from the `AllowedCosmosDenoms` param, existing ERC20 tokens can be converted back to the underlying sdk.Coin via `MsgConvertCosmosCoinFromERC20`, but no conversions from sdk.Coin -> ERC via `MsgConvertCosmosCoinToERC20` are allowed.

### EVM-Native Assets

ERC-20 tokens native to the EVM can be converted into an `sdk.Coin` in the Cosmos ecosystem. This works by transferring the tokens to `x/evmutil`'s module account and then minting an `sdk.Coin` to the receiver. Converting back is the inverse: the `sdk.Coin` of the initiator is burned and the original ERC-20 tokens that were locked into the module account are transferred back to the receiver.

EVM-native asset conversion is done through the use of the `MsgConvertERC20ToCoin` & `MsgConvertCoinToERC20` messages (see **[Messages](03_messages.md)**).

Only ERC20 contract address that are in the `EnabledConversionPairs` param (see **[Params](05_params.md)**) can be converted via these messages.

`EnabledConversionPairs` can be altered through governance.

## Module Keeper

The module Keeper provides access to an account's excess `ajolt` balance and the ability to update the balance.
