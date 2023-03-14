import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateIssueToken } from "./types/joltify/vault/tx";
import { MsgCreateCreatePool } from "./types/joltify/vault/tx";
import { MsgCreateOutboundTx } from "./types/joltify/vault/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.vault.MsgCreateIssueToken", MsgCreateIssueToken],
    ["/joltify.vault.MsgCreateCreatePool", MsgCreateCreatePool],
    ["/joltify.vault.MsgCreateOutboundTx", MsgCreateOutboundTx],
    
];

export { msgTypes }