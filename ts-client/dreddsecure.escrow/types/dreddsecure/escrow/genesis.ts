/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Escrow } from "./escrow";
import { Params } from "./params";

export const protobufPackage = "dreddsecure.escrow";

/** GenesisState defines the escrow module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  escrowList: Escrow[];
  escrowCount: number;
  pendingEscrows: number[];
  expiringEscrows: number[];
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, escrowList: [], escrowCount: 0, pendingEscrows: [], expiringEscrows: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.escrowList) {
      Escrow.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.escrowCount !== 0) {
      writer.uint32(24).uint64(message.escrowCount);
    }
    writer.uint32(34).fork();
    for (const v of message.pendingEscrows) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(42).fork();
    for (const v of message.expiringEscrows) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.escrowList.push(Escrow.decode(reader, reader.uint32()));
          break;
        case 3:
          message.escrowCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.pendingEscrows.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.pendingEscrows.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 5:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.expiringEscrows.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.expiringEscrows.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      escrowList: Array.isArray(object?.escrowList) ? object.escrowList.map((e: any) => Escrow.fromJSON(e)) : [],
      escrowCount: isSet(object.escrowCount) ? Number(object.escrowCount) : 0,
      pendingEscrows: Array.isArray(object?.pendingEscrows) ? object.pendingEscrows.map((e: any) => Number(e)) : [],
      expiringEscrows: Array.isArray(object?.expiringEscrows) ? object.expiringEscrows.map((e: any) => Number(e)) : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.escrowList) {
      obj.escrowList = message.escrowList.map((e) => e ? Escrow.toJSON(e) : undefined);
    } else {
      obj.escrowList = [];
    }
    message.escrowCount !== undefined && (obj.escrowCount = Math.round(message.escrowCount));
    if (message.pendingEscrows) {
      obj.pendingEscrows = message.pendingEscrows.map((e) => Math.round(e));
    } else {
      obj.pendingEscrows = [];
    }
    if (message.expiringEscrows) {
      obj.expiringEscrows = message.expiringEscrows.map((e) => Math.round(e));
    } else {
      obj.expiringEscrows = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.escrowList = object.escrowList?.map((e) => Escrow.fromPartial(e)) || [];
    message.escrowCount = object.escrowCount ?? 0;
    message.pendingEscrows = object.pendingEscrows?.map((e) => e) || [];
    message.expiringEscrows = object.expiringEscrows?.map((e) => e) || [];
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
