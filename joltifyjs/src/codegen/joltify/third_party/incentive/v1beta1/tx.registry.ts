import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgClaimJoltReward } from "./tx";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward", MsgClaimJoltReward]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    claimJoltReward(value: MsgClaimJoltReward) {
      return {
        typeUrl: "/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward",
        value: MsgClaimJoltReward.encode(value).finish()
      };
    }

  },
  withTypeUrl: {
    claimJoltReward(value: MsgClaimJoltReward) {
      return {
        typeUrl: "/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward",
        value
      };
    }

  },
  fromPartial: {
    claimJoltReward(value: MsgClaimJoltReward) {
      return {
        typeUrl: "/joltify.third_party.incentive.v1beta1.MsgClaimJoltReward",
        value: MsgClaimJoltReward.fromPartial(value)
      };
    }

  }
};