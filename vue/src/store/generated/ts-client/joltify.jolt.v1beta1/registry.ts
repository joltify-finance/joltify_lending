import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgBorrow } from "./types/joltify/jolt/v1beta1/tx";
import { MsgWithdraw } from "./types/joltify/jolt/v1beta1/tx";
import { MsgDeposit } from "./types/joltify/jolt/v1beta1/tx";
import { MsgRepay } from "./types/joltify/jolt/v1beta1/tx";
import { MsgLiquidate } from "./types/joltify/jolt/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.jolt.v1beta1.MsgBorrow", MsgBorrow],
    ["/joltify.jolt.v1beta1.MsgWithdraw", MsgWithdraw],
    ["/joltify.jolt.v1beta1.MsgDeposit", MsgDeposit],
    ["/joltify.jolt.v1beta1.MsgRepay", MsgRepay],
    ["/joltify.jolt.v1beta1.MsgLiquidate", MsgLiquidate],
    
];

export { msgTypes }