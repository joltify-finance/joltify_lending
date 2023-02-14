import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePool } from "./types/joltify/spv/tx";
import { MsgAddInvestors } from "./types/joltify/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.spv.MsgCreatePool", MsgCreatePool],
    ["/joltify.spv.MsgAddInvestors", MsgAddInvestors],
    
];

export { msgTypes }