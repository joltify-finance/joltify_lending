import { AminoMsg } from "@cosmjs/amino";
import { Long } from "../../helpers";
import { MsgCreatePool, MsgAddInvestors, MsgDeposit, MsgBorrow, MsgRepayInterest, MsgClaimInterest, MsgUpdatePool, MsgActivePool, MsgPayPrincipal, MsgWithdrawPrincipal, MsgSubmitWithdrawProposal, MsgTransferOwnership } from "./tx";
export interface MsgCreatePoolAminoType extends AminoMsg {
  type: "/joltify.spv.MsgCreatePool";
  value: {
    creator: string;
    project_index: number;
    pool_name: string;
    apy: string;
    target_token_amount: {
      denom: string;
      amount: string;
    };
  };
}
export interface MsgAddInvestorsAminoType extends AminoMsg {
  type: "/joltify.spv.MsgAddInvestors";
  value: {
    creator: string;
    pool_index: string;
    investor_iD: string[];
  };
}
export interface MsgDepositAminoType extends AminoMsg {
  type: "/joltify.spv.MsgDeposit";
  value: {
    creator: string;
    pool_index: string;
    token: {
      denom: string;
      amount: string;
    };
  };
}
export interface MsgBorrowAminoType extends AminoMsg {
  type: "/joltify.spv.MsgBorrow";
  value: {
    creator: string;
    pool_index: string;
    borrow_amount: {
      denom: string;
      amount: string;
    };
  };
}
export interface MsgRepayInterestAminoType extends AminoMsg {
  type: "/joltify.spv.MsgRepayInterest";
  value: {
    creator: string;
    pool_index: string;
    token: {
      denom: string;
      amount: string;
    };
  };
}
export interface MsgClaimInterestAminoType extends AminoMsg {
  type: "/joltify.spv.MsgClaimInterest";
  value: {
    creator: string;
    pool_index: string;
  };
}
export interface MsgUpdatePoolAminoType extends AminoMsg {
  type: "/joltify.spv.MsgUpdatePool";
  value: {
    creator: string;
    pool_index: string;
    pool_name: string;
    pool_apy: string;
    target_token_amount: {
      denom: string;
      amount: string;
    };
  };
}
export interface MsgActivePoolAminoType extends AminoMsg {
  type: "/joltify.spv.MsgActivePool";
  value: {
    creator: string;
    pool_index: string;
  };
}
export interface MsgPayPrincipalAminoType extends AminoMsg {
  type: "/joltify.spv.MsgPayPrincipal";
  value: {
    creator: string;
    pool_index: string;
    token: {
      denom: string;
      amount: string;
    };
  };
}
export interface MsgWithdrawPrincipalAminoType extends AminoMsg {
  type: "/joltify.spv.MsgWithdrawPrincipal";
  value: {
    creator: string;
    pool_index: string;
    token: {
      denom: string;
      amount: string;
    };
  };
}
export interface MsgSubmitWithdrawProposalAminoType extends AminoMsg {
  type: "/joltify.spv.MsgSubmitWithdrawProposal";
  value: {
    creator: string;
    pool_index: string;
  };
}
export interface MsgTransferOwnershipAminoType extends AminoMsg {
  type: "/joltify.spv.MsgTransferOwnership";
  value: {
    creator: string;
    pool_index: string;
  };
}
export const AminoConverter = {
  "/joltify.spv.MsgCreatePool": {
    aminoType: "/joltify.spv.MsgCreatePool",
    toAmino: ({
      creator,
      projectIndex,
      poolName,
      apy,
      targetTokenAmount
    }: MsgCreatePool): MsgCreatePoolAminoType["value"] => {
      return {
        creator,
        project_index: projectIndex,
        pool_name: poolName,
        apy,
        target_token_amount: {
          denom: targetTokenAmount.denom,
          amount: Long.fromValue(targetTokenAmount.amount).toString()
        }
      };
    },
    fromAmino: ({
      creator,
      project_index,
      pool_name,
      apy,
      target_token_amount
    }: MsgCreatePoolAminoType["value"]): MsgCreatePool => {
      return {
        creator,
        projectIndex: project_index,
        poolName: pool_name,
        apy,
        targetTokenAmount: {
          denom: target_token_amount.denom,
          amount: target_token_amount.amount
        }
      };
    }
  },
  "/joltify.spv.MsgAddInvestors": {
    aminoType: "/joltify.spv.MsgAddInvestors",
    toAmino: ({
      creator,
      poolIndex,
      investorID
    }: MsgAddInvestors): MsgAddInvestorsAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex,
        investor_iD: investorID
      };
    },
    fromAmino: ({
      creator,
      pool_index,
      investor_iD
    }: MsgAddInvestorsAminoType["value"]): MsgAddInvestors => {
      return {
        creator,
        poolIndex: pool_index,
        investorID: investor_iD
      };
    }
  },
  "/joltify.spv.MsgDeposit": {
    aminoType: "/joltify.spv.MsgDeposit",
    toAmino: ({
      creator,
      poolIndex,
      token
    }: MsgDeposit): MsgDepositAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex,
        token: {
          denom: token.denom,
          amount: Long.fromValue(token.amount).toString()
        }
      };
    },
    fromAmino: ({
      creator,
      pool_index,
      token
    }: MsgDepositAminoType["value"]): MsgDeposit => {
      return {
        creator,
        poolIndex: pool_index,
        token: {
          denom: token.denom,
          amount: token.amount
        }
      };
    }
  },
  "/joltify.spv.MsgBorrow": {
    aminoType: "/joltify.spv.MsgBorrow",
    toAmino: ({
      creator,
      poolIndex,
      borrowAmount
    }: MsgBorrow): MsgBorrowAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex,
        borrow_amount: {
          denom: borrowAmount.denom,
          amount: Long.fromValue(borrowAmount.amount).toString()
        }
      };
    },
    fromAmino: ({
      creator,
      pool_index,
      borrow_amount
    }: MsgBorrowAminoType["value"]): MsgBorrow => {
      return {
        creator,
        poolIndex: pool_index,
        borrowAmount: {
          denom: borrow_amount.denom,
          amount: borrow_amount.amount
        }
      };
    }
  },
  "/joltify.spv.MsgRepayInterest": {
    aminoType: "/joltify.spv.MsgRepayInterest",
    toAmino: ({
      creator,
      poolIndex,
      token
    }: MsgRepayInterest): MsgRepayInterestAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex,
        token: {
          denom: token.denom,
          amount: Long.fromValue(token.amount).toString()
        }
      };
    },
    fromAmino: ({
      creator,
      pool_index,
      token
    }: MsgRepayInterestAminoType["value"]): MsgRepayInterest => {
      return {
        creator,
        poolIndex: pool_index,
        token: {
          denom: token.denom,
          amount: token.amount
        }
      };
    }
  },
  "/joltify.spv.MsgClaimInterest": {
    aminoType: "/joltify.spv.MsgClaimInterest",
    toAmino: ({
      creator,
      poolIndex
    }: MsgClaimInterest): MsgClaimInterestAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex
      };
    },
    fromAmino: ({
      creator,
      pool_index
    }: MsgClaimInterestAminoType["value"]): MsgClaimInterest => {
      return {
        creator,
        poolIndex: pool_index
      };
    }
  },
  "/joltify.spv.MsgUpdatePool": {
    aminoType: "/joltify.spv.MsgUpdatePool",
    toAmino: ({
      creator,
      poolIndex,
      poolName,
      poolApy,
      targetTokenAmount
    }: MsgUpdatePool): MsgUpdatePoolAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex,
        pool_name: poolName,
        pool_apy: poolApy,
        target_token_amount: {
          denom: targetTokenAmount.denom,
          amount: Long.fromValue(targetTokenAmount.amount).toString()
        }
      };
    },
    fromAmino: ({
      creator,
      pool_index,
      pool_name,
      pool_apy,
      target_token_amount
    }: MsgUpdatePoolAminoType["value"]): MsgUpdatePool => {
      return {
        creator,
        poolIndex: pool_index,
        poolName: pool_name,
        poolApy: pool_apy,
        targetTokenAmount: {
          denom: target_token_amount.denom,
          amount: target_token_amount.amount
        }
      };
    }
  },
  "/joltify.spv.MsgActivePool": {
    aminoType: "/joltify.spv.MsgActivePool",
    toAmino: ({
      creator,
      poolIndex
    }: MsgActivePool): MsgActivePoolAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex
      };
    },
    fromAmino: ({
      creator,
      pool_index
    }: MsgActivePoolAminoType["value"]): MsgActivePool => {
      return {
        creator,
        poolIndex: pool_index
      };
    }
  },
  "/joltify.spv.MsgPayPrincipal": {
    aminoType: "/joltify.spv.MsgPayPrincipal",
    toAmino: ({
      creator,
      poolIndex,
      token
    }: MsgPayPrincipal): MsgPayPrincipalAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex,
        token: {
          denom: token.denom,
          amount: Long.fromValue(token.amount).toString()
        }
      };
    },
    fromAmino: ({
      creator,
      pool_index,
      token
    }: MsgPayPrincipalAminoType["value"]): MsgPayPrincipal => {
      return {
        creator,
        poolIndex: pool_index,
        token: {
          denom: token.denom,
          amount: token.amount
        }
      };
    }
  },
  "/joltify.spv.MsgWithdrawPrincipal": {
    aminoType: "/joltify.spv.MsgWithdrawPrincipal",
    toAmino: ({
      creator,
      poolIndex,
      token
    }: MsgWithdrawPrincipal): MsgWithdrawPrincipalAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex,
        token: {
          denom: token.denom,
          amount: Long.fromValue(token.amount).toString()
        }
      };
    },
    fromAmino: ({
      creator,
      pool_index,
      token
    }: MsgWithdrawPrincipalAminoType["value"]): MsgWithdrawPrincipal => {
      return {
        creator,
        poolIndex: pool_index,
        token: {
          denom: token.denom,
          amount: token.amount
        }
      };
    }
  },
  "/joltify.spv.MsgSubmitWithdrawProposal": {
    aminoType: "/joltify.spv.MsgSubmitWithdrawProposal",
    toAmino: ({
      creator,
      poolIndex
    }: MsgSubmitWithdrawProposal): MsgSubmitWithdrawProposalAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex
      };
    },
    fromAmino: ({
      creator,
      pool_index
    }: MsgSubmitWithdrawProposalAminoType["value"]): MsgSubmitWithdrawProposal => {
      return {
        creator,
        poolIndex: pool_index
      };
    }
  },
  "/joltify.spv.MsgTransferOwnership": {
    aminoType: "/joltify.spv.MsgTransferOwnership",
    toAmino: ({
      creator,
      poolIndex
    }: MsgTransferOwnership): MsgTransferOwnershipAminoType["value"] => {
      return {
        creator,
        pool_index: poolIndex
      };
    },
    fromAmino: ({
      creator,
      pool_index
    }: MsgTransferOwnershipAminoType["value"]): MsgTransferOwnership => {
      return {
        creator,
        poolIndex: pool_index
      };
    }
  }
};