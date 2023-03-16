import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface AddressV16 {
  address: Uint8Array[];
}
export interface AddressV16SDKType {
  address: Uint8Array[];
}
export interface OutboundTxV16_ItemsEntry {
  key: string;
  value?: AddressV16;
}
export interface OutboundTxV16_ItemsEntrySDKType {
  key: string;
  value?: AddressV16SDKType;
}
export interface OutboundTxV16 {
  index: string;
  items?: {
    [key: string]: AddressV16;
  };
}
export interface OutboundTxV16SDKType {
  index: string;
  items?: {
    [key: string]: AddressV16SDKType;
  };
}

function createBaseAddressV16(): AddressV16 {
  return {
    address: []
  };
}

export const AddressV16 = {
  encode(message: AddressV16, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.address) {
      writer.uint32(10).bytes(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddressV16 {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddressV16();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.address.push(reader.bytes());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<AddressV16>): AddressV16 {
    const message = createBaseAddressV16();
    message.address = object.address?.map(e => e) || [];
    return message;
  }

};

function createBaseOutboundTxV16_ItemsEntry(): OutboundTxV16_ItemsEntry {
  return {
    key: "",
    value: undefined
  };
}

export const OutboundTxV16_ItemsEntry = {
  encode(message: OutboundTxV16_ItemsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }

    if (message.value !== undefined) {
      AddressV16.encode(message.value, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutboundTxV16_ItemsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutboundTxV16_ItemsEntry();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;

        case 2:
          message.value = AddressV16.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<OutboundTxV16_ItemsEntry>): OutboundTxV16_ItemsEntry {
    const message = createBaseOutboundTxV16_ItemsEntry();
    message.key = object.key ?? "";
    message.value = object.value !== undefined && object.value !== null ? AddressV16.fromPartial(object.value) : undefined;
    return message;
  }

};

function createBaseOutboundTxV16(): OutboundTxV16 {
  return {
    index: "",
    items: {}
  };
}

export const OutboundTxV16 = {
  encode(message: OutboundTxV16, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }

    Object.entries(message.items).forEach(([key, value]) => {
      OutboundTxV16_ItemsEntry.encode({
        key: (key as any),
        value
      }, writer.uint32(18).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OutboundTxV16 {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOutboundTxV16();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;

        case 2:
          const entry2 = OutboundTxV16_ItemsEntry.decode(reader, reader.uint32());

          if (entry2.value !== undefined) {
            message.items[entry2.key] = entry2.value;
          }

          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<OutboundTxV16>): OutboundTxV16 {
    const message = createBaseOutboundTxV16();
    message.index = object.index ?? "";
    message.items = Object.entries(object.items ?? {}).reduce<{
      [key: string]: AddressV16;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = AddressV16.fromPartial(value);
      }

      return acc;
    }, {});
    return message;
  }

};