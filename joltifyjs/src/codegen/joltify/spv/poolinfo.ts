import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../google/protobuf/timestamp";
import { Long, toTimestamp, fromTimestamp, DeepPartial } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
export enum PoolInfo_POOLSTATUS {
  ACTIVE = 0,
  INACTIVE = 1,
  CLOSED = 2,
  PREPARE = 3,
  CLOSING = 4,
  UNRECOGNIZED = -1,
}
export const PoolInfo_POOLSTATUSSDKType = PoolInfo_POOLSTATUS;
export function poolInfo_POOLSTATUSFromJSON(object: any): PoolInfo_POOLSTATUS {
  switch (object) {
    case 0:
    case "ACTIVE":
      return PoolInfo_POOLSTATUS.ACTIVE;

    case 1:
    case "INACTIVE":
      return PoolInfo_POOLSTATUS.INACTIVE;

    case 2:
    case "CLOSED":
      return PoolInfo_POOLSTATUS.CLOSED;

    case 3:
    case "PREPARE":
      return PoolInfo_POOLSTATUS.PREPARE;

    case 4:
    case "CLOSING":
      return PoolInfo_POOLSTATUS.CLOSING;

    case -1:
    case "UNRECOGNIZED":
    default:
      return PoolInfo_POOLSTATUS.UNRECOGNIZED;
  }
}
export function poolInfo_POOLSTATUSToJSON(object: PoolInfo_POOLSTATUS): string {
  switch (object) {
    case PoolInfo_POOLSTATUS.ACTIVE:
      return "ACTIVE";

    case PoolInfo_POOLSTATUS.INACTIVE:
      return "INACTIVE";

    case PoolInfo_POOLSTATUS.CLOSED:
      return "CLOSED";

    case PoolInfo_POOLSTATUS.PREPARE:
      return "PREPARE";

    case PoolInfo_POOLSTATUS.CLOSING:
      return "CLOSING";

    case PoolInfo_POOLSTATUS.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
export enum PoolInfo_POOLTYPE {
  JUNIOR = 0,
  SENIOR = 1,
  UNRECOGNIZED = -1,
}
export const PoolInfo_POOLTYPESDKType = PoolInfo_POOLTYPE;
export function poolInfo_POOLTYPEFromJSON(object: any): PoolInfo_POOLTYPE {
  switch (object) {
    case 0:
    case "JUNIOR":
      return PoolInfo_POOLTYPE.JUNIOR;

    case 1:
    case "SENIOR":
      return PoolInfo_POOLTYPE.SENIOR;

    case -1:
    case "UNRECOGNIZED":
    default:
      return PoolInfo_POOLTYPE.UNRECOGNIZED;
  }
}
export function poolInfo_POOLTYPEToJSON(object: PoolInfo_POOLTYPE): string {
  switch (object) {
    case PoolInfo_POOLTYPE.JUNIOR:
      return "JUNIOR";

    case PoolInfo_POOLTYPE.SENIOR:
      return "SENIOR";

    case PoolInfo_POOLTYPE.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
export interface PoolInfo {
  index: string;
  poolName: string;
  linkedProject: number;
  ownerAddress: Uint8Array;
  apy: string;
  totalAmount?: Coin;
  payFreq: number;
  reserveFactor: string;
  /**
   * string            pool_nFT_class      = 9 [
   *    (cosmos_proto.scalar)  = "cosmos.Class",
   *    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/x/nft.Class",
   *    (gogoproto.nullable)   = false
   *  ];
   */

  poolNFTIds: string[];
  lastPaymentTime?: Date;
  poolStatus: PoolInfo_POOLSTATUS;
  borrowedAmount?: Coin;
  poolInterest: string;
  projectLength: Long;
  usableAmount?: Coin;
  targetAmount?: Coin;
  poolType: PoolInfo_POOLTYPE;
  escrowInterestAmount: string;
  escrowPrincipalAmount?: Coin;
  withdrawProposalAmount?: Coin;
  projectDueTime?: Date;
  withdrawAccounts: Uint8Array[];
  transferAccounts: Uint8Array[];
}
export interface PoolInfoSDKType {
  index: string;
  pool_name: string;
  linked_project: number;
  owner_address: Uint8Array;
  apy: string;
  total_amount?: CoinSDKType;
  pay_freq: number;
  reserve_factor: string;
  pool_nFT_ids: string[];
  last_payment_time?: Date;
  pool_status: PoolInfo_POOLSTATUS;
  borrowed_amount?: CoinSDKType;
  pool_interest: string;
  project_length: Long;
  usable_amount?: CoinSDKType;
  target_amount?: CoinSDKType;
  pool_type: PoolInfo_POOLTYPE;
  escrow_interest_amount: string;
  escrow_principal_amount?: CoinSDKType;
  withdraw_proposal_amount?: CoinSDKType;
  project_due_time?: Date;
  withdraw_accounts: Uint8Array[];
  transfer_accounts: Uint8Array[];
}
export interface PoolWithInvestors {
  poolIndex: string;
  investors: string[];
}
export interface PoolWithInvestorsSDKType {
  pool_index: string;
  investors: string[];
}
export interface PoolDepositedInvestors {
  poolIndex: string;
  walletAddress: Uint8Array[];
}
export interface PoolDepositedInvestorsSDKType {
  pool_index: string;
  wallet_address: Uint8Array[];
}
export interface WalletsLinkPool {
  walletAddress: Uint8Array;
  poolsIndex: string[];
}
export interface WalletsLinkPoolSDKType {
  wallet_address: Uint8Array;
  pools_index: string[];
}

function createBasePoolInfo(): PoolInfo {
  return {
    index: "",
    poolName: "",
    linkedProject: 0,
    ownerAddress: new Uint8Array(),
    apy: "",
    totalAmount: undefined,
    payFreq: 0,
    reserveFactor: "",
    poolNFTIds: [],
    lastPaymentTime: undefined,
    poolStatus: 0,
    borrowedAmount: undefined,
    poolInterest: "",
    projectLength: Long.UZERO,
    usableAmount: undefined,
    targetAmount: undefined,
    poolType: 0,
    escrowInterestAmount: "",
    escrowPrincipalAmount: undefined,
    withdrawProposalAmount: undefined,
    projectDueTime: undefined,
    withdrawAccounts: [],
    transferAccounts: []
  };
}

export const PoolInfo = {
  encode(message: PoolInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }

    if (message.poolName !== "") {
      writer.uint32(18).string(message.poolName);
    }

    if (message.linkedProject !== 0) {
      writer.uint32(24).int32(message.linkedProject);
    }

    if (message.ownerAddress.length !== 0) {
      writer.uint32(34).bytes(message.ownerAddress);
    }

    if (message.apy !== "") {
      writer.uint32(42).string(message.apy);
    }

    if (message.totalAmount !== undefined) {
      Coin.encode(message.totalAmount, writer.uint32(50).fork()).ldelim();
    }

    if (message.payFreq !== 0) {
      writer.uint32(56).int32(message.payFreq);
    }

    if (message.reserveFactor !== "") {
      writer.uint32(66).string(message.reserveFactor);
    }

    for (const v of message.poolNFTIds) {
      writer.uint32(74).string(v!);
    }

    if (message.lastPaymentTime !== undefined) {
      Timestamp.encode(toTimestamp(message.lastPaymentTime), writer.uint32(82).fork()).ldelim();
    }

    if (message.poolStatus !== 0) {
      writer.uint32(88).int32(message.poolStatus);
    }

    if (message.borrowedAmount !== undefined) {
      Coin.encode(message.borrowedAmount, writer.uint32(98).fork()).ldelim();
    }

    if (message.poolInterest !== "") {
      writer.uint32(106).string(message.poolInterest);
    }

    if (!message.projectLength.isZero()) {
      writer.uint32(112).uint64(message.projectLength);
    }

    if (message.usableAmount !== undefined) {
      Coin.encode(message.usableAmount, writer.uint32(122).fork()).ldelim();
    }

    if (message.targetAmount !== undefined) {
      Coin.encode(message.targetAmount, writer.uint32(130).fork()).ldelim();
    }

    if (message.poolType !== 0) {
      writer.uint32(136).int32(message.poolType);
    }

    if (message.escrowInterestAmount !== "") {
      writer.uint32(146).string(message.escrowInterestAmount);
    }

    if (message.escrowPrincipalAmount !== undefined) {
      Coin.encode(message.escrowPrincipalAmount, writer.uint32(154).fork()).ldelim();
    }

    if (message.withdrawProposalAmount !== undefined) {
      Coin.encode(message.withdrawProposalAmount, writer.uint32(162).fork()).ldelim();
    }

    if (message.projectDueTime !== undefined) {
      Timestamp.encode(toTimestamp(message.projectDueTime), writer.uint32(170).fork()).ldelim();
    }

    for (const v of message.withdrawAccounts) {
      writer.uint32(178).bytes(v!);
    }

    for (const v of message.transferAccounts) {
      writer.uint32(186).bytes(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolInfo();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;

        case 2:
          message.poolName = reader.string();
          break;

        case 3:
          message.linkedProject = reader.int32();
          break;

        case 4:
          message.ownerAddress = reader.bytes();
          break;

        case 5:
          message.apy = reader.string();
          break;

        case 6:
          message.totalAmount = Coin.decode(reader, reader.uint32());
          break;

        case 7:
          message.payFreq = reader.int32();
          break;

        case 8:
          message.reserveFactor = reader.string();
          break;

        case 9:
          message.poolNFTIds.push(reader.string());
          break;

        case 10:
          message.lastPaymentTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 11:
          message.poolStatus = (reader.int32() as any);
          break;

        case 12:
          message.borrowedAmount = Coin.decode(reader, reader.uint32());
          break;

        case 13:
          message.poolInterest = reader.string();
          break;

        case 14:
          message.projectLength = (reader.uint64() as Long);
          break;

        case 15:
          message.usableAmount = Coin.decode(reader, reader.uint32());
          break;

        case 16:
          message.targetAmount = Coin.decode(reader, reader.uint32());
          break;

        case 17:
          message.poolType = (reader.int32() as any);
          break;

        case 18:
          message.escrowInterestAmount = reader.string();
          break;

        case 19:
          message.escrowPrincipalAmount = Coin.decode(reader, reader.uint32());
          break;

        case 20:
          message.withdrawProposalAmount = Coin.decode(reader, reader.uint32());
          break;

        case 21:
          message.projectDueTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;

        case 22:
          message.withdrawAccounts.push(reader.bytes());
          break;

        case 23:
          message.transferAccounts.push(reader.bytes());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<PoolInfo>): PoolInfo {
    const message = createBasePoolInfo();
    message.index = object.index ?? "";
    message.poolName = object.poolName ?? "";
    message.linkedProject = object.linkedProject ?? 0;
    message.ownerAddress = object.ownerAddress ?? new Uint8Array();
    message.apy = object.apy ?? "";
    message.totalAmount = object.totalAmount !== undefined && object.totalAmount !== null ? Coin.fromPartial(object.totalAmount) : undefined;
    message.payFreq = object.payFreq ?? 0;
    message.reserveFactor = object.reserveFactor ?? "";
    message.poolNFTIds = object.poolNFTIds?.map(e => e) || [];
    message.lastPaymentTime = object.lastPaymentTime ?? undefined;
    message.poolStatus = object.poolStatus ?? 0;
    message.borrowedAmount = object.borrowedAmount !== undefined && object.borrowedAmount !== null ? Coin.fromPartial(object.borrowedAmount) : undefined;
    message.poolInterest = object.poolInterest ?? "";
    message.projectLength = object.projectLength !== undefined && object.projectLength !== null ? Long.fromValue(object.projectLength) : Long.UZERO;
    message.usableAmount = object.usableAmount !== undefined && object.usableAmount !== null ? Coin.fromPartial(object.usableAmount) : undefined;
    message.targetAmount = object.targetAmount !== undefined && object.targetAmount !== null ? Coin.fromPartial(object.targetAmount) : undefined;
    message.poolType = object.poolType ?? 0;
    message.escrowInterestAmount = object.escrowInterestAmount ?? "";
    message.escrowPrincipalAmount = object.escrowPrincipalAmount !== undefined && object.escrowPrincipalAmount !== null ? Coin.fromPartial(object.escrowPrincipalAmount) : undefined;
    message.withdrawProposalAmount = object.withdrawProposalAmount !== undefined && object.withdrawProposalAmount !== null ? Coin.fromPartial(object.withdrawProposalAmount) : undefined;
    message.projectDueTime = object.projectDueTime ?? undefined;
    message.withdrawAccounts = object.withdrawAccounts?.map(e => e) || [];
    message.transferAccounts = object.transferAccounts?.map(e => e) || [];
    return message;
  }

};

function createBasePoolWithInvestors(): PoolWithInvestors {
  return {
    poolIndex: "",
    investors: []
  };
}

export const PoolWithInvestors = {
  encode(message: PoolWithInvestors, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }

    for (const v of message.investors) {
      writer.uint32(18).string(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolWithInvestors {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolWithInvestors();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.poolIndex = reader.string();
          break;

        case 2:
          message.investors.push(reader.string());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<PoolWithInvestors>): PoolWithInvestors {
    const message = createBasePoolWithInvestors();
    message.poolIndex = object.poolIndex ?? "";
    message.investors = object.investors?.map(e => e) || [];
    return message;
  }

};

function createBasePoolDepositedInvestors(): PoolDepositedInvestors {
  return {
    poolIndex: "",
    walletAddress: []
  };
}

export const PoolDepositedInvestors = {
  encode(message: PoolDepositedInvestors, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.poolIndex !== "") {
      writer.uint32(10).string(message.poolIndex);
    }

    for (const v of message.walletAddress) {
      writer.uint32(34).bytes(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PoolDepositedInvestors {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePoolDepositedInvestors();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.poolIndex = reader.string();
          break;

        case 4:
          message.walletAddress.push(reader.bytes());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<PoolDepositedInvestors>): PoolDepositedInvestors {
    const message = createBasePoolDepositedInvestors();
    message.poolIndex = object.poolIndex ?? "";
    message.walletAddress = object.walletAddress?.map(e => e) || [];
    return message;
  }

};

function createBaseWalletsLinkPool(): WalletsLinkPool {
  return {
    walletAddress: new Uint8Array(),
    poolsIndex: []
  };
}

export const WalletsLinkPool = {
  encode(message: WalletsLinkPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.walletAddress.length !== 0) {
      writer.uint32(10).bytes(message.walletAddress);
    }

    for (const v of message.poolsIndex) {
      writer.uint32(18).string(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WalletsLinkPool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWalletsLinkPool();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.walletAddress = reader.bytes();
          break;

        case 2:
          message.poolsIndex.push(reader.string());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<WalletsLinkPool>): WalletsLinkPool {
    const message = createBaseWalletsLinkPool();
    message.walletAddress = object.walletAddress ?? new Uint8Array();
    message.poolsIndex = object.poolsIndex?.map(e => e) || [];
    return message;
  }

};