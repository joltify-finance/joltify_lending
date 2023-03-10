import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateOutboundTx } from "./types/joltify/vault/tx";
import { MsgCreateIssueToken } from "./types/joltify/vault/tx";
import { MsgCreateCreatePool } from "./types/joltify/vault/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.vault.MsgCreateOutboundTx", MsgCreateOutboundTx],
    ["/joltify.vault.MsgCreateIssueToken", MsgCreateIssueToken],
    ["/joltify.vault.MsgCreateCreatePool", MsgCreateCreatePool],
    
];

export { msgTypes }