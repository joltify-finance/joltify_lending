import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgCreatePool, MsgAddInvestors, MsgDeposit, MsgBorrow, MsgRepayInterest, MsgClaimInterest, MsgUpdatePool, MsgActivePool, MsgPayPrincipal, MsgWithdrawPrincipal, MsgSubmitWithdrawProposal, MsgTransferOwnership } from "./tx";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/joltify.spv.MsgCreatePool", MsgCreatePool], ["/joltify.spv.MsgAddInvestors", MsgAddInvestors], ["/joltify.spv.MsgDeposit", MsgDeposit], ["/joltify.spv.MsgBorrow", MsgBorrow], ["/joltify.spv.MsgRepayInterest", MsgRepayInterest], ["/joltify.spv.MsgClaimInterest", MsgClaimInterest], ["/joltify.spv.MsgUpdatePool", MsgUpdatePool], ["/joltify.spv.MsgActivePool", MsgActivePool], ["/joltify.spv.MsgPayPrincipal", MsgPayPrincipal], ["/joltify.spv.MsgWithdrawPrincipal", MsgWithdrawPrincipal], ["/joltify.spv.MsgSubmitWithdrawProposal", MsgSubmitWithdrawProposal], ["/joltify.spv.MsgTransferOwnership", MsgTransferOwnership]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    createPool(value: MsgCreatePool) {
      return {
        typeUrl: "/joltify.spv.MsgCreatePool",
        value: MsgCreatePool.encode(value).finish()
      };
    },

    addInvestors(value: MsgAddInvestors) {
      return {
        typeUrl: "/joltify.spv.MsgAddInvestors",
        value: MsgAddInvestors.encode(value).finish()
      };
    },

    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.spv.MsgDeposit",
        value: MsgDeposit.encode(value).finish()
      };
    },

    borrow(value: MsgBorrow) {
      return {
        typeUrl: "/joltify.spv.MsgBorrow",
        value: MsgBorrow.encode(value).finish()
      };
    },

    repayInterest(value: MsgRepayInterest) {
      return {
        typeUrl: "/joltify.spv.MsgRepayInterest",
        value: MsgRepayInterest.encode(value).finish()
      };
    },

    claimInterest(value: MsgClaimInterest) {
      return {
        typeUrl: "/joltify.spv.MsgClaimInterest",
        value: MsgClaimInterest.encode(value).finish()
      };
    },

    updatePool(value: MsgUpdatePool) {
      return {
        typeUrl: "/joltify.spv.MsgUpdatePool",
        value: MsgUpdatePool.encode(value).finish()
      };
    },

    activePool(value: MsgActivePool) {
      return {
        typeUrl: "/joltify.spv.MsgActivePool",
        value: MsgActivePool.encode(value).finish()
      };
    },

    payPrincipal(value: MsgPayPrincipal) {
      return {
        typeUrl: "/joltify.spv.MsgPayPrincipal",
        value: MsgPayPrincipal.encode(value).finish()
      };
    },

    withdrawPrincipal(value: MsgWithdrawPrincipal) {
      return {
        typeUrl: "/joltify.spv.MsgWithdrawPrincipal",
        value: MsgWithdrawPrincipal.encode(value).finish()
      };
    },

    submitWithdrawProposal(value: MsgSubmitWithdrawProposal) {
      return {
        typeUrl: "/joltify.spv.MsgSubmitWithdrawProposal",
        value: MsgSubmitWithdrawProposal.encode(value).finish()
      };
    },

    transferOwnership(value: MsgTransferOwnership) {
      return {
        typeUrl: "/joltify.spv.MsgTransferOwnership",
        value: MsgTransferOwnership.encode(value).finish()
      };
    }

  },
  withTypeUrl: {
    createPool(value: MsgCreatePool) {
      return {
        typeUrl: "/joltify.spv.MsgCreatePool",
        value
      };
    },

    addInvestors(value: MsgAddInvestors) {
      return {
        typeUrl: "/joltify.spv.MsgAddInvestors",
        value
      };
    },

    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.spv.MsgDeposit",
        value
      };
    },

    borrow(value: MsgBorrow) {
      return {
        typeUrl: "/joltify.spv.MsgBorrow",
        value
      };
    },

    repayInterest(value: MsgRepayInterest) {
      return {
        typeUrl: "/joltify.spv.MsgRepayInterest",
        value
      };
    },

    claimInterest(value: MsgClaimInterest) {
      return {
        typeUrl: "/joltify.spv.MsgClaimInterest",
        value
      };
    },

    updatePool(value: MsgUpdatePool) {
      return {
        typeUrl: "/joltify.spv.MsgUpdatePool",
        value
      };
    },

    activePool(value: MsgActivePool) {
      return {
        typeUrl: "/joltify.spv.MsgActivePool",
        value
      };
    },

    payPrincipal(value: MsgPayPrincipal) {
      return {
        typeUrl: "/joltify.spv.MsgPayPrincipal",
        value
      };
    },

    withdrawPrincipal(value: MsgWithdrawPrincipal) {
      return {
        typeUrl: "/joltify.spv.MsgWithdrawPrincipal",
        value
      };
    },

    submitWithdrawProposal(value: MsgSubmitWithdrawProposal) {
      return {
        typeUrl: "/joltify.spv.MsgSubmitWithdrawProposal",
        value
      };
    },

    transferOwnership(value: MsgTransferOwnership) {
      return {
        typeUrl: "/joltify.spv.MsgTransferOwnership",
        value
      };
    }

  },
  fromPartial: {
    createPool(value: MsgCreatePool) {
      return {
        typeUrl: "/joltify.spv.MsgCreatePool",
        value: MsgCreatePool.fromPartial(value)
      };
    },

    addInvestors(value: MsgAddInvestors) {
      return {
        typeUrl: "/joltify.spv.MsgAddInvestors",
        value: MsgAddInvestors.fromPartial(value)
      };
    },

    deposit(value: MsgDeposit) {
      return {
        typeUrl: "/joltify.spv.MsgDeposit",
        value: MsgDeposit.fromPartial(value)
      };
    },

    borrow(value: MsgBorrow) {
      return {
        typeUrl: "/joltify.spv.MsgBorrow",
        value: MsgBorrow.fromPartial(value)
      };
    },

    repayInterest(value: MsgRepayInterest) {
      return {
        typeUrl: "/joltify.spv.MsgRepayInterest",
        value: MsgRepayInterest.fromPartial(value)
      };
    },

    claimInterest(value: MsgClaimInterest) {
      return {
        typeUrl: "/joltify.spv.MsgClaimInterest",
        value: MsgClaimInterest.fromPartial(value)
      };
    },

    updatePool(value: MsgUpdatePool) {
      return {
        typeUrl: "/joltify.spv.MsgUpdatePool",
        value: MsgUpdatePool.fromPartial(value)
      };
    },

    activePool(value: MsgActivePool) {
      return {
        typeUrl: "/joltify.spv.MsgActivePool",
        value: MsgActivePool.fromPartial(value)
      };
    },

    payPrincipal(value: MsgPayPrincipal) {
      return {
        typeUrl: "/joltify.spv.MsgPayPrincipal",
        value: MsgPayPrincipal.fromPartial(value)
      };
    },

    withdrawPrincipal(value: MsgWithdrawPrincipal) {
      return {
        typeUrl: "/joltify.spv.MsgWithdrawPrincipal",
        value: MsgWithdrawPrincipal.fromPartial(value)
      };
    },

    submitWithdrawProposal(value: MsgSubmitWithdrawProposal) {
      return {
        typeUrl: "/joltify.spv.MsgSubmitWithdrawProposal",
        value: MsgSubmitWithdrawProposal.fromPartial(value)
      };
    },

    transferOwnership(value: MsgTransferOwnership) {
      return {
        typeUrl: "/joltify.spv.MsgTransferOwnership",
        value: MsgTransferOwnership.fromPartial(value)
      };
    }

  }
};