import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgClaimSwapReward } from "./types/joltify/third_party/incentive/v1beta1/tx";
import { MsgClaimUSDXMintingReward } from "./types/joltify/third_party/incentive/v1beta1/tx";
import { MsgClaimJoltReward } from "./types/joltify/third_party/incentive/v1beta1/tx";
import { MsgClaimSavingsReward } from "./types/joltify/third_party/incentive/v1beta1/tx";
import { MsgClaimDelegatorReward } from "./types/joltify/third_party/incentive/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.third_party.incentive.v1beta1.MsgClaimSwapReward", MsgClaimSwapReward],
    ["/joltify.third_party.incentive.v1beta1.MsgClaimUSDXMintingReward", MsgClaimUSDXMintingReward],
    ["/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward", MsgClaimJoltReward],
    ["/joltify.third_party.incentive.v1beta1.MsgClaimSavingsReward", MsgClaimSavingsReward],
    ["/joltify.third_party.incentive.v1beta1.MsgClaimDelegatorReward", MsgClaimDelegatorReward],
    
];

export { msgTypes }