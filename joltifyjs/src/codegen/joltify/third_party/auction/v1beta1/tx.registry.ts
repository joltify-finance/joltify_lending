import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgPlaceBid } from "./tx";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/joltify.third_party.auction.v1beta1.MsgPlaceBid", MsgPlaceBid]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    placeBid(value: MsgPlaceBid) {
      return {
        typeUrl: "/joltify.third_party.auction.v1beta1.MsgPlaceBid",
        value: MsgPlaceBid.encode(value).finish()
      };
    }

  },
  withTypeUrl: {
    placeBid(value: MsgPlaceBid) {
      return {
        typeUrl: "/joltify.third_party.auction.v1beta1.MsgPlaceBid",
        value
      };
    }

  },
  fromPartial: {
    placeBid(value: MsgPlaceBid) {
      return {
        typeUrl: "/joltify.third_party.auction.v1beta1.MsgPlaceBid",
        value: MsgPlaceBid.fromPartial(value)
      };
    }

  }
};