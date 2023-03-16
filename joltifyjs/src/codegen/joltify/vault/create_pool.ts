import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface PoolProposal {
  poolPubKey: string;
  poolAddr: Uint8Array;
  nodes: Uint8Array[];
}
export interface PoolProposalSDKType {
  pool_pubKey: string;
  pool_addr: Uint8Array;
  nodes: Uint8Array[];
}
export interface CreatePool {
  blockHeight: string;
  validators: Uint8Array[];
  proposal: PoolProposal[];
}
export interface CreatePoolSDKType {
  block_height: string;
  validators: Uint8Array[];
  proposal: PoolProposalSDKType[];
}

function createBasePoolProposal(): PoolProposal {
  return {
    poolPubKey: "",
    poolAddr: new Uint8Array(),
    nodes: []
  };
}

export const PoolProposal = {
  encode(message: PoolProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolPubKey !== "") {
      writer.uint32(10).string(message.poolPubKey);
    }

    if (message.poolAddr.length !== 0) {
      writer.uint32(18).bytes(message.poolAddr);
    }

    for (const v of message.nodes) {
      writer.uint32(26).bytes(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolProposal();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.poolPubKey = reader.string();
          break;

        case 2:
          message.poolAddr = reader.bytes();
          break;

        case 3:
          message.nodes.push(reader.bytes());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<PoolProposal>): PoolProposal {
    const message = createBasePoolProposal();
    message.poolPubKey = object.poolPubKey ?? "";
    message.poolAddr = object.poolAddr ?? new Uint8Array();
    message.nodes = object.nodes?.map(e => e) || [];
    return message;
  }

};

function createBaseCreatePool(): CreatePool {
  return {
    blockHeight: "",
    validators: [],
    proposal: []
  };
}

export const CreatePool = {
  encode(message: CreatePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.blockHeight !== "") {
      writer.uint32(10).string(message.blockHeight);
    }

    for (const v of message.validators) {
      writer.uint32(18).bytes(v!);
    }

    for (const v of message.proposal) {
      PoolProposal.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreatePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreatePool();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.blockHeight = reader.string();
          break;

        case 2:
          message.validators.push(reader.bytes());
          break;

        case 3:
          message.proposal.push(PoolProposal.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<CreatePool>): CreatePool {
    const message = createBaseCreatePool();
    message.blockHeight = object.blockHeight ?? "";
    message.validators = object.validators?.map(e => e) || [];
    message.proposal = object.proposal?.map(e => PoolProposal.fromPartial(e)) || [];
    return message;
  }

};