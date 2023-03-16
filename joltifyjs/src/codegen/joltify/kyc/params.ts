import { Coin, CoinSDKType } from "../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial, Long } from "../../helpers";
export interface BasicInfo {
  description: string;
  projectsUrl: string;
  projectCountry: string;
  businessNumber: string;
  reserved: Uint8Array;
  projectName: string;
}
export interface BasicInfoSDKType {
  description: string;
  projects_url: string;
  project_country: string;
  business_number: string;
  reserved: Uint8Array;
  project_name: string;
}
/** Market defines an asset in the pricefeed. */

export interface ProjectInfo {
  index: number;
  SPVName: string;
  basicInfo?: BasicInfo;
  projectOwner: Uint8Array;
  projectLength: Long;
  projectTargetAmount?: Coin;
  baseApy: string;
  payFreq: string;
}
/** Market defines an asset in the pricefeed. */

export interface ProjectInfoSDKType {
  index: number;
  SPV_name: string;
  basic_info?: BasicInfoSDKType;
  project_owner: Uint8Array;
  project_length: Long;
  project_target_amount?: CoinSDKType;
  base_apy: string;
  pay_freq: string;
}
/** Params defines the parameters for the module. */

export interface Params {
  projectsInfo: ProjectInfo[];
  submitter: Uint8Array[];
}
/** Params defines the parameters for the module. */

export interface ParamsSDKType {
  projects_info: ProjectInfoSDKType[];
  submitter: Uint8Array[];
}

function createBaseBasicInfo(): BasicInfo {
  return {
    description: "",
    projectsUrl: "",
    projectCountry: "",
    businessNumber: "",
    reserved: new Uint8Array(),
    projectName: ""
  };
}

export const BasicInfo = {
  encode(message: BasicInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.description !== "") {
      writer.uint32(10).string(message.description);
    }

    if (message.projectsUrl !== "") {
      writer.uint32(18).string(message.projectsUrl);
    }

    if (message.projectCountry !== "") {
      writer.uint32(26).string(message.projectCountry);
    }

    if (message.businessNumber !== "") {
      writer.uint32(34).string(message.businessNumber);
    }

    if (message.reserved.length !== 0) {
      writer.uint32(42).bytes(message.reserved);
    }

    if (message.projectName !== "") {
      writer.uint32(50).string(message.projectName);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BasicInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBasicInfo();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.description = reader.string();
          break;

        case 2:
          message.projectsUrl = reader.string();
          break;

        case 3:
          message.projectCountry = reader.string();
          break;

        case 4:
          message.businessNumber = reader.string();
          break;

        case 5:
          message.reserved = reader.bytes();
          break;

        case 6:
          message.projectName = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<BasicInfo>): BasicInfo {
    const message = createBaseBasicInfo();
    message.description = object.description ?? "";
    message.projectsUrl = object.projectsUrl ?? "";
    message.projectCountry = object.projectCountry ?? "";
    message.businessNumber = object.businessNumber ?? "";
    message.reserved = object.reserved ?? new Uint8Array();
    message.projectName = object.projectName ?? "";
    return message;
  }

};

function createBaseProjectInfo(): ProjectInfo {
  return {
    index: 0,
    SPVName: "",
    basicInfo: undefined,
    projectOwner: new Uint8Array(),
    projectLength: Long.UZERO,
    projectTargetAmount: undefined,
    baseApy: "",
    payFreq: ""
  };
}

export const ProjectInfo = {
  encode(message: ProjectInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== 0) {
      writer.uint32(8).int32(message.index);
    }

    if (message.SPVName !== "") {
      writer.uint32(18).string(message.SPVName);
    }

    if (message.basicInfo !== undefined) {
      BasicInfo.encode(message.basicInfo, writer.uint32(26).fork()).ldelim();
    }

    if (message.projectOwner.length !== 0) {
      writer.uint32(34).bytes(message.projectOwner);
    }

    if (!message.projectLength.isZero()) {
      writer.uint32(40).uint64(message.projectLength);
    }

    if (message.projectTargetAmount !== undefined) {
      Coin.encode(message.projectTargetAmount, writer.uint32(50).fork()).ldelim();
    }

    if (message.baseApy !== "") {
      writer.uint32(58).string(message.baseApy);
    }

    if (message.payFreq !== "") {
      writer.uint32(66).string(message.payFreq);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProjectInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProjectInfo();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.index = reader.int32();
          break;

        case 2:
          message.SPVName = reader.string();
          break;

        case 3:
          message.basicInfo = BasicInfo.decode(reader, reader.uint32());
          break;

        case 4:
          message.projectOwner = reader.bytes();
          break;

        case 5:
          message.projectLength = (reader.uint64() as Long);
          break;

        case 6:
          message.projectTargetAmount = Coin.decode(reader, reader.uint32());
          break;

        case 7:
          message.baseApy = reader.string();
          break;

        case 8:
          message.payFreq = reader.string();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<ProjectInfo>): ProjectInfo {
    const message = createBaseProjectInfo();
    message.index = object.index ?? 0;
    message.SPVName = object.SPVName ?? "";
    message.basicInfo = object.basicInfo !== undefined && object.basicInfo !== null ? BasicInfo.fromPartial(object.basicInfo) : undefined;
    message.projectOwner = object.projectOwner ?? new Uint8Array();
    message.projectLength = object.projectLength !== undefined && object.projectLength !== null ? Long.fromValue(object.projectLength) : Long.UZERO;
    message.projectTargetAmount = object.projectTargetAmount !== undefined && object.projectTargetAmount !== null ? Coin.fromPartial(object.projectTargetAmount) : undefined;
    message.baseApy = object.baseApy ?? "";
    message.payFreq = object.payFreq ?? "";
    return message;
  }

};

function createBaseParams(): Params {
  return {
    projectsInfo: [],
    submitter: []
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.projectsInfo) {
      ProjectInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    for (const v of message.submitter) {
      writer.uint32(34).bytes(v!);
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
          message.projectsInfo.push(ProjectInfo.decode(reader, reader.uint32()));
          break;

        case 4:
          message.submitter.push(reader.bytes());
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
    message.projectsInfo = object.projectsInfo?.map(e => ProjectInfo.fromPartial(e)) || [];
    message.submitter = object.submitter?.map(e => e) || [];
    return message;
  }

};