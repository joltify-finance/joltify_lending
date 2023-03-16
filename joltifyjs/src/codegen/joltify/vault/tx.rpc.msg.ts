import { Rpc } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
import { MsgCreateOutboundTx, MsgCreateOutboundTxResponse, MsgCreateIssueToken, MsgCreateIssueTokenResponse, MsgCreateCreatePool, MsgCreateCreatePoolResponse } from "./tx";
/** Msg defines the Msg service. */

export interface Msg {
  createOutboundTx(request: MsgCreateOutboundTx): Promise<MsgCreateOutboundTxResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */

  createIssueToken(request: MsgCreateIssueToken): Promise<MsgCreateIssueTokenResponse>;
  createCreatePool(request: MsgCreateCreatePool): Promise<MsgCreateCreatePoolResponse>;
}
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.createOutboundTx = this.createOutboundTx.bind(this);
    this.createIssueToken = this.createIssueToken.bind(this);
    this.createCreatePool = this.createCreatePool.bind(this);
  }

  createOutboundTx(request: MsgCreateOutboundTx): Promise<MsgCreateOutboundTxResponse> {
    const data = MsgCreateOutboundTx.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Msg", "CreateOutboundTx", data);
    return promise.then(data => MsgCreateOutboundTxResponse.decode(new _m0.Reader(data)));
  }

  createIssueToken(request: MsgCreateIssueToken): Promise<MsgCreateIssueTokenResponse> {
    const data = MsgCreateIssueToken.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Msg", "CreateIssueToken", data);
    return promise.then(data => MsgCreateIssueTokenResponse.decode(new _m0.Reader(data)));
  }

  createCreatePool(request: MsgCreateCreatePool): Promise<MsgCreateCreatePoolResponse> {
    const data = MsgCreateCreatePool.encode(request).finish();
    const promise = this.rpc.request("joltify.vault.Msg", "CreateCreatePool", data);
    return promise.then(data => MsgCreateCreatePoolResponse.decode(new _m0.Reader(data)));
  }

}