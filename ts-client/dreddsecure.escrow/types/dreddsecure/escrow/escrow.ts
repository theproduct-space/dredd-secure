/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "dreddsecure.escrow";

export interface Escrow {
  id: number;
  status: string;
  initiator: string;
  fulfiller: string;
  initiatorCoins: Coin[];
  fulfillerCoins: Coin[];
  startDate: string;
  endDate: string;
}

function createBaseEscrow(): Escrow {
  return {
    id: 0,
    status: "",
    initiator: "",
    fulfiller: "",
    initiatorCoins: [],
    fulfillerCoins: [],
    startDate: "",
    endDate: "",
  };
}

export const Escrow = {
  encode(message: Escrow, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.status !== "") {
      writer.uint32(18).string(message.status);
    }
    if (message.initiator !== "") {
      writer.uint32(26).string(message.initiator);
    }
    if (message.fulfiller !== "") {
      writer.uint32(34).string(message.fulfiller);
    }
    for (const v of message.initiatorCoins) {
      Coin.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.fulfillerCoins) {
      Coin.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.startDate !== "") {
      writer.uint32(58).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(66).string(message.endDate);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Escrow {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEscrow();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.status = reader.string();
          break;
        case 3:
          message.initiator = reader.string();
          break;
        case 4:
          message.fulfiller = reader.string();
          break;
        case 5:
          message.initiatorCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        case 6:
          message.fulfillerCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        case 7:
          message.startDate = reader.string();
          break;
        case 8:
          message.endDate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Escrow {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      status: isSet(object.status) ? String(object.status) : "",
      initiator: isSet(object.initiator) ? String(object.initiator) : "",
      fulfiller: isSet(object.fulfiller) ? String(object.fulfiller) : "",
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

  toJSON(message: Escrow): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.status !== undefined && (obj.status = message.status);
    message.initiator !== undefined && (obj.initiator = message.initiator);
    message.fulfiller !== undefined && (obj.fulfiller = message.fulfiller);
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

  fromPartial<I extends Exact<DeepPartial<Escrow>, I>>(object: I): Escrow {
    const message = createBaseEscrow();
    message.id = object.id ?? 0;
    message.status = object.status ?? "";
    message.initiator = object.initiator ?? "";
    message.fulfiller = object.fulfiller ?? "";
    message.initiatorCoins = object.initiatorCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.fulfillerCoins = object.fulfillerCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    return message;
  },
};

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
