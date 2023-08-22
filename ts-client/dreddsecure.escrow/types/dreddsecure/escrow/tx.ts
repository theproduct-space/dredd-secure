/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "dreddsecure.escrow";

export interface MsgCreateEscrow {
  creator: string;
  initiatorCoins: Coin[];
  fulfillerCoins: Coin[];
  tips: Coin[];
  startDate: string;
  endDate: string;
  apiConditions: string;
}

export interface MsgCreateEscrowResponse {
}

export interface MsgCancelEscrow {
  creator: string;
  id: number;
}

export interface MsgCancelEscrowResponse {
}

export interface MsgFulfillEscrow {
  creator: string;
  id: number;
}

export interface MsgFulfillEscrowResponse {
}

export interface MsgOptOutEscrow {
  creator: string;
  id: number;
}

export interface MsgOptOutEscrowResponse {
}

function createBaseMsgCreateEscrow(): MsgCreateEscrow {
  return {
    creator: "",
    initiatorCoins: [],
    fulfillerCoins: [],
    tips: [],
    startDate: "",
    endDate: "",
    apiConditions: "",
  };
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
    for (const v of message.tips) {
      Coin.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.startDate !== "") {
      writer.uint32(42).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(50).string(message.endDate);
    }
    if (message.apiConditions !== "") {
      writer.uint32(58).string(message.apiConditions);
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
          message.tips.push(Coin.decode(reader, reader.uint32()));
          break;
        case 5:
          message.startDate = reader.string();
          break;
        case 6:
          message.endDate = reader.string();
          break;
        case 7:
          message.apiConditions = reader.string();
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
      tips: Array.isArray(object?.tips) ? object.tips.map((e: any) => Coin.fromJSON(e)) : [],
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
      apiConditions: isSet(object.apiConditions) ? String(object.apiConditions) : "",
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
    if (message.tips) {
      obj.tips = message.tips.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.tips = [];
    }
    message.startDate !== undefined && (obj.startDate = message.startDate);
    message.endDate !== undefined && (obj.endDate = message.endDate);
    message.apiConditions !== undefined && (obj.apiConditions = message.apiConditions);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateEscrow>, I>>(object: I): MsgCreateEscrow {
    const message = createBaseMsgCreateEscrow();
    message.creator = object.creator ?? "";
    message.initiatorCoins = object.initiatorCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.fulfillerCoins = object.fulfillerCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.tips = object.tips?.map((e) => Coin.fromPartial(e)) || [];
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    message.apiConditions = object.apiConditions ?? "";
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

function createBaseMsgCancelEscrow(): MsgCancelEscrow {
  return { creator: "", id: 0 };
}

export const MsgCancelEscrow = {
  encode(message: MsgCancelEscrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelEscrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelEscrow();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelEscrow {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgCancelEscrow): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelEscrow>, I>>(object: I): MsgCancelEscrow {
    const message = createBaseMsgCancelEscrow();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgCancelEscrowResponse(): MsgCancelEscrowResponse {
  return {};
}

export const MsgCancelEscrowResponse = {
  encode(_: MsgCancelEscrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelEscrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelEscrowResponse();
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

  fromJSON(_: any): MsgCancelEscrowResponse {
    return {};
  },

  toJSON(_: MsgCancelEscrowResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCancelEscrowResponse>, I>>(_: I): MsgCancelEscrowResponse {
    const message = createBaseMsgCancelEscrowResponse();
    return message;
  },
};

function createBaseMsgFulfillEscrow(): MsgFulfillEscrow {
  return { creator: "", id: 0 };
}

export const MsgFulfillEscrow = {
  encode(message: MsgFulfillEscrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgFulfillEscrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgFulfillEscrow();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFulfillEscrow {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgFulfillEscrow): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgFulfillEscrow>, I>>(object: I): MsgFulfillEscrow {
    const message = createBaseMsgFulfillEscrow();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgFulfillEscrowResponse(): MsgFulfillEscrowResponse {
  return {};
}

export const MsgFulfillEscrowResponse = {
  encode(_: MsgFulfillEscrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgFulfillEscrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgFulfillEscrowResponse();
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

  fromJSON(_: any): MsgFulfillEscrowResponse {
    return {};
  },

  toJSON(_: MsgFulfillEscrowResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgFulfillEscrowResponse>, I>>(_: I): MsgFulfillEscrowResponse {
    const message = createBaseMsgFulfillEscrowResponse();
    return message;
  },
};

function createBaseMsgOptOutEscrow(): MsgOptOutEscrow {
  return { creator: "", id: 0 };
}

export const MsgOptOutEscrow = {
  encode(message: MsgOptOutEscrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgOptOutEscrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgOptOutEscrow();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgOptOutEscrow {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgOptOutEscrow): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgOptOutEscrow>, I>>(object: I): MsgOptOutEscrow {
    const message = createBaseMsgOptOutEscrow();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgOptOutEscrowResponse(): MsgOptOutEscrowResponse {
  return {};
}

export const MsgOptOutEscrowResponse = {
  encode(_: MsgOptOutEscrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgOptOutEscrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgOptOutEscrowResponse();
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

  fromJSON(_: any): MsgOptOutEscrowResponse {
    return {};
  },

  toJSON(_: MsgOptOutEscrowResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgOptOutEscrowResponse>, I>>(_: I): MsgOptOutEscrowResponse {
    const message = createBaseMsgOptOutEscrowResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateEscrow(request: MsgCreateEscrow): Promise<MsgCreateEscrowResponse>;
  CancelEscrow(request: MsgCancelEscrow): Promise<MsgCancelEscrowResponse>;
  FulfillEscrow(request: MsgFulfillEscrow): Promise<MsgFulfillEscrowResponse>;
  OptOutEscrow(request: MsgOptOutEscrow): Promise<MsgOptOutEscrowResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateEscrow = this.CreateEscrow.bind(this);
    this.CancelEscrow = this.CancelEscrow.bind(this);
    this.FulfillEscrow = this.FulfillEscrow.bind(this);
    this.OptOutEscrow = this.OptOutEscrow.bind(this);
  }
  CreateEscrow(request: MsgCreateEscrow): Promise<MsgCreateEscrowResponse> {
    const data = MsgCreateEscrow.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Msg", "CreateEscrow", data);
    return promise.then((data) => MsgCreateEscrowResponse.decode(new _m0.Reader(data)));
  }

  CancelEscrow(request: MsgCancelEscrow): Promise<MsgCancelEscrowResponse> {
    const data = MsgCancelEscrow.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Msg", "CancelEscrow", data);
    return promise.then((data) => MsgCancelEscrowResponse.decode(new _m0.Reader(data)));
  }

  FulfillEscrow(request: MsgFulfillEscrow): Promise<MsgFulfillEscrowResponse> {
    const data = MsgFulfillEscrow.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Msg", "FulfillEscrow", data);
    return promise.then((data) => MsgFulfillEscrowResponse.decode(new _m0.Reader(data)));
  }

  OptOutEscrow(request: MsgOptOutEscrow): Promise<MsgOptOutEscrowResponse> {
    const data = MsgOptOutEscrow.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Msg", "OptOutEscrow", data);
    return promise.then((data) => MsgOptOutEscrowResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
