import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateEscrow } from "./types/dreddsecure/escrow/tx";
import { MsgFulfillEscrow } from "./types/dreddsecure/escrow/tx";
import { MsgCancelEscrow } from "./types/dreddsecure/escrow/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dreddsecure.escrow.MsgCreateEscrow", MsgCreateEscrow],
    ["/dreddsecure.escrow.MsgFulfillEscrow", MsgFulfillEscrow],
    ["/dreddsecure.escrow.MsgCancelEscrow", MsgCancelEscrow],
    
];

export { msgTypes }