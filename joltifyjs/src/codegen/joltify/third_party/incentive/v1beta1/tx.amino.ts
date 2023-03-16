import { AminoMsg } from "@cosmjs/amino";
import { MsgClaimJoltReward } from "./tx";
export interface MsgClaimJoltRewardAminoType extends AminoMsg {
  type: "/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward";
  value: {
    sender: string;
    denoms_to_claim: {
      denom: string;
      multiplier_name: string;
    }[];
  };
}
export const AminoConverter = {
  "/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward": {
    aminoType: "/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward",
    toAmino: ({
      sender,
      denomsToClaim
    }: MsgClaimJoltReward): MsgClaimJoltRewardAminoType["value"] => {
      return {
        sender,
        denoms_to_claim: denomsToClaim.map(el0 => ({
          denom: el0.denom,
          multiplier_name: el0.multiplierName
        }))
      };
    },
    fromAmino: ({
      sender,
      denoms_to_claim
    }: MsgClaimJoltRewardAminoType["value"]): MsgClaimJoltReward => {
      return {
        sender,
        denomsToClaim: denoms_to_claim.map(el0 => ({
          denom: el0.denom,
          multiplierName: el0.multiplier_name
        }))
      };
    }
  }
};