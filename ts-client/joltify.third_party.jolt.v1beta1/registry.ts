import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRepay } from "./types/joltify/third_party/jolt/v1beta1/tx";
import { MsgBorrow } from "./types/joltify/third_party/jolt/v1beta1/tx";
import { MsgDeposit } from "./types/joltify/third_party/jolt/v1beta1/tx";
import { MsgWithdraw } from "./types/joltify/third_party/jolt/v1beta1/tx";
import { MsgLiquidate } from "./types/joltify/third_party/jolt/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.third_party.jolt.v1beta1.MsgRepay", MsgRepay],
    ["/joltify.third_party.jolt.v1beta1.MsgBorrow", MsgBorrow],
    ["/joltify.third_party.jolt.v1beta1.MsgDeposit", MsgDeposit],
    ["/joltify.third_party.jolt.v1beta1.MsgWithdraw", MsgWithdraw],
    ["/joltify.third_party.jolt.v1beta1.MsgLiquidate", MsgLiquidate],
    
];

export { msgTypes }