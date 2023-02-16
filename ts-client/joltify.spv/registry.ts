import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddInvestors } from "./types/joltify/spv/tx";
import { MsgDeposit } from "./types/joltify/spv/tx";
import { MsgCreatePool } from "./types/joltify/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.spv.MsgAddInvestors", MsgAddInvestors],
    ["/joltify.spv.MsgDeposit", MsgDeposit],
    ["/joltify.spv.MsgCreatePool", MsgCreatePool],
    
];

export { msgTypes }