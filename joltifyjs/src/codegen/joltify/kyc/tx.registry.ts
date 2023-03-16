import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgUploadInvestor } from "./tx";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/joltifyfinance.joltify_lending.kyc.MsgUploadInvestor", MsgUploadInvestor]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    uploadInvestor(value: MsgUploadInvestor) {
      return {
        typeUrl: "/joltifyfinance.joltify_lending.kyc.MsgUploadInvestor",
        value: MsgUploadInvestor.encode(value).finish()
      };
    }

  },
  withTypeUrl: {
    uploadInvestor(value: MsgUploadInvestor) {
      return {
        typeUrl: "/joltifyfinance.joltify_lending.kyc.MsgUploadInvestor",
        value
      };
    }

  },
  fromPartial: {
    uploadInvestor(value: MsgUploadInvestor) {
      return {
        typeUrl: "/joltifyfinance.joltify_lending.kyc.MsgUploadInvestor",
        value: MsgUploadInvestor.fromPartial(value)
      };
    }

  }
};