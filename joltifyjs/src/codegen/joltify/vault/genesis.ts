import { Params, ParamsSDKType, Validators, ValidatorsSDKType, StandbyPower, StandbyPowerSDKType } from "./staking";
import { OutboundTx, OutboundTxSDKType } from "./outbound_tx";
import { IssueToken, IssueTokenSDKType } from "./issue_token";
import { CreatePool, CreatePoolSDKType } from "./create_pool";
import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import { CoinsQuota, CoinsQuotaSDKType } from "./quota";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/** GenesisState defines the vault module's genesis state. */

export interface GenesisState {
  /** params defines all the paramaters of related to deposit. */
  params?: Params;
  outboundTxList: OutboundTx[];
  /** this line is used by starport scaffolding # genesis/proto/state */

  issueTokenList: IssueToken[];
  /**
   * this line is used by starport scaffolding # genesis/proto/stateField
   * this line is used by starport scaffolding # ibc/genesis/proto
   */

  createPoolList: CreatePool[];
  validatorinfoList: Validators[];
  latestTwoPool: CreatePool[];
  standbypowerList: StandbyPower[];
  feeCollectedList: Coin[];
  coinsQuota?: CoinsQuota;
  exported: boolean;
}
/** GenesisState defines the vault module's genesis state. */

export interface GenesisStateSDKType {
  params?: ParamsSDKType;
  outbound_tx_list: OutboundTxSDKType[];
  issue_token_list: IssueTokenSDKType[];
  create_pool_list: CreatePoolSDKType[];
  validatorinfo_list: ValidatorsSDKType[];
  latest_twoPool: CreatePoolSDKType[];
  standbypower_list: StandbyPowerSDKType[];
  feeCollected_list: CoinSDKType[];
  coinsQuota?: CoinsQuotaSDKType;
  exported: boolean;
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    outboundTxList: [],
    issueTokenList: [],
    createPoolList: [],
    validatorinfoList: [],
    latestTwoPool: [],
    standbypowerList: [],
    feeCollectedList: [],
    coinsQuota: undefined,
    exported: false
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.outboundTxList) {
      OutboundTx.encode(v!, writer.uint32(42).fork()).ldelim();
    }

    for (const v of message.issueTokenList) {
      IssueToken.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.createPoolList) {
      CreatePool.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    for (const v of message.validatorinfoList) {
      Validators.encode(v!, writer.uint32(50).fork()).ldelim();
    }

    for (const v of message.latestTwoPool) {
      CreatePool.encode(v!, writer.uint32(82).fork()).ldelim();
    }

    for (const v of message.standbypowerList) {
      StandbyPower.encode(v!, writer.uint32(58).fork()).ldelim();
    }

    for (const v of message.feeCollectedList) {
      Coin.encode(v!, writer.uint32(66).fork()).ldelim();
    }

    if (message.coinsQuota !== undefined) {
      CoinsQuota.encode(message.coinsQuota, writer.uint32(74).fork()).ldelim();
    }

    if (message.exported === true) {
      writer.uint32(32).bool(message.exported);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;

        case 5:
          message.outboundTxList.push(OutboundTx.decode(reader, reader.uint32()));
          break;

        case 2:
          message.issueTokenList.push(IssueToken.decode(reader, reader.uint32()));
          break;

        case 3:
          message.createPoolList.push(CreatePool.decode(reader, reader.uint32()));
          break;

        case 6:
          message.validatorinfoList.push(Validators.decode(reader, reader.uint32()));
          break;

        case 10:
          message.latestTwoPool.push(CreatePool.decode(reader, reader.uint32()));
          break;

        case 7:
          message.standbypowerList.push(StandbyPower.decode(reader, reader.uint32()));
          break;

        case 8:
          message.feeCollectedList.push(Coin.decode(reader, reader.uint32()));
          break;

        case 9:
          message.coinsQuota = CoinsQuota.decode(reader, reader.uint32());
          break;

        case 4:
          message.exported = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = createBaseGenesisState();
    message.params = object.params !== undefined && object.params !== null ? Params.fromPartial(object.params) : undefined;
    message.outboundTxList = object.outboundTxList?.map(e => OutboundTx.fromPartial(e)) || [];
    message.issueTokenList = object.issueTokenList?.map(e => IssueToken.fromPartial(e)) || [];
    message.createPoolList = object.createPoolList?.map(e => CreatePool.fromPartial(e)) || [];
    message.validatorinfoList = object.validatorinfoList?.map(e => Validators.fromPartial(e)) || [];
    message.latestTwoPool = object.latestTwoPool?.map(e => CreatePool.fromPartial(e)) || [];
    message.standbypowerList = object.standbypowerList?.map(e => StandbyPower.fromPartial(e)) || [];
    message.feeCollectedList = object.feeCollectedList?.map(e => Coin.fromPartial(e)) || [];
    message.coinsQuota = object.coinsQuota !== undefined && object.coinsQuota !== null ? CoinsQuota.fromPartial(object.coinsQuota) : undefined;
    message.exported = object.exported ?? false;
    return message;
  }

};