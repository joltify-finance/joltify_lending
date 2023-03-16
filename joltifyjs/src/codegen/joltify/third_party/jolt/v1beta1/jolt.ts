import { Coin, CoinSDKType } from "../../../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../../../helpers";
/** Params defines the parameters for the jolt module. */

export interface Params {
  moneyMarkets: MoneyMarket[];
  minimumBorrowUsdValue: string;
  surplusAuctionThreshold: string;
}
/** Params defines the parameters for the jolt module. */

export interface ParamsSDKType {
  money_markets: MoneyMarketSDKType[];
  minimum_borrow_usd_value: string;
  surplus_auction_threshold: string;
}
/** MoneyMarket is a money market for an individual asset. */

export interface MoneyMarket {
  denom: string;
  borrowLimit?: BorrowLimit;
  spotMarketId: string;
  conversionFactor: string;
  interestRateModel?: InterestRateModel;
  reserveFactor: string;
  keeperRewardPercentage: string;
}
/** MoneyMarket is a money market for an individual asset. */

export interface MoneyMarketSDKType {
  denom: string;
  borrow_limit?: BorrowLimitSDKType;
  spot_market_id: string;
  conversion_factor: string;
  interest_rate_model?: InterestRateModelSDKType;
  reserve_factor: string;
  keeper_reward_percentage: string;
}
/** BorrowLimit enforces restrictions on a money market. */

export interface BorrowLimit {
  hasMaxLimit: boolean;
  maximumLimit: string;
  loanToValue: string;
}
/** BorrowLimit enforces restrictions on a money market. */

export interface BorrowLimitSDKType {
  has_max_limit: boolean;
  maximum_limit: string;
  loan_to_value: string;
}
/** InterestRateModel contains information about an asset's interest rate. */

export interface InterestRateModel {
  baseRateApy: string;
  baseMultiplier: string;
  kink: string;
  jumpMultiplier: string;
}
/** InterestRateModel contains information about an asset's interest rate. */

export interface InterestRateModelSDKType {
  base_rate_apy: string;
  base_multiplier: string;
  kink: string;
  jump_multiplier: string;
}
/** Deposit defines an amount of coins deposited into a jolt module account. */

export interface Deposit {
  depositor: string;
  amount: Coin[];
  index: SupplyInterestFactor[];
}
/** Deposit defines an amount of coins deposited into a jolt module account. */

export interface DepositSDKType {
  depositor: string;
  amount: CoinSDKType[];
  index: SupplyInterestFactorSDKType[];
}
/** Borrow defines an amount of coins borrowed from a jolt module account. */

export interface Borrow {
  borrower: string;
  amount: Coin[];
  index: BorrowInterestFactor[];
}
/** Borrow defines an amount of coins borrowed from a jolt module account. */

export interface BorrowSDKType {
  borrower: string;
  amount: CoinSDKType[];
  index: BorrowInterestFactorSDKType[];
}
/** SupplyInterestFactor defines an individual borrow interest factor. */

export interface SupplyInterestFactor {
  denom: string;
  value: string;
}
/** SupplyInterestFactor defines an individual borrow interest factor. */

export interface SupplyInterestFactorSDKType {
  denom: string;
  value: string;
}
/** BorrowInterestFactor defines an individual borrow interest factor. */

export interface BorrowInterestFactor {
  denom: string;
  value: string;
}
/** BorrowInterestFactor defines an individual borrow interest factor. */

export interface BorrowInterestFactorSDKType {
  denom: string;
  value: string;
}
/** CoinsProto defines a Protobuf wrapper around a Coins slice */

export interface CoinsProto {
  coins: Coin[];
}
/** CoinsProto defines a Protobuf wrapper around a Coins slice */

export interface CoinsProtoSDKType {
  coins: CoinSDKType[];
}

function createBaseParams(): Params {
  return {
    moneyMarkets: [],
    minimumBorrowUsdValue: "",
    surplusAuctionThreshold: ""
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.moneyMarkets) {
      MoneyMarket.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    if (message.minimumBorrowUsdValue !== "") {
      writer.uint32(18).string(message.minimumBorrowUsdValue);
    }

    if (message.surplusAuctionThreshold !== "") {
      writer.uint32(26).string(message.surplusAuctionThreshold);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.moneyMarkets.push(MoneyMarket.decode(reader, reader.uint32()));
          break;

        case 2:
          message.minimumBorrowUsdValue = reader.string();
          break;

        case 3:
          message.surplusAuctionThreshold = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = createBaseParams();
    message.moneyMarkets = object.moneyMarkets?.map(e => MoneyMarket.fromPartial(e)) || [];
    message.minimumBorrowUsdValue = object.minimumBorrowUsdValue ?? "";
    message.surplusAuctionThreshold = object.surplusAuctionThreshold ?? "";
    return message;
  }

};

function createBaseMoneyMarket(): MoneyMarket {
  return {
    denom: "",
    borrowLimit: undefined,
    spotMarketId: "",
    conversionFactor: "",
    interestRateModel: undefined,
    reserveFactor: "",
    keeperRewardPercentage: ""
  };
}

export const MoneyMarket = {
  encode(message: MoneyMarket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.borrowLimit !== undefined) {
      BorrowLimit.encode(message.borrowLimit, writer.uint32(18).fork()).ldelim();
    }

    if (message.spotMarketId !== "") {
      writer.uint32(26).string(message.spotMarketId);
    }

    if (message.conversionFactor !== "") {
      writer.uint32(34).string(message.conversionFactor);
    }

    if (message.interestRateModel !== undefined) {
      InterestRateModel.encode(message.interestRateModel, writer.uint32(42).fork()).ldelim();
    }

    if (message.reserveFactor !== "") {
      writer.uint32(50).string(message.reserveFactor);
    }

    if (message.keeperRewardPercentage !== "") {
      writer.uint32(58).string(message.keeperRewardPercentage);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MoneyMarket {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMoneyMarket();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.borrowLimit = BorrowLimit.decode(reader, reader.uint32());
          break;

        case 3:
          message.spotMarketId = reader.string();
          break;

        case 4:
          message.conversionFactor = reader.string();
          break;

        case 5:
          message.interestRateModel = InterestRateModel.decode(reader, reader.uint32());
          break;

        case 6:
          message.reserveFactor = reader.string();
          break;

        case 7:
          message.keeperRewardPercentage = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MoneyMarket>): MoneyMarket {
    const message = createBaseMoneyMarket();
    message.denom = object.denom ?? "";
    message.borrowLimit = object.borrowLimit !== undefined && object.borrowLimit !== null ? BorrowLimit.fromPartial(object.borrowLimit) : undefined;
    message.spotMarketId = object.spotMarketId ?? "";
    message.conversionFactor = object.conversionFactor ?? "";
    message.interestRateModel = object.interestRateModel !== undefined && object.interestRateModel !== null ? InterestRateModel.fromPartial(object.interestRateModel) : undefined;
    message.reserveFactor = object.reserveFactor ?? "";
    message.keeperRewardPercentage = object.keeperRewardPercentage ?? "";
    return message;
  }

};

function createBaseBorrowLimit(): BorrowLimit {
  return {
    hasMaxLimit: false,
    maximumLimit: "",
    loanToValue: ""
  };
}

export const BorrowLimit = {
  encode(message: BorrowLimit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.hasMaxLimit === true) {
      writer.uint32(8).bool(message.hasMaxLimit);
    }

    if (message.maximumLimit !== "") {
      writer.uint32(18).string(message.maximumLimit);
    }

    if (message.loanToValue !== "") {
      writer.uint32(26).string(message.loanToValue);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BorrowLimit {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBorrowLimit();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.hasMaxLimit = reader.bool();
          break;

        case 2:
          message.maximumLimit = reader.string();
          break;

        case 3:
          message.loanToValue = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BorrowLimit>): BorrowLimit {
    const message = createBaseBorrowLimit();
    message.hasMaxLimit = object.hasMaxLimit ?? false;
    message.maximumLimit = object.maximumLimit ?? "";
    message.loanToValue = object.loanToValue ?? "";
    return message;
  }

};

function createBaseInterestRateModel(): InterestRateModel {
  return {
    baseRateApy: "",
    baseMultiplier: "",
    kink: "",
    jumpMultiplier: ""
  };
}

export const InterestRateModel = {
  encode(message: InterestRateModel, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseRateApy !== "") {
      writer.uint32(10).string(message.baseRateApy);
    }

    if (message.baseMultiplier !== "") {
      writer.uint32(18).string(message.baseMultiplier);
    }

    if (message.kink !== "") {
      writer.uint32(26).string(message.kink);
    }

    if (message.jumpMultiplier !== "") {
      writer.uint32(34).string(message.jumpMultiplier);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InterestRateModel {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInterestRateModel();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.baseRateApy = reader.string();
          break;

        case 2:
          message.baseMultiplier = reader.string();
          break;

        case 3:
          message.kink = reader.string();
          break;

        case 4:
          message.jumpMultiplier = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<InterestRateModel>): InterestRateModel {
    const message = createBaseInterestRateModel();
    message.baseRateApy = object.baseRateApy ?? "";
    message.baseMultiplier = object.baseMultiplier ?? "";
    message.kink = object.kink ?? "";
    message.jumpMultiplier = object.jumpMultiplier ?? "";
    return message;
  }

};

function createBaseDeposit(): Deposit {
  return {
    depositor: "",
    amount: [],
    index: []
  };
}

export const Deposit = {
  encode(message: Deposit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.depositor !== "") {
      writer.uint32(10).string(message.depositor);
    }

    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.index) {
      SupplyInterestFactor.encode(v!, writer.uint32(26).fork()).ldelim();
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
          message.depositor = reader.string();
          break;

        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;

        case 3:
          message.index.push(SupplyInterestFactor.decode(reader, reader.uint32()));
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
    message.depositor = object.depositor ?? "";
    message.amount = object.amount?.map(e => Coin.fromPartial(e)) || [];
    message.index = object.index?.map(e => SupplyInterestFactor.fromPartial(e)) || [];
    return message;
  }

};

function createBaseBorrow(): Borrow {
  return {
    borrower: "",
    amount: [],
    index: []
  };
}

export const Borrow = {
  encode(message: Borrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.borrower !== "") {
      writer.uint32(10).string(message.borrower);
    }

    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }

    for (const v of message.index) {
      BorrowInterestFactor.encode(v!, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Borrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBorrow();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.borrower = reader.string();
          break;

        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;

        case 3:
          message.index.push(BorrowInterestFactor.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<Borrow>): Borrow {
    const message = createBaseBorrow();
    message.borrower = object.borrower ?? "";
    message.amount = object.amount?.map(e => Coin.fromPartial(e)) || [];
    message.index = object.index?.map(e => BorrowInterestFactor.fromPartial(e)) || [];
    return message;
  }

};

function createBaseSupplyInterestFactor(): SupplyInterestFactor {
  return {
    denom: "",
    value: ""
  };
}

export const SupplyInterestFactor = {
  encode(message: SupplyInterestFactor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SupplyInterestFactor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSupplyInterestFactor();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.value = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<SupplyInterestFactor>): SupplyInterestFactor {
    const message = createBaseSupplyInterestFactor();
    message.denom = object.denom ?? "";
    message.value = object.value ?? "";
    return message;
  }

};

function createBaseBorrowInterestFactor(): BorrowInterestFactor {
  return {
    denom: "",
    value: ""
  };
}

export const BorrowInterestFactor = {
  encode(message: BorrowInterestFactor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }

    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BorrowInterestFactor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBorrowInterestFactor();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;

        case 2:
          message.value = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BorrowInterestFactor>): BorrowInterestFactor {
    const message = createBaseBorrowInterestFactor();
    message.denom = object.denom ?? "";
    message.value = object.value ?? "";
    return message;
  }

};

function createBaseCoinsProto(): CoinsProto {
  return {
    coins: []
  };
}

export const CoinsProto = {
  encode(message: CoinsProto, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.coins) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CoinsProto {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCoinsProto();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.coins.push(Coin.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<CoinsProto>): CoinsProto {
    const message = createBaseCoinsProto();
    message.coins = object.coins?.map(e => Coin.fromPartial(e)) || [];
    return message;
  }

};