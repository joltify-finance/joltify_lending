import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgDeposit } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgLiquidate } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgCreateCDP } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgWithdraw } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgDrawDebt } from "./types/joltify/third_party/cdp/v1beta1/tx";
import { MsgRepayDebt } from "./types/joltify/third_party/cdp/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.third_party.cdp.v1beta1.MsgDeposit", MsgDeposit],
    ["/joltify.third_party.cdp.v1beta1.MsgLiquidate", MsgLiquidate],
    ["/joltify.third_party.cdp.v1beta1.MsgCreateCDP", MsgCreateCDP],
    ["/joltify.third_party.cdp.v1beta1.MsgWithdraw", MsgWithdraw],
    ["/joltify.third_party.cdp.v1beta1.MsgDrawDebt", MsgDrawDebt],
    ["/joltify.third_party.cdp.v1beta1.MsgRepayDebt", MsgRepayDebt],
    
];

export { msgTypes }