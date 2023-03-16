import { AminoMsg } from "@cosmjs/amino";
import { Long } from "../../../../helpers";
import { MsgPlaceBid } from "./tx";
export interface MsgPlaceBidAminoType extends AminoMsg {
  type: "/joltify.third_party.auction.v1beta1.MsgPlaceBid";
  value: {
    auction_id: string;
    bidder: string;
    amount: {
      denom: string;
      amount: string;
    };
  };
}
export const AminoConverter = {
  "/joltify.third_party.auction.v1beta1.MsgPlaceBid": {
    aminoType: "/joltify.third_party.auction.v1beta1.MsgPlaceBid",
    toAmino: ({
      auctionId,
      bidder,
      amount
    }: MsgPlaceBid): MsgPlaceBidAminoType["value"] => {
      return {
        auction_id: auctionId.toString(),
        bidder,
        amount: {
          denom: amount.denom,
          amount: Long.fromValue(amount.amount).toString()
        }
      };
    },
    fromAmino: ({
      auction_id,
      bidder,
      amount
    }: MsgPlaceBidAminoType["value"]): MsgPlaceBid => {
      return {
        auctionId: Long.fromString(auction_id),
        bidder,
        amount: {
          denom: amount.denom,
          amount: amount.amount
        }
      };
    }
  }
};