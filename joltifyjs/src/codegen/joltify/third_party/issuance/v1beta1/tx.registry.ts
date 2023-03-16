import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgIssueTokens, MsgRedeemTokens, MsgBlockAddress, MsgUnblockAddress, MsgSetPauseStatus } from "./tx";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/joltify.third_party.issuance.v1beta1.MsgIssueTokens", MsgIssueTokens], ["/joltify.third_party.issuance.v1beta1.MsgRedeemTokens", MsgRedeemTokens], ["/joltify.third_party.issuance.v1beta1.MsgBlockAddress", MsgBlockAddress], ["/joltify.third_party.issuance.v1beta1.MsgUnblockAddress", MsgUnblockAddress], ["/joltify.third_party.issuance.v1beta1.MsgSetPauseStatus", MsgSetPauseStatus]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    issueTokens(value: MsgIssueTokens) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgIssueTokens",
        value: MsgIssueTokens.encode(value).finish()
      };
    },

    redeemTokens(value: MsgRedeemTokens) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgRedeemTokens",
        value: MsgRedeemTokens.encode(value).finish()
      };
    },

    blockAddress(value: MsgBlockAddress) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgBlockAddress",
        value: MsgBlockAddress.encode(value).finish()
      };
    },

    unblockAddress(value: MsgUnblockAddress) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgUnblockAddress",
        value: MsgUnblockAddress.encode(value).finish()
      };
    },

    setPauseStatus(value: MsgSetPauseStatus) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgSetPauseStatus",
        value: MsgSetPauseStatus.encode(value).finish()
      };
    }

  },
  withTypeUrl: {
    issueTokens(value: MsgIssueTokens) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgIssueTokens",
        value
      };
    },

    redeemTokens(value: MsgRedeemTokens) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgRedeemTokens",
        value
      };
    },

    blockAddress(value: MsgBlockAddress) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgBlockAddress",
        value
      };
    },

    unblockAddress(value: MsgUnblockAddress) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgUnblockAddress",
        value
      };
    },

    setPauseStatus(value: MsgSetPauseStatus) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgSetPauseStatus",
        value
      };
    }

  },
  fromPartial: {
    issueTokens(value: MsgIssueTokens) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgIssueTokens",
        value: MsgIssueTokens.fromPartial(value)
      };
    },

    redeemTokens(value: MsgRedeemTokens) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgRedeemTokens",
        value: MsgRedeemTokens.fromPartial(value)
      };
    },

    blockAddress(value: MsgBlockAddress) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgBlockAddress",
        value: MsgBlockAddress.fromPartial(value)
      };
    },

    unblockAddress(value: MsgUnblockAddress) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgUnblockAddress",
        value: MsgUnblockAddress.fromPartial(value)
      };
    },

    setPauseStatus(value: MsgSetPauseStatus) {
      return {
        typeUrl: "/joltify.third_party.issuance.v1beta1.MsgSetPauseStatus",
        value: MsgSetPauseStatus.fromPartial(value)
      };
    }

  }
};