import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeposit } from "./types/joltify/spv/tx";
import { MsgPayPrincipal } from "./types/joltify/spv/tx";
import { MsgRepayInterest } from "./types/joltify/spv/tx";
import { MsgUpdatePool } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgWithdrawPrincipal } from "./types/joltify/spv/tx";
import { MsgClaimInterest } from "./types/joltify/spv/tx";
import { MsgBorrow } from "./types/joltify/spv/tx";
import { MsgActivePool } from "./types/joltify/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.spv.MsgDeposit", MsgDeposit],
    ["/joltify.spv.MsgPayPrincipal", MsgPayPrincipal],
    ["/joltify.spv.MsgRepayInterest", MsgRepayInterest],
    ["/joltify.spv.MsgUpdatePool", MsgUpdatePool],
    ["/joltify.spv.MsgAddInvestors", MsgAddInvestors],
    ["/joltify.spv.MsgCreatePool", MsgCreatePool],
    ["/joltify.spv.MsgWithdrawPrincipal", MsgWithdrawPrincipal],
    ["/joltify.spv.MsgClaimInterest", MsgClaimInterest],
    ["/joltify.spv.MsgBorrow", MsgBorrow],
    ["/joltify.spv.MsgActivePool", MsgActivePool],
    
];

export { msgTypes }