import { Rpc } from "../../../../helpers";
import * as _m0 from "protobufjs/minimal";
import { MsgClaimJoltReward, MsgClaimJoltRewardResponse } from "./tx";
/** Msg defines the incentive Msg service. */

export interface Msg {
  /** ClaimJoltReward is a message type used to claim Hard liquidity provider rewards */
  claimJoltReward(request: MsgClaimJoltReward): Promise<MsgClaimJoltRewardResponse>;
}
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.claimJoltReward = this.claimJoltReward.bind(this);
  }

  claimJoltReward(request: MsgClaimJoltReward): Promise<MsgClaimJoltRewardResponse> {
    const data = MsgClaimJoltReward.encode(request).finish();
    const promise = this.rpc.request("joltify.third_party.incentive.v1beta1.Msg", "ClaimJoltReward", data);
    return promise.then(data => MsgClaimJoltRewardResponse.decode(new _m0.Reader(data)));
  }

}