import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgWithdraw } from "./types/joltify/savings/v1beta1/tx";
import { MsgDeposit } from "./types/joltify/savings/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.savings.v1beta1.MsgWithdraw", MsgWithdraw],
    ["/joltify.savings.v1beta1.MsgDeposit", MsgDeposit],
    
];

export { msgTypes }