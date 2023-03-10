import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgWithdrawPrincipal } from "./types/joltify/spv/tx";
import { MsgRepayInterest } from "./types/joltify/spv/tx";
import { MsgClaimInterest } from "./types/joltify/spv/tx";
import { MsgSubmitWithdrawProposal } from "./types/joltify/spv/tx";
import { MsgUpdatePool } from "./types/joltify/spv/tx";
import { MsgTransferOwnership } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgDeposit } from "./types/joltify/spv/tx";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgActivePool } from "./types/joltify/spv/tx";
import { MsgPayPrincipal } from "./types/joltify/spv/tx";
import { MsgBorrow } from "./types/joltify/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.spv.MsgWithdrawPrincipal", MsgWithdrawPrincipal],
    ["/joltify.spv.MsgRepayInterest", MsgRepayInterest],
    ["/joltify.spv.MsgClaimInterest", MsgClaimInterest],
    ["/joltify.spv.MsgSubmitWithdrawProposal", MsgSubmitWithdrawProposal],
    ["/joltify.spv.MsgUpdatePool", MsgUpdatePool],
    ["/joltify.spv.MsgTransferOwnership", MsgTransferOwnership],
    ["/joltify.spv.MsgAddInvestors", MsgAddInvestors],
    ["/joltify.spv.MsgDeposit", MsgDeposit],
    ["/joltify.spv.MsgCreatePool", MsgCreatePool],
    ["/joltify.spv.MsgActivePool", MsgActivePool],
    ["/joltify.spv.MsgPayPrincipal", MsgPayPrincipal],
    ["/joltify.spv.MsgBorrow", MsgBorrow],
    
];

export { msgTypes }