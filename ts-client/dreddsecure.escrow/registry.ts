import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgFulfillEscrow } from "./types/dreddsecure/escrow/tx";
import { MsgCancelEscrow } from "./types/dreddsecure/escrow/tx";
import { MsgCreateEscrow } from "./types/dreddsecure/escrow/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dreddsecure.escrow.MsgFulfillEscrow", MsgFulfillEscrow],
    ["/dreddsecure.escrow.MsgCancelEscrow", MsgCancelEscrow],
    ["/dreddsecure.escrow.MsgCreateEscrow", MsgCreateEscrow],
    
];

export { msgTypes }