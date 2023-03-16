import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export enum DepositorInfo_DEPOSITTYPE {
  withdraw_proposal = 0,
  transfer_request = 1,
  deposit_close = 2,
  unset = 3,
  processed = 4,
  deactive = 5,
  UNRECOGNIZED = -1,
}
export const DepositorInfo_DEPOSITTYPESDKType = DepositorInfo_DEPOSITTYPE;
export function depositorInfo_DEPOSITTYPEFromJSON(object: any): DepositorInfo_DEPOSITTYPE {
  switch (object) {
    case 0:
    case "withdraw_proposal":
      return DepositorInfo_DEPOSITTYPE.withdraw_proposal;

    case 1:
    case "transfer_request":
      return DepositorInfo_DEPOSITTYPE.transfer_request;

    case 2:
    case "deposit_close":
      return DepositorInfo_DEPOSITTYPE.deposit_close;

    case 3:
    case "unset":
      return DepositorInfo_DEPOSITTYPE.unset;

    case 4:
    case "processed":
      return DepositorInfo_DEPOSITTYPE.processed;

    case 5:
    case "deactive":
      return DepositorInfo_DEPOSITTYPE.deactive;

    case -1:
    case "UNRECOGNIZED":
    default:
      return DepositorInfo_DEPOSITTYPE.UNRECOGNIZED;
  }
}
export function depositorInfo_DEPOSITTYPEToJSON(object: DepositorInfo_DEPOSITTYPE): string {
  switch (object) {
    case DepositorInfo_DEPOSITTYPE.withdraw_proposal:
      return "withdraw_proposal";

    case DepositorInfo_DEPOSITTYPE.transfer_request:
      return "transfer_request";

    case DepositorInfo_DEPOSITTYPE.deposit_close:
      return "deposit_close";

    case DepositorInfo_DEPOSITTYPE.unset:
      return "unset";

    case DepositorInfo_DEPOSITTYPE.processed:
      return "processed";

    case DepositorInfo_DEPOSITTYPE.deactive:
      return "deactive";

    case DepositorInfo_DEPOSITTYPE.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
export interface DepositorInfo {
  investorId: string;
  depositorAddress: Uint8Array;
  poolIndex: string;
  lockedAmount?: Coin;
  withdrawalAmount?: Coin;
  incentiveAmount?: Coin;
  linkedNFT: string[];
  depositType: DepositorInfo_DEPOSITTYPE;
  pendingInterest?: Coin;
}
export interface DepositorInfoSDKType {
  investor_id: string;
  depositor_address: Uint8Array;
  pool_index: string;
  locked_amount?: CoinSDKType;
  withdrawal_amount?: CoinSDKType;
  incentive_amount?: CoinSDKType;
  linkedNFT: string[];
  deposit_type: DepositorInfo_DEPOSITTYPE;
  pending_interest?: CoinSDKType;
}

function createBaseDepositorInfo(): DepositorInfo {
  return {
    investorId: "",
    depositorAddress: new Uint8Array(),
    poolIndex: "",
    lockedAmount: undefined,
    withdrawalAmount: undefined,
    incentiveAmount: undefined,
    linkedNFT: [],
    depositType: 0,
    pendingInterest: undefined
  };
}

export const DepositorInfo = {
  encode(message: DepositorInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.investorId !== "") {
      writer.uint32(10).string(message.investorId);
    }

    if (message.depositorAddress.length !== 0) {
      writer.uint32(18).bytes(message.depositorAddress);
    }

    if (message.poolIndex !== "") {
      writer.uint32(26).string(message.poolIndex);
    }

    if (message.lockedAmount !== undefined) {
      Coin.encode(message.lockedAmount, writer.uint32(34).fork()).ldelim();
    }

    if (message.withdrawalAmount !== undefined) {
      Coin.encode(message.withdrawalAmount, writer.uint32(42).fork()).ldelim();
    }

    if (message.incentiveAmount !== undefined) {
      Coin.encode(message.incentiveAmount, writer.uint32(50).fork()).ldelim();
    }

    for (const v of message.linkedNFT) {
      writer.uint32(58).string(v!);
    }

    if (message.depositType !== 0) {
      writer.uint32(64).int32(message.depositType);
    }

    if (message.pendingInterest !== undefined) {
      Coin.encode(message.pendingInterest, writer.uint32(74).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DepositorInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDepositorInfo();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.investorId = reader.string();
          break;

        case 2:
          message.depositorAddress = reader.bytes();
          break;

        case 3:
          message.poolIndex = reader.string();
          break;

        case 4:
          message.lockedAmount = Coin.decode(reader, reader.uint32());
          break;

        case 5:
          message.withdrawalAmount = Coin.decode(reader, reader.uint32());
          break;

        case 6:
          message.incentiveAmount = Coin.decode(reader, reader.uint32());
          break;

        case 7:
          message.linkedNFT.push(reader.string());
          break;

        case 8:
          message.depositType = (reader.int32() as any);
          break;

        case 9:
          message.pendingInterest = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<DepositorInfo>): DepositorInfo {
    const message = createBaseDepositorInfo();
    message.investorId = object.investorId ?? "";
    message.depositorAddress = object.depositorAddress ?? new Uint8Array();
    message.poolIndex = object.poolIndex ?? "";
    message.lockedAmount = object.lockedAmount !== undefined && object.lockedAmount !== null ? Coin.fromPartial(object.lockedAmount) : undefined;
    message.withdrawalAmount = object.withdrawalAmount !== undefined && object.withdrawalAmount !== null ? Coin.fromPartial(object.withdrawalAmount) : undefined;
    message.incentiveAmount = object.incentiveAmount !== undefined && object.incentiveAmount !== null ? Coin.fromPartial(object.incentiveAmount) : undefined;
    message.linkedNFT = object.linkedNFT?.map(e => e) || [];
    message.depositType = object.depositType ?? 0;
    message.pendingInterest = object.pendingInterest !== undefined && object.pendingInterest !== null ? Coin.fromPartial(object.pendingInterest) : undefined;
    return message;
  }

};