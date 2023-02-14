import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateCreatePool } from "./types/joltify/vault/tx";
import { MsgCreateOutboundTx } from "./types/joltify/vault/tx";
import { MsgCreateIssueToken } from "./types/joltify/vault/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/joltify.vault.MsgCreateCreatePool", MsgCreateCreatePool],
    ["/joltify.vault.MsgCreateOutboundTx", MsgCreateOutboundTx],
    ["/joltify.vault.MsgCreateIssueToken", MsgCreateIssueToken],
    
];

export { msgTypes }