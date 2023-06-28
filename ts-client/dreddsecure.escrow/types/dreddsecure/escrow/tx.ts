/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "dreddsecure.escrow";

export interface MsgCreateEscrow {
  creator: string;
  initiatorCoins: Coin[];
  fulfillerCoins: Coin[];
  startDate: string;
  endDate: string;
}

export interface MsgCreateEscrowResponse {
}

function createBaseMsgCreateEscrow(): MsgCreateEscrow {
  return { creator: "", initiatorCoins: [], fulfillerCoins: [], startDate: "", endDate: "" };
}

export const MsgCreateEscrow = {
  encode(message: MsgCreateEscrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    for (const v of message.initiatorCoins) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.fulfillerCoins) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.startDate !== "") {
      writer.uint32(34).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(42).string(message.endDate);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateEscrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateEscrow();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.initiatorCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        case 3:
          message.fulfillerCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          message.startDate = reader.string();
          break;
        case 5:
          message.endDate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateEscrow {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      initiatorCoins: Array.isArray(object?.initiatorCoins)
        ? object.initiatorCoins.map((e: any) => Coin.fromJSON(e))
        : [],
      fulfillerCoins: Array.isArray(object?.fulfillerCoins)
        ? object.fulfillerCoins.map((e: any) => Coin.fromJSON(e))
        : [],
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
    };
  },

  toJSON(message: MsgCreateEscrow): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.initiatorCoins) {
      obj.initiatorCoins = message.initiatorCoins.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.initiatorCoins = [];
    }
    if (message.fulfillerCoins) {
      obj.fulfillerCoins = message.fulfillerCoins.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.fulfillerCoins = [];
    }
    message.startDate !== undefined && (obj.startDate = message.startDate);
    message.endDate !== undefined && (obj.endDate = message.endDate);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateEscrow>, I>>(object: I): MsgCreateEscrow {
    const message = createBaseMsgCreateEscrow();
    message.creator = object.creator ?? "";
    message.initiatorCoins = object.initiatorCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.fulfillerCoins = object.fulfillerCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    return message;
  },
};

function createBaseMsgCreateEscrowResponse(): MsgCreateEscrowResponse {
  return {};
}

export const MsgCreateEscrowResponse = {
  encode(_: MsgCreateEscrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateEscrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateEscrowResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateEscrowResponse {
    return {};
  },

  toJSON(_: MsgCreateEscrowResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateEscrowResponse>, I>>(_: I): MsgCreateEscrowResponse {
    const message = createBaseMsgCreateEscrowResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateEscrow(request: MsgCreateEscrow): Promise<MsgCreateEscrowResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateEscrow = this.CreateEscrow.bind(this);
  }
  CreateEscrow(request: MsgCreateEscrow): Promise<MsgCreateEscrowResponse> {
    const data = MsgCreateEscrow.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Msg", "CreateEscrow", data);
    return promise.then((data) => MsgCreateEscrowResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
