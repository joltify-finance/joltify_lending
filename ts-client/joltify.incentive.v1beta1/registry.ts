import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgClaimUSDXMintingReward } from "./types/joltify/incentive/v1beta1/tx";
import { MsgClaimJoltReward } from "./types/joltify/incentive/v1beta1/tx";
import { MsgClaimSavingsReward } from "./types/joltify/incentive/v1beta1/tx";
import { MsgClaimDelegatorReward } from "./types/joltify/incentive/v1beta1/tx";
import { MsgClaimSwapReward } from "./types/joltify/incentive/v1beta1/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.incentive.v1beta1.MsgClaimUSDXMintingReward", MsgClaimUSDXMintingReward],
    ["/joltify.incentive.v1beta1.MsgClaimJoltReward", MsgClaimJoltReward],
    ["/joltify.incentive.v1beta1.MsgClaimSavingsReward", MsgClaimSavingsReward],
    ["/joltify.incentive.v1beta1.MsgClaimDelegatorReward", MsgClaimDelegatorReward],
    ["/joltify.incentive.v1beta1.MsgClaimSwapReward", MsgClaimSwapReward],
    
];

export { msgTypes }