import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface Entity {
  address: Uint8Array;
  feecoin: Coin[];
}
export interface EntitySDKType {
  address: Uint8Array;
  feecoin: CoinSDKType[];
}
export interface Proposals {
  entry: Entity[];
}
export interface ProposalsSDKType {
  entry: EntitySDKType[];
}
export interface OutboundTx_ItemsEntry {
  key: string;
  value?: Proposals;
}
export interface OutboundTx_ItemsEntrySDKType {
  key: string;
  value?: ProposalsSDKType;
}
export interface OutboundTx {
  index: string;
  processed: boolean;
  items?: {
    [key: string]: Proposals;
  };
  chainType: string;
  inTxHash: string;
  receiverAddress: Uint8Array;
  needMint: boolean;
  feecoin: Coin[];
}
export interface OutboundTxSDKType {
  index: string;
  processed: boolean;
  items?: {
    [key: string]: ProposalsSDKType;
  };
  chainType: string;
  inTxHash: string;
  receiverAddress: Uint8Array;
  needMint: boolean;
  feecoin: CoinSDKType[];
}

function createBaseEntity(): Entity {
  return {
    address: new Uint8Array(),
    feecoin: []
  };
}

export const Entity = {
  encode(message: Entity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address.length !== 0) {
      writer.uint32(10).bytes(message.address);
    }

    for (const v of message.feecoin) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Entity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEntity();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.address = reader.bytes();
          break;

        case 2:
          message.feecoin.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Entity>): Entity {
    const message = createBaseEntity();
    message.address = object.address ?? new Uint8Array();
    message.feecoin = object.feecoin?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};

function createBaseProposals(): Proposals {
  return {
    entry: []
  };
}

export const Proposals = {
  encode(message: Proposals, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.entry) {
      Entity.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Proposals {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProposals();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.entry.push(Entity.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Proposals>): Proposals {
    const message = createBaseProposals();
    message.entry = object.entry?.map(e => Entity.fromPartial(e)) || [];
    return message;
  }

};

function createBaseOutboundTx_ItemsEntry(): OutboundTx_ItemsEntry {
  return {
    key: "",
    value: undefined
  };
}

export const OutboundTx_ItemsEntry = {
  encode(message: OutboundTx_ItemsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }

    if (message.value !== undefined) {
      Proposals.encode(message.value, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutboundTx_ItemsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutboundTx_ItemsEntry();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;

        case 2:
          message.value = Proposals.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<OutboundTx_ItemsEntry>): OutboundTx_ItemsEntry {
    const message = createBaseOutboundTx_ItemsEntry();
    message.key = object.key ?? "";
    message.value = object.value !== undefined && object.value !== null ? Proposals.fromPartial(object.value) : undefined;
    return message;
  }

};

function createBaseOutboundTx(): OutboundTx {
  return {
    index: "",
    processed: false,
    items: {},
    chainType: "",
    inTxHash: "",
    receiverAddress: new Uint8Array(),
    needMint: false,
    feecoin: []
  };
}

export const OutboundTx = {
  encode(message: OutboundTx, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }

    if (message.processed === true) {
      writer.uint32(16).bool(message.processed);
    }

    Object.entries(message.items).forEach(([key, value]) => {
      OutboundTx_ItemsEntry.encode({
        key: (key as any),
        value
      }, writer.uint32(26).fork()).ldelim();
    });

    if (message.chainType !== "") {
      writer.uint32(34).string(message.chainType);
    }

    if (message.inTxHash !== "") {
      writer.uint32(42).string(message.inTxHash);
    }

    if (message.receiverAddress.length !== 0) {
      writer.uint32(50).bytes(message.receiverAddress);
    }

    if (message.needMint === true) {
      writer.uint32(56).bool(message.needMint);
    }

    for (const v of message.feecoin) {
      Coin.encode(v!, writer.uint32(66).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutboundTx {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutboundTx();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;

        case 2:
          message.processed = reader.bool();
          break;

        case 3:
          const entry3 = OutboundTx_ItemsEntry.decode(reader, reader.uint32());

          if (entry3.value !== undefined) {
            message.items[entry3.key] = entry3.value;
          }

          break;

        case 4:
          message.chainType = reader.string();
          break;

        case 5:
          message.inTxHash = reader.string();
          break;

        case 6:
          message.receiverAddress = reader.bytes();
          break;

        case 7:
          message.needMint = reader.bool();
          break;

        case 8:
          message.feecoin.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<OutboundTx>): OutboundTx {
    const message = createBaseOutboundTx();
    message.index = object.index ?? "";
    message.processed = object.processed ?? false;
    message.items = Object.entries(object.items ?? {}).reduce<{
      [key: string]: Proposals;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = Proposals.fromPartial(value);
      }

      return acc;
    }, {});
    message.chainType = object.chainType ?? "";
    message.inTxHash = object.inTxHash ?? "";
    message.receiverAddress = object.receiverAddress ?? new Uint8Array();
    message.needMint = object.needMint ?? false;
    message.feecoin = object.feecoin?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};