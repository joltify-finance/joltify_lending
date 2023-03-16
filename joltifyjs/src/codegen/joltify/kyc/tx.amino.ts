import { AminoMsg } from "@cosmjs/amino";
import { MsgUploadInvestor } from "./tx";
export interface MsgUploadInvestorAminoType extends AminoMsg {
  type: "/joltifyfinance.joltify_lending.kyc.MsgUploadInvestor";
  value: {
    creator: string;
    investorId: string;
    walletAddress: string[];
  };
}
export const AminoConverter = {
  "/joltifyfinance.joltify_lending.kyc.MsgUploadInvestor": {
    aminoType: "/joltifyfinance.joltify_lending.kyc.MsgUploadInvestor",
    toAmino: ({
      creator,
      investorId,
      walletAddress
    }: MsgUploadInvestor): MsgUploadInvestorAminoType["value"] => {
      return {
        creator,
        investorId,
        walletAddress
      };
    },
    fromAmino: ({
      creator,
      investorId,
      walletAddress
    }: MsgUploadInvestorAminoType["value"]): MsgUploadInvestor => {
      return {
        creator,
        investorId,
        walletAddress
      };
    }
  }
};