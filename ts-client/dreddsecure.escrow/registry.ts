import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCancelEscrow } from "./types/dreddsecure/escrow/tx";
import { MsgFulfillEscrow } from "./types/dreddsecure/escrow/tx";
import { MsgCreateEscrow } from "./types/dreddsecure/escrow/tx";
import { MsgOptOutEscrow } from "./types/dreddsecure/escrow/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/dreddsecure.escrow.MsgCancelEscrow", MsgCancelEscrow],
    ["/dreddsecure.escrow.MsgFulfillEscrow", MsgFulfillEscrow],
    ["/dreddsecure.escrow.MsgCreateEscrow", MsgCreateEscrow],
    ["/dreddsecure.escrow.MsgOptOutEscrow", MsgOptOutEscrow],
    
];

export { msgTypes }