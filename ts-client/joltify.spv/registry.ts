import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgClaimInterest } from "./types/joltify/spv/tx";
import { MsgActivePool } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgDeposit } from "./types/joltify/spv/tx";
import { MsgBorrow } from "./types/joltify/spv/tx";
import { MsgRepayInterest } from "./types/joltify/spv/tx";
import { MsgUpdatePool } from "./types/joltify/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.spv.MsgCreatePool", MsgCreatePool],
    ["/joltify.spv.MsgClaimInterest", MsgClaimInterest],
    ["/joltify.spv.MsgActivePool", MsgActivePool],
    ["/joltify.spv.MsgAddInvestors", MsgAddInvestors],
    ["/joltify.spv.MsgDeposit", MsgDeposit],
    ["/joltify.spv.MsgBorrow", MsgBorrow],
    ["/joltify.spv.MsgRepayInterest", MsgRepayInterest],
    ["/joltify.spv.MsgUpdatePool", MsgUpdatePool],
    
];

export { msgTypes }