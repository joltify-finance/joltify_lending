import { GeneratedType, Registry } from "@cosmjs/proto-signing";
import { MsgCreateOutboundTx, MsgCreateIssueToken, MsgCreateCreatePool } from "./tx";
export const registry: ReadonlyArray<[string, GeneratedType]> = [["/joltify.vault.MsgCreateOutboundTx", MsgCreateOutboundTx], ["/joltify.vault.MsgCreateIssueToken", MsgCreateIssueToken], ["/joltify.vault.MsgCreateCreatePool", MsgCreateCreatePool]];
export const load = (protoRegistry: Registry) => {
  registry.forEach(([typeUrl, mod]) => {
    protoRegistry.register(typeUrl, mod);
  });
};
export const MessageComposer = {
  encoded: {
    createOutboundTx(value: MsgCreateOutboundTx) {
      return {
        typeUrl: "/joltify.vault.MsgCreateOutboundTx",
        value: MsgCreateOutboundTx.encode(value).finish()
      };
    },

    createIssueToken(value: MsgCreateIssueToken) {
      return {
        typeUrl: "/joltify.vault.MsgCreateIssueToken",
        value: MsgCreateIssueToken.encode(value).finish()
      };
    },

    createCreatePool(value: MsgCreateCreatePool) {
      return {
        typeUrl: "/joltify.vault.MsgCreateCreatePool",
        value: MsgCreateCreatePool.encode(value).finish()
      };
    }

  },
  withTypeUrl: {
    createOutboundTx(value: MsgCreateOutboundTx) {
      return {
        typeUrl: "/joltify.vault.MsgCreateOutboundTx",
        value
      };
    },

    createIssueToken(value: MsgCreateIssueToken) {
      return {
        typeUrl: "/joltify.vault.MsgCreateIssueToken",
        value
      };
    },

    createCreatePool(value: MsgCreateCreatePool) {
      return {
        typeUrl: "/joltify.vault.MsgCreateCreatePool",
        value
      };
    }

  },
  fromPartial: {
    createOutboundTx(value: MsgCreateOutboundTx) {
      return {
        typeUrl: "/joltify.vault.MsgCreateOutboundTx",
        value: MsgCreateOutboundTx.fromPartial(value)
      };
    },

    createIssueToken(value: MsgCreateIssueToken) {
      return {
        typeUrl: "/joltify.vault.MsgCreateIssueToken",
        value: MsgCreateIssueToken.fromPartial(value)
      };
    },

    createCreatePool(value: MsgCreateCreatePool) {
      return {
        typeUrl: "/joltify.vault.MsgCreateCreatePool",
        value: MsgCreateCreatePool.fromPartial(value)
      };
    }

  }
};