import { Rpc } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
import { MsgUploadInvestor, MsgUploadInvestorResponse } from "./tx";
/** Msg defines the Msg service. */

export interface Msg {
  uploadInvestor(request: MsgUploadInvestor): Promise<MsgUploadInvestorResponse>;
}
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.uploadInvestor = this.uploadInvestor.bind(this);
  }

  uploadInvestor(request: MsgUploadInvestor): Promise<MsgUploadInvestorResponse> {
    const data = MsgUploadInvestor.encode(request).finish();
    const promise = this.rpc.request("joltifyfinance.joltify_lending.kyc.Msg", "UploadInvestor", data);
    return promise.then(data => MsgUploadInvestorResponse.decode(new _m0.Reader(data)));
  }

}