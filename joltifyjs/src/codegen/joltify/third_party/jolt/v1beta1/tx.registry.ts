import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgDeposit, MsgWithdraw, MsgBorrow, MsgRepay, MsgLiquidate } from "./tx";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/joltify.third_party.jolt.v1beta1.MsgDeposit", MsgDeposit], ["/joltify.third_party.jolt.v1beta1.MsgWithdraw", MsgWithdraw], ["/joltify.third_party.jolt.v1beta1.MsgBorrow", MsgBorrow], ["/joltify.third_party.jolt.v1beta1.MsgRepay", MsgRepay], ["/joltify.third_party.jolt.v1beta1.MsgLiquidate", MsgLiquidate]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgDeposit",
        value: MsgDeposit.encode(value).finish()
      };
    },

    withdraw(value: MsgWithdraw) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgWithdraw",
        value: MsgWithdraw.encode(value).finish()
      };
    },

    borrow(value: MsgBorrow) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgBorrow",
        value: MsgBorrow.encode(value).finish()
      };
    },

    repay(value: MsgRepay) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgRepay",
        value: MsgRepay.encode(value).finish()
      };
    },

    liquidate(value: MsgLiquidate) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgLiquidate",
        value: MsgLiquidate.encode(value).finish()
      };
    }

  },
  withTypeUrl: {
    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgDeposit",
        value
      };
    },

    withdraw(value: MsgWithdraw) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgWithdraw",
        value
      };
    },

    borrow(value: MsgBorrow) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgBorrow",
        value
      };
    },

    repay(value: MsgRepay) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgRepay",
        value
      };
    },

    liquidate(value: MsgLiquidate) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgLiquidate",
        value
      };
    }

  },
  fromPartial: {
    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgDeposit",
        value: MsgDeposit.fromPartial(value)
      };
    },

    withdraw(value: MsgWithdraw) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgWithdraw",
        value: MsgWithdraw.fromPartial(value)
      };
    },

    borrow(value: MsgBorrow) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgBorrow",
        value: MsgBorrow.fromPartial(value)
      };
    },

    repay(value: MsgRepay) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgRepay",
        value: MsgRepay.fromPartial(value)
      };
    },

    liquidate(value: MsgLiquidate) {
      return {
        typeUrl: "/joltify.third_party.jolt.v1beta1.MsgLiquidate",
        value: MsgLiquidate.fromPartial(value)
      };
    }

  }
};