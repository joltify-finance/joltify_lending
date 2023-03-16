import { GeneratedType, Registry, OfflineSigner } from "@cosmjs/proto-signing";
import { defaultRegistryTypes, AminoTypes, SigningStargateClient } from "@cosmjs/stargate";
import { HttpEndpoint } from "@cosmjs/tendermint-rpc";
import * as joltifyKycTxRegistry from "../joltify/kyc/tx.registry";
import * as joltifyKycTxAmino from "../joltify/kyc/tx.amino";
export const joltifyfinanceAminoConverters = { ...joltifyKycTxAmino.AminoConverter
};
export const joltifyfinanceProtoRegistry: ReadonlyArray<[string, GeneratedType]> = [...joltifyKycTxRegistry.registry];
export const getSigningJoltifyfinanceClientOptions = ({
  defaultTypes = defaultRegistryTypes
}: {
  defaultTypes?: ReadonlyArray<[string, GeneratedType]>;
} = {}): {
  registry: Registry;
  aminoTypes: AminoTypes;
} => {
  const registry = new Registry([...defaultTypes, ...joltifyfinanceProtoRegistry]);
  const aminoTypes = new AminoTypes({ ...joltifyfinanceAminoConverters
  });
  return {
    registry,
    aminoTypes
  };
};
export const getSigningJoltifyfinanceClient = async ({
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
  } = getSigningJoltifyfinanceClientOptions({
    defaultTypes
  });
  const client = await SigningStargateClient.connectWithSigner(rpcEndpoint, signer, {
    registry,
    aminoTypes
  });
  return client;
};