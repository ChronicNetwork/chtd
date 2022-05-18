// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgStoreCode } from "./types/cht/tx";
import { MsgIBCSend } from "./types/cht/ibc";
import { MsgInstantiateContract } from "./types/cht/tx";
import { MsgUpdateAdmin } from "./types/cht/tx";
import { MsgIBCCloseChannel } from "./types/cht/ibc";
import { MsgMigrateContract } from "./types/cht/tx";
import { MsgExecuteContract } from "./types/cht/tx";
import { MsgClearAdmin } from "./types/cht/tx";


const types = [
  ["/ChronicNetwork.cht.cht.MsgStoreCode", MsgStoreCode],
  ["/ChronicNetwork.cht.cht.MsgIBCSend", MsgIBCSend],
  ["/ChronicNetwork.cht.cht.MsgInstantiateContract", MsgInstantiateContract],
  ["/ChronicNetwork.cht.cht.MsgUpdateAdmin", MsgUpdateAdmin],
  ["/ChronicNetwork.cht.cht.MsgIBCCloseChannel", MsgIBCCloseChannel],
  ["/ChronicNetwork.cht.cht.MsgMigrateContract", MsgMigrateContract],
  ["/ChronicNetwork.cht.cht.MsgExecuteContract", MsgExecuteContract],
  ["/ChronicNetwork.cht.cht.MsgClearAdmin", MsgClearAdmin],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgStoreCode: (data: MsgStoreCode): EncodeObject => ({ typeUrl: "/ChronicNetwork.cht.cht.MsgStoreCode", value: MsgStoreCode.fromPartial( data ) }),
    msgIBCSend: (data: MsgIBCSend): EncodeObject => ({ typeUrl: "/ChronicNetwork.cht.cht.MsgIBCSend", value: MsgIBCSend.fromPartial( data ) }),
    msgInstantiateContract: (data: MsgInstantiateContract): EncodeObject => ({ typeUrl: "/ChronicNetwork.cht.cht.MsgInstantiateContract", value: MsgInstantiateContract.fromPartial( data ) }),
    msgUpdateAdmin: (data: MsgUpdateAdmin): EncodeObject => ({ typeUrl: "/ChronicNetwork.cht.cht.MsgUpdateAdmin", value: MsgUpdateAdmin.fromPartial( data ) }),
    msgIBCCloseChannel: (data: MsgIBCCloseChannel): EncodeObject => ({ typeUrl: "/ChronicNetwork.cht.cht.MsgIBCCloseChannel", value: MsgIBCCloseChannel.fromPartial( data ) }),
    msgMigrateContract: (data: MsgMigrateContract): EncodeObject => ({ typeUrl: "/ChronicNetwork.cht.cht.MsgMigrateContract", value: MsgMigrateContract.fromPartial( data ) }),
    msgExecuteContract: (data: MsgExecuteContract): EncodeObject => ({ typeUrl: "/ChronicNetwork.cht.cht.MsgExecuteContract", value: MsgExecuteContract.fromPartial( data ) }),
    msgClearAdmin: (data: MsgClearAdmin): EncodeObject => ({ typeUrl: "/ChronicNetwork.cht.cht.MsgClearAdmin", value: MsgClearAdmin.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
