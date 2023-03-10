import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRedeemTokens } from "./types/joltify/third_party/issuance/v1beta1/tx";
import { MsgBlockAddress } from "./types/joltify/third_party/issuance/v1beta1/tx";
import { MsgUnblockAddress } from "./types/joltify/third_party/issuance/v1beta1/tx";
import { MsgIssueTokens } from "./types/joltify/third_party/issuance/v1beta1/tx";
import { MsgSetPauseStatus } from "./types/joltify/third_party/issuance/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.third_party.issuance.v1beta1.MsgRedeemTokens", MsgRedeemTokens],
    ["/joltify.third_party.issuance.v1beta1.MsgBlockAddress", MsgBlockAddress],
    ["/joltify.third_party.issuance.v1beta1.MsgUnblockAddress", MsgUnblockAddress],
    ["/joltify.third_party.issuance.v1beta1.MsgIssueTokens", MsgIssueTokens],
    ["/joltify.third_party.issuance.v1beta1.MsgSetPauseStatus", MsgSetPauseStatus],
    
];

export { msgTypes }