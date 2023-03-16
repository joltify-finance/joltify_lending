import { GeneratedType, Registry, OfflineSigner } from "@cosmjs/proto-signing";
import { defaultRegistryTypes, AminoTypes, SigningStargateClient } from "@cosmjs/stargate";
import { HttpEndpoint } from "@cosmjs/tendermint-rpc";
import * as joltifySpvTxRegistry from "./spv/tx.registry";
import * as joltifyThirdPartyAuctionV1beta1TxRegistry from "./third_party/auction/v1beta1/tx.registry";
import * as joltifyThirdPartyCdpV1beta1TxRegistry from "./third_party/cdp/v1beta1/tx.registry";
import * as joltifyThirdPartyIncentiveV1beta1TxRegistry from "./third_party/incentive/v1beta1/tx.registry";
import * as joltifyThirdPartyIssuanceV1beta1TxRegistry from "./third_party/issuance/v1beta1/tx.registry";
import * as joltifyThirdPartyJoltV1beta1TxRegistry from "./third_party/jolt/v1beta1/tx.registry";
import * as joltifyThirdPartyPricefeedV1beta1TxRegistry from "./third_party/pricefeed/v1beta1/tx.registry";
import * as joltifyVaultTxRegistry from "./vault/tx.registry";
import * as joltifySpvTxAmino from "./spv/tx.amino";
import * as joltifyThirdPartyAuctionV1beta1TxAmino from "./third_party/auction/v1beta1/tx.amino";
import * as joltifyThirdPartyCdpV1beta1TxAmino from "./third_party/cdp/v1beta1/tx.amino";
import * as joltifyThirdPartyIncentiveV1beta1TxAmino from "./third_party/incentive/v1beta1/tx.amino";
import * as joltifyThirdPartyIssuanceV1beta1TxAmino from "./third_party/issuance/v1beta1/tx.amino";
import * as joltifyThirdPartyJoltV1beta1TxAmino from "./third_party/jolt/v1beta1/tx.amino";
import * as joltifyThirdPartyPricefeedV1beta1TxAmino from "./third_party/pricefeed/v1beta1/tx.amino";
import * as joltifyVaultTxAmino from "./vault/tx.amino";
export const joltifyAminoConverters = { ...joltifySpvTxAmino.AminoConverter,
  ...joltifyThirdPartyAuctionV1beta1TxAmino.AminoConverter,
  ...joltifyThirdPartyCdpV1beta1TxAmino.AminoConverter,
  ...joltifyThirdPartyIncentiveV1beta1TxAmino.AminoConverter,
  ...joltifyThirdPartyIssuanceV1beta1TxAmino.AminoConverter,
  ...joltifyThirdPartyJoltV1beta1TxAmino.AminoConverter,
  ...joltifyThirdPartyPricefeedV1beta1TxAmino.AminoConverter,
  ...joltifyVaultTxAmino.AminoConverter
};
export const joltifyProtoRegistry: ReadonlyArray<[string, GeneratedType]> = [...joltifySpvTxRegistry.registry, ...joltifyThirdPartyAuctionV1beta1TxRegistry.registry, ...joltifyThirdPartyCdpV1beta1TxRegistry.registry, ...joltifyThirdPartyIncentiveV1beta1TxRegistry.registry, ...joltifyThirdPartyIssuanceV1beta1TxRegistry.registry, ...joltifyThirdPartyJoltV1beta1TxRegistry.registry, ...joltifyThirdPartyPricefeedV1beta1TxRegistry.registry, ...joltifyVaultTxRegistry.registry];
export const getSigningJoltifyClientOptions = ({
  defaultTypes = defaultRegistryTypes
}: {
  defaultTypes?: ReadonlyArray<[string, GeneratedType]>;
} = {}): {
  registry: Registry;
  aminoTypes: AminoTypes;
} => {
  const registry = new Registry([...defaultTypes, ...joltifyProtoRegistry]);
  const aminoTypes = new AminoTypes({ ...joltifyAminoConverters
  });
  return {
    registry,
    aminoTypes
  };
};
export const getSigningJoltifyClient = async ({
  rpcEndpoint,
  signer,
  defaultTypes = defaultRegistryTypes
}: {
  rpcEndpoint: string | HttpEndpoint;
  signer: OfflineSigner;
  defaultTypes?: ReadonlyArray<[string, GeneratedType]>;
}) => {
  const {
    registry,
    aminoTypes
  } = getSigningJoltifyClientOptions({
    defaultTypes
  });
  const client = await SigningStargateClient.connectWithSigner(rpcEndpoint, signer, {
    registry,
    aminoTypes
  });
  return client;
};