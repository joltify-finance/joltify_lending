import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgCreateCDP, MsgDeposit, MsgWithdraw, MsgDrawDebt, MsgRepayDebt, MsgLiquidate } from "./tx";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/joltify.third_party.cdp.v1beta1.MsgCreateCDP", MsgCreateCDP], ["/joltify.third_party.cdp.v1beta1.MsgDeposit", MsgDeposit], ["/joltify.third_party.cdp.v1beta1.MsgWithdraw", MsgWithdraw], ["/joltify.third_party.cdp.v1beta1.MsgDrawDebt", MsgDrawDebt], ["/joltify.third_party.cdp.v1beta1.MsgRepayDebt", MsgRepayDebt], ["/joltify.third_party.cdp.v1beta1.MsgLiquidate", MsgLiquidate]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    createCDP(value: MsgCreateCDP) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgCreateCDP",
        value: MsgCreateCDP.encode(value).finish()
      };
    },

    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgDeposit",
        value: MsgDeposit.encode(value).finish()
      };
    },

    withdraw(value: MsgWithdraw) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgWithdraw",
        value: MsgWithdraw.encode(value).finish()
      };
    },

    drawDebt(value: MsgDrawDebt) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgDrawDebt",
        value: MsgDrawDebt.encode(value).finish()
      };
    },

    repayDebt(value: MsgRepayDebt) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgRepayDebt",
        value: MsgRepayDebt.encode(value).finish()
      };
    },

    liquidate(value: MsgLiquidate) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgLiquidate",
        value: MsgLiquidate.encode(value).finish()
      };
    }

  },
  withTypeUrl: {
    createCDP(value: MsgCreateCDP) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgCreateCDP",
        value
      };
    },

    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgDeposit",
        value
      };
    },

    withdraw(value: MsgWithdraw) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgWithdraw",
        value
      };
    },

    drawDebt(value: MsgDrawDebt) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgDrawDebt",
        value
      };
    },

    repayDebt(value: MsgRepayDebt) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgRepayDebt",
        value
      };
    },

    liquidate(value: MsgLiquidate) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgLiquidate",
        value
      };
    }

  },
  fromPartial: {
    createCDP(value: MsgCreateCDP) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgCreateCDP",
        value: MsgCreateCDP.fromPartial(value)
      };
    },

    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgDeposit",
        value: MsgDeposit.fromPartial(value)
      };
    },

    withdraw(value: MsgWithdraw) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgWithdraw",
        value: MsgWithdraw.fromPartial(value)
      };
    },

    drawDebt(value: MsgDrawDebt) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgDrawDebt",
        value: MsgDrawDebt.fromPartial(value)
      };
    },

    repayDebt(value: MsgRepayDebt) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgRepayDebt",
        value: MsgRepayDebt.fromPartial(value)
      };
    },

    liquidate(value: MsgLiquidate) {
      return {
        typeUrl: "/joltify.third_party.cdp.v1beta1.MsgLiquidate",
        value: MsgLiquidate.fromPartial(value)
      };
    }

  }
};