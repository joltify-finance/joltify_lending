import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRepayDebt } from "./types/joltify/cdp/v1beta1/tx";
import { MsgCreateCDP } from "./types/joltify/cdp/v1beta1/tx";
import { MsgDeposit } from "./types/joltify/cdp/v1beta1/tx";
import { MsgLiquidate } from "./types/joltify/cdp/v1beta1/tx";
import { MsgWithdraw } from "./types/joltify/cdp/v1beta1/tx";
import { MsgDrawDebt } from "./types/joltify/cdp/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.cdp.v1beta1.MsgRepayDebt", MsgRepayDebt],
    ["/joltify.cdp.v1beta1.MsgCreateCDP", MsgCreateCDP],
    ["/joltify.cdp.v1beta1.MsgDeposit", MsgDeposit],
    ["/joltify.cdp.v1beta1.MsgLiquidate", MsgLiquidate],
    ["/joltify.cdp.v1beta1.MsgWithdraw", MsgWithdraw],
    ["/joltify.cdp.v1beta1.MsgDrawDebt", MsgDrawDebt],
    
];

export { msgTypes }