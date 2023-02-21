import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgClaimInterest } from "./types/joltifylending/spv/tx";
import { MsgRepayInterest } from "./types/joltifylending/spv/tx";
import { MsgUpdatePool } from "./types/spv/tx";
import { MsgActivePool } from "./types/spv/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltifyfinance.joltify_lending.spv.MsgClaimInterest", MsgClaimInterest],
    ["/joltifyfinance.joltify_lending.spv.MsgRepayInterest", MsgRepayInterest],
    ["/joltifyfinance.joltify_lending.spv.MsgUpdatePool", MsgUpdatePool],
    ["/joltifyfinance.joltify_lending.spv.MsgActivePool", MsgActivePool],
    
];

export { msgTypes }