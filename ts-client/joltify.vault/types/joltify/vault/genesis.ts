/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { CreatePool } from "./create_pool";
import { IssueToken } from "./issue_token";
import { OutboundTx } from "./outbound_tx";
import { coinsQuota } from "./quota";
import { Params, StandbyPower, Validators } from "./staking";

export const protobufPackage = "joltify.vault";

/** GenesisState defines the vault module's genesis state. */
export interface GenesisState {
  /** params defines all the paramaters of related to deposit. */
  params: Params | undefined;
  outboundTxList: OutboundTx[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  issueTokenList: IssueToken[];
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  createPoolList: CreatePool[];
  /** this line is used by starport scaffolding # ibc/genesis/proto */
  validatorinfoList: Validators[];
  latestTwoPool: CreatePool[];
  standbypowerList: StandbyPower[];
  feeCollectedList: Coin[];
  coinsQuota: coinsQuota | undefined;
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
    exported: false,
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
      coinsQuota.encode(message.coinsQuota, writer.uint32(74).fork()).ldelim();
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
          message.coinsQuota = coinsQuota.decode(reader, reader.uint32());
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

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      outboundTxList: Array.isArray(object?.outboundTxList)
        ? object.outboundTxList.map((e: any) => OutboundTx.fromJSON(e))
        : [],
      issueTokenList: Array.isArray(object?.issueTokenList)
        ? object.issueTokenList.map((e: any) => IssueToken.fromJSON(e))
        : [],
      createPoolList: Array.isArray(object?.createPoolList)
        ? object.createPoolList.map((e: any) => CreatePool.fromJSON(e))
        : [],
      validatorinfoList: Array.isArray(object?.validatorinfoList)
        ? object.validatorinfoList.map((e: any) => Validators.fromJSON(e))
        : [],
      latestTwoPool: Array.isArray(object?.latestTwoPool)
        ? object.latestTwoPool.map((e: any) => CreatePool.fromJSON(e))
        : [],
      standbypowerList: Array.isArray(object?.standbypowerList)
        ? object.standbypowerList.map((e: any) => StandbyPower.fromJSON(e))
        : [],
      feeCollectedList: Array.isArray(object?.feeCollectedList)
        ? object.feeCollectedList.map((e: any) => Coin.fromJSON(e))
        : [],
      coinsQuota: isSet(object.coinsQuota) ? coinsQuota.fromJSON(object.coinsQuota) : undefined,
      exported: isSet(object.exported) ? Boolean(object.exported) : false,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.outboundTxList) {
      obj.outboundTxList = message.outboundTxList.map((e) => e ? OutboundTx.toJSON(e) : undefined);
    } else {
      obj.outboundTxList = [];
    }
    if (message.issueTokenList) {
      obj.issueTokenList = message.issueTokenList.map((e) => e ? IssueToken.toJSON(e) : undefined);
    } else {
      obj.issueTokenList = [];
    }
    if (message.createPoolList) {
      obj.createPoolList = message.createPoolList.map((e) => e ? CreatePool.toJSON(e) : undefined);
    } else {
      obj.createPoolList = [];
    }
    if (message.validatorinfoList) {
      obj.validatorinfoList = message.validatorinfoList.map((e) => e ? Validators.toJSON(e) : undefined);
    } else {
      obj.validatorinfoList = [];
    }
    if (message.latestTwoPool) {
      obj.latestTwoPool = message.latestTwoPool.map((e) => e ? CreatePool.toJSON(e) : undefined);
    } else {
      obj.latestTwoPool = [];
    }
    if (message.standbypowerList) {
      obj.standbypowerList = message.standbypowerList.map((e) => e ? StandbyPower.toJSON(e) : undefined);
    } else {
      obj.standbypowerList = [];
    }
    if (message.feeCollectedList) {
      obj.feeCollectedList = message.feeCollectedList.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.feeCollectedList = [];
    }
    message.coinsQuota !== undefined
      && (obj.coinsQuota = message.coinsQuota ? coinsQuota.toJSON(message.coinsQuota) : undefined);
    message.exported !== undefined && (obj.exported = message.exported);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.outboundTxList = object.outboundTxList?.map((e) => OutboundTx.fromPartial(e)) || [];
    message.issueTokenList = object.issueTokenList?.map((e) => IssueToken.fromPartial(e)) || [];
    message.createPoolList = object.createPoolList?.map((e) => CreatePool.fromPartial(e)) || [];
    message.validatorinfoList = object.validatorinfoList?.map((e) => Validators.fromPartial(e)) || [];
    message.latestTwoPool = object.latestTwoPool?.map((e) => CreatePool.fromPartial(e)) || [];
    message.standbypowerList = object.standbypowerList?.map((e) => StandbyPower.fromPartial(e)) || [];
    message.feeCollectedList = object.feeCollectedList?.map((e) => Coin.fromPartial(e)) || [];
    message.coinsQuota = (object.coinsQuota !== undefined && object.coinsQuota !== null)
      ? coinsQuota.fromPartial(object.coinsQuota)
      : undefined;
    message.exported = object.exported ?? false;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
