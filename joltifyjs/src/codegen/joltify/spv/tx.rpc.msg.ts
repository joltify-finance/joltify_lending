import { Rpc } from "../../helpers";
import * as _m0 from "protobufjs/minimal";
import { MsgCreatePool, MsgCreatePoolResponse, MsgAddInvestors, MsgAddInvestorsResponse, MsgDeposit, MsgDepositResponse, MsgBorrow, MsgBorrowResponse, MsgRepayInterest, MsgRepayInterestResponse, MsgClaimInterest, MsgClaimInterestResponse, MsgUpdatePool, MsgUpdatePoolResponse, MsgActivePool, MsgActivePoolResponse, MsgPayPrincipal, MsgPayPrincipalResponse, MsgWithdrawPrincipal, MsgWithdrawPrincipalResponse, MsgSubmitWithdrawProposal, MsgSubmitWithdrawProposalResponse, MsgTransferOwnership, MsgTransferOwnershipResponse } from "./tx";
/** Msg defines the Msg service. */

export interface Msg {
  createPool(request: MsgCreatePool): Promise<MsgCreatePoolResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */

  addInvestors(request: MsgAddInvestors): Promise<MsgAddInvestorsResponse>;
  deposit(request: MsgDeposit): Promise<MsgDepositResponse>;
  borrow(request: MsgBorrow): Promise<MsgBorrowResponse>;
  repayInterest(request: MsgRepayInterest): Promise<MsgRepayInterestResponse>;
  claimInterest(request: MsgClaimInterest): Promise<MsgClaimInterestResponse>;
  updatePool(request: MsgUpdatePool): Promise<MsgUpdatePoolResponse>;
  activePool(request: MsgActivePool): Promise<MsgActivePoolResponse>;
  payPrincipal(request: MsgPayPrincipal): Promise<MsgPayPrincipalResponse>;
  withdrawPrincipal(request: MsgWithdrawPrincipal): Promise<MsgWithdrawPrincipalResponse>;
  submitWithdrawProposal(request: MsgSubmitWithdrawProposal): Promise<MsgSubmitWithdrawProposalResponse>;
  transferOwnership(request: MsgTransferOwnership): Promise<MsgTransferOwnershipResponse>;
}
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.createPool = this.createPool.bind(this);
    this.addInvestors = this.addInvestors.bind(this);
    this.deposit = this.deposit.bind(this);
    this.borrow = this.borrow.bind(this);
    this.repayInterest = this.repayInterest.bind(this);
    this.claimInterest = this.claimInterest.bind(this);
    this.updatePool = this.updatePool.bind(this);
    this.activePool = this.activePool.bind(this);
    this.payPrincipal = this.payPrincipal.bind(this);
    this.withdrawPrincipal = this.withdrawPrincipal.bind(this);
    this.submitWithdrawProposal = this.submitWithdrawProposal.bind(this);
    this.transferOwnership = this.transferOwnership.bind(this);
  }

  createPool(request: MsgCreatePool): Promise<MsgCreatePoolResponse> {
    const data = MsgCreatePool.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "CreatePool", data);
    return promise.then(data => MsgCreatePoolResponse.decode(new _m0.Reader(data)));
  }

  addInvestors(request: MsgAddInvestors): Promise<MsgAddInvestorsResponse> {
    const data = MsgAddInvestors.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "AddInvestors", data);
    return promise.then(data => MsgAddInvestorsResponse.decode(new _m0.Reader(data)));
  }

  deposit(request: MsgDeposit): Promise<MsgDepositResponse> {
    const data = MsgDeposit.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "Deposit", data);
    return promise.then(data => MsgDepositResponse.decode(new _m0.Reader(data)));
  }

  borrow(request: MsgBorrow): Promise<MsgBorrowResponse> {
    const data = MsgBorrow.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "Borrow", data);
    return promise.then(data => MsgBorrowResponse.decode(new _m0.Reader(data)));
  }

  repayInterest(request: MsgRepayInterest): Promise<MsgRepayInterestResponse> {
    const data = MsgRepayInterest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "RepayInterest", data);
    return promise.then(data => MsgRepayInterestResponse.decode(new _m0.Reader(data)));
  }

  claimInterest(request: MsgClaimInterest): Promise<MsgClaimInterestResponse> {
    const data = MsgClaimInterest.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "ClaimInterest", data);
    return promise.then(data => MsgClaimInterestResponse.decode(new _m0.Reader(data)));
  }

  updatePool(request: MsgUpdatePool): Promise<MsgUpdatePoolResponse> {
    const data = MsgUpdatePool.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "UpdatePool", data);
    return promise.then(data => MsgUpdatePoolResponse.decode(new _m0.Reader(data)));
  }

  activePool(request: MsgActivePool): Promise<MsgActivePoolResponse> {
    const data = MsgActivePool.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "ActivePool", data);
    return promise.then(data => MsgActivePoolResponse.decode(new _m0.Reader(data)));
  }

  payPrincipal(request: MsgPayPrincipal): Promise<MsgPayPrincipalResponse> {
    const data = MsgPayPrincipal.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "PayPrincipal", data);
    return promise.then(data => MsgPayPrincipalResponse.decode(new _m0.Reader(data)));
  }

  withdrawPrincipal(request: MsgWithdrawPrincipal): Promise<MsgWithdrawPrincipalResponse> {
    const data = MsgWithdrawPrincipal.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "WithdrawPrincipal", data);
    return promise.then(data => MsgWithdrawPrincipalResponse.decode(new _m0.Reader(data)));
  }

  submitWithdrawProposal(request: MsgSubmitWithdrawProposal): Promise<MsgSubmitWithdrawProposalResponse> {
    const data = MsgSubmitWithdrawProposal.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "SubmitWithdrawProposal", data);
    return promise.then(data => MsgSubmitWithdrawProposalResponse.decode(new _m0.Reader(data)));
  }

  transferOwnership(request: MsgTransferOwnership): Promise<MsgTransferOwnershipResponse> {
    const data = MsgTransferOwnership.encode(request).finish();
    const promise = this.rpc.request("joltify.spv.Msg", "TransferOwnership", data);
    return promise.then(data => MsgTransferOwnershipResponse.decode(new _m0.Reader(data)));
  }

}