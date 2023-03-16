import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
export interface MsgCreatePool {
  creator: string;
  projectIndex: number;
  poolName: string;
  apy: string;
  targetTokenAmount?: Coin;
}
export interface MsgCreatePoolSDKType {
  creator: string;
  project_index: number;
  pool_name: string;
  apy: string;
  target_token_amount?: CoinSDKType;
}
export interface MsgCreatePoolResponse {
  poolIndex: string[];
}
export interface MsgCreatePoolResponseSDKType {
  pool_index: string[];
}
export interface MsgAddInvestors {
  creator: string;
  poolIndex: string;
  investorID: string[];
}
export interface MsgAddInvestorsSDKType {
  creator: string;
  pool_index: string;
  investor_iD: string[];
}
export interface MsgAddInvestorsResponse {
  operationResult: boolean;
}
export interface MsgAddInvestorsResponseSDKType {
  operation_result: boolean;
}
export interface MsgDeposit {
  creator: string;
  poolIndex: string;
  token?: Coin;
}
export interface MsgDepositSDKType {
  creator: string;
  pool_index: string;
  token?: CoinSDKType;
}
export interface MsgDepositResponse {}
export interface MsgDepositResponseSDKType {}
export interface MsgBorrow {
  creator: string;
  poolIndex: string;
  borrowAmount?: Coin;
}
export interface MsgBorrowSDKType {
  creator: string;
  pool_index: string;
  borrow_amount?: CoinSDKType;
}
export interface MsgBorrowResponse {
  borrowAmount: string;
}
export interface MsgBorrowResponseSDKType {
  borrow_amount: string;
}
export interface MsgRepayInterest {
  creator: string;
  poolIndex: string;
  token?: Coin;
}
export interface MsgRepayInterestSDKType {
  creator: string;
  pool_index: string;
  token?: CoinSDKType;
}
export interface MsgRepayInterestResponse {}
export interface MsgRepayInterestResponseSDKType {}
export interface MsgClaimInterest {
  creator: string;
  poolIndex: string;
}
export interface MsgClaimInterestSDKType {
  creator: string;
  pool_index: string;
}
export interface MsgClaimInterestResponse {
  amount: string;
}
export interface MsgClaimInterestResponseSDKType {
  amount: string;
}
export interface MsgUpdatePool {
  creator: string;
  poolIndex: string;
  poolName: string;
  poolApy: string;
  targetTokenAmount?: Coin;
}
export interface MsgUpdatePoolSDKType {
  creator: string;
  pool_index: string;
  pool_name: string;
  pool_apy: string;
  target_token_amount?: CoinSDKType;
}
export interface MsgUpdatePoolResponse {}
export interface MsgUpdatePoolResponseSDKType {}
export interface MsgActivePool {
  creator: string;
  poolIndex: string;
}
export interface MsgActivePoolSDKType {
  creator: string;
  pool_index: string;
}
export interface MsgActivePoolResponse {}
export interface MsgActivePoolResponseSDKType {}
export interface MsgPayPrincipal {
  creator: string;
  poolIndex: string;
  token?: Coin;
}
export interface MsgPayPrincipalSDKType {
  creator: string;
  pool_index: string;
  token?: CoinSDKType;
}
export interface MsgPayPrincipalResponse {}
export interface MsgPayPrincipalResponseSDKType {}
export interface MsgWithdrawPrincipal {
  creator: string;
  poolIndex: string;
  token?: Coin;
}
export interface MsgWithdrawPrincipalSDKType {
  creator: string;
  pool_index: string;
  token?: CoinSDKType;
}
export interface MsgWithdrawPrincipalResponse {
  amount: string;
}
export interface MsgWithdrawPrincipalResponseSDKType {
  amount: string;
}
export interface MsgSubmitWithdrawProposal {
  creator: string;
  poolIndex: string;
}
export interface MsgSubmitWithdrawProposalSDKType {
  creator: string;
  pool_index: string;
}
export interface MsgSubmitWithdrawProposalResponse {
  operationResult: boolean;
}
export interface MsgSubmitWithdrawProposalResponseSDKType {
  operation_result: boolean;
}
export interface MsgTransferOwnership {
  creator: string;
  poolIndex: string;
}
export interface MsgTransferOwnershipSDKType {
  creator: string;
  pool_index: string;
}
export interface MsgTransferOwnershipResponse {
  operationResult: boolean;
}
export interface MsgTransferOwnershipResponseSDKType {
  operation_result: boolean;
}

function createBaseMsgCreatePool(): MsgCreatePool {
  return {
    creator: "",
    projectIndex: 0,
    poolName: "",
    apy: "",
    targetTokenAmount: undefined
  };
}

export const MsgCreatePool = {
  encode(message: MsgCreatePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.projectIndex !== 0) {
      writer.uint32(16).int32(message.projectIndex);
    }

    if (message.poolName !== "") {
      writer.uint32(26).string(message.poolName);
    }

    if (message.apy !== "") {
      writer.uint32(34).string(message.apy);
    }

    if (message.targetTokenAmount !== undefined) {
      Coin.encode(message.targetTokenAmount, writer.uint32(42).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreatePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreatePool();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.projectIndex = reader.int32();
          break;

        case 3:
          message.poolName = reader.string();
          break;

        case 4:
          message.apy = reader.string();
          break;

        case 5:
          message.targetTokenAmount = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCreatePool>): MsgCreatePool {
    const message = createBaseMsgCreatePool();
    message.creator = object.creator ?? "";
    message.projectIndex = object.projectIndex ?? 0;
    message.poolName = object.poolName ?? "";
    message.apy = object.apy ?? "";
    message.targetTokenAmount = object.targetTokenAmount !== undefined && object.targetTokenAmount !== null ? Coin.fromPartial(object.targetTokenAmount) : undefined;
    return message;
  }

};

function createBaseMsgCreatePoolResponse(): MsgCreatePoolResponse {
  return {
    poolIndex: []
  };
}

export const MsgCreatePoolResponse = {
  encode(message: MsgCreatePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.poolIndex) {
      writer.uint32(10).string(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreatePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreatePoolResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.poolIndex.push(reader.string());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCreatePoolResponse>): MsgCreatePoolResponse {
    const message = createBaseMsgCreatePoolResponse();
    message.poolIndex = object.poolIndex?.map(e => e) || [];
    return message;
  }

};

function createBaseMsgAddInvestors(): MsgAddInvestors {
  return {
    creator: "",
    poolIndex: "",
    investorID: []
  };
}

export const MsgAddInvestors = {
  encode(message: MsgAddInvestors, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    for (const v of message.investorID) {
      writer.uint32(26).string(v!);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddInvestors {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddInvestors();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        case 3:
          message.investorID.push(reader.string());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgAddInvestors>): MsgAddInvestors {
    const message = createBaseMsgAddInvestors();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.investorID = object.investorID?.map(e => e) || [];
    return message;
  }

};

function createBaseMsgAddInvestorsResponse(): MsgAddInvestorsResponse {
  return {
    operationResult: false
  };
}

export const MsgAddInvestorsResponse = {
  encode(message: MsgAddInvestorsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.operationResult === true) {
      writer.uint32(8).bool(message.operationResult);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddInvestorsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddInvestorsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.operationResult = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgAddInvestorsResponse>): MsgAddInvestorsResponse {
    const message = createBaseMsgAddInvestorsResponse();
    message.operationResult = object.operationResult ?? false;
    return message;
  }

};

function createBaseMsgDeposit(): MsgDeposit {
  return {
    creator: "",
    poolIndex: "",
    token: undefined
  };
}

export const MsgDeposit = {
  encode(message: MsgDeposit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeposit {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeposit();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        case 3:
          message.token = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgDeposit>): MsgDeposit {
    const message = createBaseMsgDeposit();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.token = object.token !== undefined && object.token !== null ? Coin.fromPartial(object.token) : undefined;
    return message;
  }

};

function createBaseMsgDepositResponse(): MsgDepositResponse {
  return {};
}

export const MsgDepositResponse = {
  encode(_: MsgDepositResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDepositResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDepositResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgDepositResponse>): MsgDepositResponse {
    const message = createBaseMsgDepositResponse();
    return message;
  }

};

function createBaseMsgBorrow(): MsgBorrow {
  return {
    creator: "",
    poolIndex: "",
    borrowAmount: undefined
  };
}

export const MsgBorrow = {
  encode(message: MsgBorrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    if (message.borrowAmount !== undefined) {
      Coin.encode(message.borrowAmount, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBorrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBorrow();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        case 3:
          message.borrowAmount = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgBorrow>): MsgBorrow {
    const message = createBaseMsgBorrow();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.borrowAmount = object.borrowAmount !== undefined && object.borrowAmount !== null ? Coin.fromPartial(object.borrowAmount) : undefined;
    return message;
  }

};

function createBaseMsgBorrowResponse(): MsgBorrowResponse {
  return {
    borrowAmount: ""
  };
}

export const MsgBorrowResponse = {
  encode(message: MsgBorrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.borrowAmount !== "") {
      writer.uint32(10).string(message.borrowAmount);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBorrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBorrowResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.borrowAmount = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgBorrowResponse>): MsgBorrowResponse {
    const message = createBaseMsgBorrowResponse();
    message.borrowAmount = object.borrowAmount ?? "";
    return message;
  }

};

function createBaseMsgRepayInterest(): MsgRepayInterest {
  return {
    creator: "",
    poolIndex: "",
    token: undefined
  };
}

export const MsgRepayInterest = {
  encode(message: MsgRepayInterest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(34).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRepayInterest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRepayInterest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        case 4:
          message.token = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgRepayInterest>): MsgRepayInterest {
    const message = createBaseMsgRepayInterest();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.token = object.token !== undefined && object.token !== null ? Coin.fromPartial(object.token) : undefined;
    return message;
  }

};

function createBaseMsgRepayInterestResponse(): MsgRepayInterestResponse {
  return {};
}

export const MsgRepayInterestResponse = {
  encode(_: MsgRepayInterestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRepayInterestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRepayInterestResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgRepayInterestResponse>): MsgRepayInterestResponse {
    const message = createBaseMsgRepayInterestResponse();
    return message;
  }

};

function createBaseMsgClaimInterest(): MsgClaimInterest {
  return {
    creator: "",
    poolIndex: ""
  };
}

export const MsgClaimInterest = {
  encode(message: MsgClaimInterest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimInterest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimInterest();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgClaimInterest>): MsgClaimInterest {
    const message = createBaseMsgClaimInterest();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  }

};

function createBaseMsgClaimInterestResponse(): MsgClaimInterestResponse {
  return {
    amount: ""
  };
}

export const MsgClaimInterestResponse = {
  encode(message: MsgClaimInterestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimInterestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimInterestResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgClaimInterestResponse>): MsgClaimInterestResponse {
    const message = createBaseMsgClaimInterestResponse();
    message.amount = object.amount ?? "";
    return message;
  }

};

function createBaseMsgUpdatePool(): MsgUpdatePool {
  return {
    creator: "",
    poolIndex: "",
    poolName: "",
    poolApy: "",
    targetTokenAmount: undefined
  };
}

export const MsgUpdatePool = {
  encode(message: MsgUpdatePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    if (message.poolName !== "") {
      writer.uint32(26).string(message.poolName);
    }

    if (message.poolApy !== "") {
      writer.uint32(34).string(message.poolApy);
    }

    if (message.targetTokenAmount !== undefined) {
      Coin.encode(message.targetTokenAmount, writer.uint32(50).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdatePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdatePool();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        case 3:
          message.poolName = reader.string();
          break;

        case 4:
          message.poolApy = reader.string();
          break;

        case 6:
          message.targetTokenAmount = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgUpdatePool>): MsgUpdatePool {
    const message = createBaseMsgUpdatePool();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.poolName = object.poolName ?? "";
    message.poolApy = object.poolApy ?? "";
    message.targetTokenAmount = object.targetTokenAmount !== undefined && object.targetTokenAmount !== null ? Coin.fromPartial(object.targetTokenAmount) : undefined;
    return message;
  }

};

function createBaseMsgUpdatePoolResponse(): MsgUpdatePoolResponse {
  return {};
}

export const MsgUpdatePoolResponse = {
  encode(_: MsgUpdatePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdatePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdatePoolResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgUpdatePoolResponse>): MsgUpdatePoolResponse {
    const message = createBaseMsgUpdatePoolResponse();
    return message;
  }

};

function createBaseMsgActivePool(): MsgActivePool {
  return {
    creator: "",
    poolIndex: ""
  };
}

export const MsgActivePool = {
  encode(message: MsgActivePool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgActivePool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgActivePool();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgActivePool>): MsgActivePool {
    const message = createBaseMsgActivePool();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  }

};

function createBaseMsgActivePoolResponse(): MsgActivePoolResponse {
  return {};
}

export const MsgActivePoolResponse = {
  encode(_: MsgActivePoolResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgActivePoolResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgActivePoolResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgActivePoolResponse>): MsgActivePoolResponse {
    const message = createBaseMsgActivePoolResponse();
    return message;
  }

};

function createBaseMsgPayPrincipal(): MsgPayPrincipal {
  return {
    creator: "",
    poolIndex: "",
    token: undefined
  };
}

export const MsgPayPrincipal = {
  encode(message: MsgPayPrincipal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPayPrincipal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPayPrincipal();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        case 3:
          message.token = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgPayPrincipal>): MsgPayPrincipal {
    const message = createBaseMsgPayPrincipal();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.token = object.token !== undefined && object.token !== null ? Coin.fromPartial(object.token) : undefined;
    return message;
  }

};

function createBaseMsgPayPrincipalResponse(): MsgPayPrincipalResponse {
  return {};
}

export const MsgPayPrincipalResponse = {
  encode(_: MsgPayPrincipalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPayPrincipalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPayPrincipalResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgPayPrincipalResponse>): MsgPayPrincipalResponse {
    const message = createBaseMsgPayPrincipalResponse();
    return message;
  }

};

function createBaseMsgWithdrawPrincipal(): MsgWithdrawPrincipal {
  return {
    creator: "",
    poolIndex: "",
    token: undefined
  };
}

export const MsgWithdrawPrincipal = {
  encode(message: MsgWithdrawPrincipal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    if (message.token !== undefined) {
      Coin.encode(message.token, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawPrincipal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawPrincipal();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        case 3:
          message.token = Coin.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgWithdrawPrincipal>): MsgWithdrawPrincipal {
    const message = createBaseMsgWithdrawPrincipal();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    message.token = object.token !== undefined && object.token !== null ? Coin.fromPartial(object.token) : undefined;
    return message;
  }

};

function createBaseMsgWithdrawPrincipalResponse(): MsgWithdrawPrincipalResponse {
  return {
    amount: ""
  };
}

export const MsgWithdrawPrincipalResponse = {
  encode(message: MsgWithdrawPrincipalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgWithdrawPrincipalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawPrincipalResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgWithdrawPrincipalResponse>): MsgWithdrawPrincipalResponse {
    const message = createBaseMsgWithdrawPrincipalResponse();
    message.amount = object.amount ?? "";
    return message;
  }

};

function createBaseMsgSubmitWithdrawProposal(): MsgSubmitWithdrawProposal {
  return {
    creator: "",
    poolIndex: ""
  };
}

export const MsgSubmitWithdrawProposal = {
  encode(message: MsgSubmitWithdrawProposal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitWithdrawProposal {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitWithdrawProposal();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgSubmitWithdrawProposal>): MsgSubmitWithdrawProposal {
    const message = createBaseMsgSubmitWithdrawProposal();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  }

};

function createBaseMsgSubmitWithdrawProposalResponse(): MsgSubmitWithdrawProposalResponse {
  return {
    operationResult: false
  };
}

export const MsgSubmitWithdrawProposalResponse = {
  encode(message: MsgSubmitWithdrawProposalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.operationResult === true) {
      writer.uint32(8).bool(message.operationResult);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitWithdrawProposalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitWithdrawProposalResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.operationResult = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgSubmitWithdrawProposalResponse>): MsgSubmitWithdrawProposalResponse {
    const message = createBaseMsgSubmitWithdrawProposalResponse();
    message.operationResult = object.operationResult ?? false;
    return message;
  }

};

function createBaseMsgTransferOwnership(): MsgTransferOwnership {
  return {
    creator: "",
    poolIndex: ""
  };
}

export const MsgTransferOwnership = {
  encode(message: MsgTransferOwnership, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }

    if (message.poolIndex !== "") {
      writer.uint32(18).string(message.poolIndex);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferOwnership {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferOwnership();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;

        case 2:
          message.poolIndex = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgTransferOwnership>): MsgTransferOwnership {
    const message = createBaseMsgTransferOwnership();
    message.creator = object.creator ?? "";
    message.poolIndex = object.poolIndex ?? "";
    return message;
  }

};

function createBaseMsgTransferOwnershipResponse(): MsgTransferOwnershipResponse {
  return {
    operationResult: false
  };
}

export const MsgTransferOwnershipResponse = {
  encode(message: MsgTransferOwnershipResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.operationResult === true) {
      writer.uint32(8).bool(message.operationResult);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferOwnershipResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferOwnershipResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.operationResult = reader.bool();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgTransferOwnershipResponse>): MsgTransferOwnershipResponse {
    const message = createBaseMsgTransferOwnershipResponse();
    message.operationResult = object.operationResult ?? false;
    return message;
  }

};