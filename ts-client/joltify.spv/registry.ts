import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRepayInterest } from "./types/joltify/spv/tx";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgBorrow } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgDeposit } from "./types/joltify/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.spv.MsgRepayInterest", MsgRepayInterest],
    ["/joltify.spv.MsgCreatePool", MsgCreatePool],
    ["/joltify.spv.MsgBorrow", MsgBorrow],
    ["/joltify.spv.MsgAddInvestors", MsgAddInvestors],
    ["/joltify.spv.MsgDeposit", MsgDeposit],
    
];

export { msgTypes }