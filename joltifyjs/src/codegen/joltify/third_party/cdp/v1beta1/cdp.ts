import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../../../google/protobuf/timestamp";
import { Long, toTimestamp, fromTimestamp, DeepPartial } from "../../../../helpers";
import * as _m0 from "protobufjs/minimal";
/** CDP defines the state of a single collateralized debt position. */

export interface CDP {
  id: Long;
  owner: Uint8Array;
  type: string;
  collateral?: Coin;
  principal?: Coin;
  accumulatedFees?: Coin;
  feesUpdated?: Date;
  interestFactor: string;
}
/** CDP defines the state of a single collateralized debt position. */

export interface CDPSDKType {
  id: Long;
  owner: Uint8Array;
  type: string;
  collateral?: CoinSDKType;
  principal?: CoinSDKType;
  accumulated_fees?: CoinSDKType;
  fees_updated?: Date;
  interest_factor: string;
}
/** Deposit defines an amount of coins deposited by an account to a cdp */

export interface Deposit {
  cdpId: Long;
  depositor: string;
  amount?: Coin;
}
/** Deposit defines an amount of coins deposited by an account to a cdp */

export interface DepositSDKType {
  cdp_id: Long;
  depositor: string;
  amount?: CoinSDKType;
}
/** TotalPrincipal defines the total principal of a given collateral type */

export interface TotalPrincipal {
  collateralType: string;
  amount?: Coin;
}
/** TotalPrincipal defines the total principal of a given collateral type */

export interface TotalPrincipalSDKType {
  collateral_type: string;
  amount?: CoinSDKType;
}
/** TotalCollateral defines the total collateral of a given collateral type */

export interface TotalCollateral {
  collateralType: string;
  amount?: Coin;
}
/** TotalCollateral defines the total collateral of a given collateral type */

export interface TotalCollateralSDKType {
  collateral_type: string;
  amount?: CoinSDKType;
}
/** OwnerCDPIndex defines the cdp ids for a single cdp owner */

export interface OwnerCDPIndex {
  cdpIds: Long[];
}
/** OwnerCDPIndex defines the cdp ids for a single cdp owner */

export interface OwnerCDPIndexSDKType {
  cdp_ids: Long[];
}

function createBaseCDP(): CDP {
  return {
    id: Long.UZERO,
    owner: new Uint8Array(),
    type: "",
    collateral: undefined,
    principal: undefined,
    accumulatedFees: undefined,
    feesUpdated: undefined,
    interestFactor: ""
  };
}

export const CDP = {
  encode(message: CDP, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.id.isZero()) {
      writer.uint32(8).uint64(message.id);
    }

    if (message.owner.length !== 0) {
      writer.uint32(18).bytes(message.owner);
    }

    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }

    if (message.collateral !== undefined) {
      Coin.encode(message.collateral, writer.uint32(34).fork()).ldelim();
    }

    if (message.principal !== undefined) {
      Coin.encode(message.principal, writer.uint32(42).fork()).ldelim();
    }

    if (message.accumulatedFees !== undefined) {
      Coin.encode(message.accumulatedFees, writer.uint32(50).fork()).ldelim();
    }

    if (message.feesUpdated !== undefined) {
      Timestamp.encode(toTimestamp(message.feesUpdated), writer.uint32(58).fork()).ldelim();
    }

    if (message.interestFactor !== "") {
      writer.uint32(66).string(message.interestFactor);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CDP {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCDP();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.id = (reader.uint64() as Long);
          break;

        case 2:
          message.owner = reader.bytes();
          break;

        case 3:
          message.type = reader.string();
          break;

        case 4:
          message.collateral = Coin.decode(reader, reader.uint32());
          break;

        case 5:
          message.principal = Coin.decode(reader, reader.uint32());
          break;

        case 6:
          message.accumulatedFees = Coin.decode(reader, reader.uint32());
          break;

        case 7:
          message.feesUpdated = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 8:
          message.interestFactor = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<CDP>): CDP {
    const message = createBaseCDP();
    message.id = object.id !== undefined && object.id !== null ? Long.fromValue(object.id) : Long.UZERO;
    message.owner = object.owner ?? new Uint8Array();
    message.type = object.type ?? "";
    message.collateral = object.collateral !== undefined && object.collateral !== null ? Coin.fromPartial(object.collateral) : undefined;
    message.principal = object.principal !== undefined && object.principal !== null ? Coin.fromPartial(object.principal) : undefined;
    message.accumulatedFees = object.accumulatedFees !== undefined && object.accumulatedFees !== null ? Coin.fromPartial(object.accumulatedFees) : undefined;
    message.feesUpdated = object.feesUpdated ?? undefined;
    message.interestFactor = object.interestFactor ?? "";
    return message;
  }

};

function createBaseDeposit(): Deposit {
  return {
    cdpId: Long.UZERO,
    depositor: "",
    amount: undefined
  };
}

export const Deposit = {
  encode(message: Deposit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.cdpId.isZero()) {
      writer.uint32(8).uint64(message.cdpId);
    }

    if (message.depositor !== "") {
      writer.uint32(18).string(message.depositor);
    }

    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Deposit {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeposit();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.cdpId = (reader.uint64() as Long);
          break;

        case 2:
          message.depositor = reader.string();
          break;

        case 3:
          message.amount = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Deposit>): Deposit {
    const message = createBaseDeposit();
    message.cdpId = object.cdpId !== undefined && object.cdpId !== null ? Long.fromValue(object.cdpId) : Long.UZERO;
    message.depositor = object.depositor ?? "";
    message.amount = object.amount !== undefined && object.amount !== null ? Coin.fromPartial(object.amount) : undefined;
    return message;
  }

};

function createBaseTotalPrincipal(): TotalPrincipal {
  return {
    collateralType: "",
    amount: undefined
  };
}

export const TotalPrincipal = {
  encode(message: TotalPrincipal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TotalPrincipal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTotalPrincipal();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.amount = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<TotalPrincipal>): TotalPrincipal {
    const message = createBaseTotalPrincipal();
    message.collateralType = object.collateralType ?? "";
    message.amount = object.amount !== undefined && object.amount !== null ? Coin.fromPartial(object.amount) : undefined;
    return message;
  }

};

function createBaseTotalCollateral(): TotalCollateral {
  return {
    collateralType: "",
    amount: undefined
  };
}

export const TotalCollateral = {
  encode(message: TotalCollateral, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collateralType !== "") {
      writer.uint32(10).string(message.collateralType);
    }

    if (message.amount !== undefined) {
      Coin.encode(message.amount, writer.uint32(18).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TotalCollateral {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTotalCollateral();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.collateralType = reader.string();
          break;

        case 2:
          message.amount = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<TotalCollateral>): TotalCollateral {
    const message = createBaseTotalCollateral();
    message.collateralType = object.collateralType ?? "";
    message.amount = object.amount !== undefined && object.amount !== null ? Coin.fromPartial(object.amount) : undefined;
    return message;
  }

};

function createBaseOwnerCDPIndex(): OwnerCDPIndex {
  return {
    cdpIds: []
  };
}

export const OwnerCDPIndex = {
  encode(message: OwnerCDPIndex, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    writer.uint32(10).fork();

    for (const v of message.cdpIds) {
      writer.uint64(v);
    }

    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OwnerCDPIndex {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOwnerCDPIndex();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;

            while (reader.pos < end2) {
              message.cdpIds.push((reader.uint64() as Long));
            }
          } else {
            message.cdpIds.push((reader.uint64() as Long));
          }

          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<OwnerCDPIndex>): OwnerCDPIndex {
    const message = createBaseOwnerCDPIndex();
    message.cdpIds = object.cdpIds?.map(e => Long.fromValue(e)) || [];
    return message;
  }

};