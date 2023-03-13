import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgPayPrincipal } from "./types/joltify/spv/tx";
import { MsgWithdrawPrincipal } from "./types/joltify/spv/tx";
import { MsgDeposit } from "./types/joltify/spv/tx";
import { MsgBorrow } from "./types/joltify/spv/tx";
import { MsgClaimInterest } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgUpdatePool } from "./types/joltify/spv/tx";
import { MsgTransferOwnership } from "./types/joltify/spv/tx";
import { MsgRepayInterest } from "./types/joltify/spv/tx";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgActivePool } from "./types/joltify/spv/tx";
import { MsgSubmitWithdrawProposal } from "./types/joltify/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.spv.MsgPayPrincipal", MsgPayPrincipal],
    ["/joltify.spv.MsgWithdrawPrincipal", MsgWithdrawPrincipal],
    ["/joltify.spv.MsgDeposit", MsgDeposit],
    ["/joltify.spv.MsgBorrow", MsgBorrow],
    ["/joltify.spv.MsgClaimInterest", MsgClaimInterest],
    ["/joltify.spv.MsgAddInvestors", MsgAddInvestors],
    ["/joltify.spv.MsgUpdatePool", MsgUpdatePool],
    ["/joltify.spv.MsgTransferOwnership", MsgTransferOwnership],
    ["/joltify.spv.MsgRepayInterest", MsgRepayInterest],
    ["/joltify.spv.MsgCreatePool", MsgCreatePool],
    ["/joltify.spv.MsgActivePool", MsgActivePool],
    ["/joltify.spv.MsgSubmitWithdrawProposal", MsgSubmitWithdrawProposal],
    
];

export { msgTypes }