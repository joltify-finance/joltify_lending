import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeposit } from "./types/joltify/spv/tx";
import { MsgClaimInterest } from "./types/joltify/spv/tx";
import { MsgActivePool } from "./types/joltify/spv/tx";
import { MsgPayPrincipal } from "./types/joltify/spv/tx";
import { MsgWithdrawPrincipal } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgBorrow } from "./types/joltify/spv/tx";
import { MsgUpdatePool } from "./types/joltify/spv/tx";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgRepayInterest } from "./types/joltify/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.spv.MsgDeposit", MsgDeposit],
    ["/joltify.spv.MsgClaimInterest", MsgClaimInterest],
    ["/joltify.spv.MsgActivePool", MsgActivePool],
    ["/joltify.spv.MsgPayPrincipal", MsgPayPrincipal],
    ["/joltify.spv.MsgWithdrawPrincipal", MsgWithdrawPrincipal],
    ["/joltify.spv.MsgAddInvestors", MsgAddInvestors],
    ["/joltify.spv.MsgBorrow", MsgBorrow],
    ["/joltify.spv.MsgUpdatePool", MsgUpdatePool],
    ["/joltify.spv.MsgCreatePool", MsgCreatePool],
    ["/joltify.spv.MsgRepayInterest", MsgRepayInterest],
    
];

export { msgTypes }