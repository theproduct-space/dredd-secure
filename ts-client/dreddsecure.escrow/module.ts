// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgCreateEscrow } from "./types/dreddsecure/escrow/tx";
import { MsgFulfillEscrow } from "./types/dreddsecure/escrow/tx";

import { Escrow as typeEscrow} from "./types"
import { Params as typeParams} from "./types"

export { MsgCreateEscrow, MsgFulfillEscrow };

type sendMsgCreateEscrowParams = {
  value: MsgCreateEscrow,
  fee?: StdFee,
  memo?: string
};

type sendMsgFulfillEscrowParams = {
  value: MsgFulfillEscrow,
  fee?: StdFee,
  memo?: string
};


type msgCreateEscrowParams = {
  value: MsgCreateEscrow,
};

type msgFulfillEscrowParams = {
  value: MsgFulfillEscrow,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgCreateEscrow({ value, fee, memo }: sendMsgCreateEscrowParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateEscrow: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateEscrow({ value: MsgCreateEscrow.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateEscrow: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgFulfillEscrow({ value, fee, memo }: sendMsgFulfillEscrowParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgFulfillEscrow: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgFulfillEscrow({ value: MsgFulfillEscrow.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgFulfillEscrow: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgCreateEscrow({ value }: msgCreateEscrowParams): EncodeObject {
			try {
				return { typeUrl: "/dreddsecure.escrow.MsgCreateEscrow", value: MsgCreateEscrow.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateEscrow: Could not create message: ' + e.message)
			}
		},
		
		msgFulfillEscrow({ value }: msgFulfillEscrowParams): EncodeObject {
			try {
				return { typeUrl: "/dreddsecure.escrow.MsgFulfillEscrow", value: MsgFulfillEscrow.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgFulfillEscrow: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						Escrow: getStructure(typeEscrow.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			DreddsecureEscrow: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;