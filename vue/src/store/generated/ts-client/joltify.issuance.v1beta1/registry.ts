import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgIssueTokens } from "./types/joltify/issuance/v1beta1/tx";
import { MsgBlockAddress } from "./types/joltify/issuance/v1beta1/tx";
import { MsgSetPauseStatus } from "./types/joltify/issuance/v1beta1/tx";
import { MsgRedeemTokens } from "./types/joltify/issuance/v1beta1/tx";
import { MsgUnblockAddress } from "./types/joltify/issuance/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.issuance.v1beta1.MsgIssueTokens", MsgIssueTokens],
    ["/joltify.issuance.v1beta1.MsgBlockAddress", MsgBlockAddress],
    ["/joltify.issuance.v1beta1.MsgSetPauseStatus", MsgSetPauseStatus],
    ["/joltify.issuance.v1beta1.MsgRedeemTokens", MsgRedeemTokens],
    ["/joltify.issuance.v1beta1.MsgUnblockAddress", MsgUnblockAddress],
    
];

export { msgTypes }