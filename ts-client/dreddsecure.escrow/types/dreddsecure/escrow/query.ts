/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Escrow } from "./escrow";
import { Params } from "./params";

export const protobufPackage = "dreddsecure.escrow";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetEscrowRequest {
  id: number;
}

export interface QueryGetEscrowResponse {
  Escrow: Escrow | undefined;
}

export interface QueryAllEscrowRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEscrowResponse {
  Escrow: Escrow[];
  pagination: PageResponse | undefined;
}

export interface QueryEscrowsByAddressRequest {
  address: string;
  pagination: PageRequest | undefined;
}

export interface QueryEscrowsByAddressResponse {
  Escrow: Escrow[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetEscrowRequest(): QueryGetEscrowRequest {
  return { id: 0 };
}

export const QueryGetEscrowRequest = {
  encode(message: QueryGetEscrowRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEscrowRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEscrowRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEscrowRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetEscrowRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEscrowRequest>, I>>(object: I): QueryGetEscrowRequest {
    const message = createBaseQueryGetEscrowRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetEscrowResponse(): QueryGetEscrowResponse {
  return { Escrow: undefined };
}

export const QueryGetEscrowResponse = {
  encode(message: QueryGetEscrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Escrow !== undefined) {
      Escrow.encode(message.Escrow, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEscrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEscrowResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Escrow = Escrow.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEscrowResponse {
    return { Escrow: isSet(object.Escrow) ? Escrow.fromJSON(object.Escrow) : undefined };
  },

  toJSON(message: QueryGetEscrowResponse): unknown {
    const obj: any = {};
    message.Escrow !== undefined && (obj.Escrow = message.Escrow ? Escrow.toJSON(message.Escrow) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEscrowResponse>, I>>(object: I): QueryGetEscrowResponse {
    const message = createBaseQueryGetEscrowResponse();
    message.Escrow = (object.Escrow !== undefined && object.Escrow !== null)
      ? Escrow.fromPartial(object.Escrow)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEscrowRequest(): QueryAllEscrowRequest {
  return { pagination: undefined };
}

export const QueryAllEscrowRequest = {
  encode(message: QueryAllEscrowRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEscrowRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEscrowRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEscrowRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllEscrowRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEscrowRequest>, I>>(object: I): QueryAllEscrowRequest {
    const message = createBaseQueryAllEscrowRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEscrowResponse(): QueryAllEscrowResponse {
  return { Escrow: [], pagination: undefined };
}

export const QueryAllEscrowResponse = {
  encode(message: QueryAllEscrowResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Escrow) {
      Escrow.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEscrowResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEscrowResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Escrow.push(Escrow.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEscrowResponse {
    return {
      Escrow: Array.isArray(object?.Escrow) ? object.Escrow.map((e: any) => Escrow.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllEscrowResponse): unknown {
    const obj: any = {};
    if (message.Escrow) {
      obj.Escrow = message.Escrow.map((e) => e ? Escrow.toJSON(e) : undefined);
    } else {
      obj.Escrow = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEscrowResponse>, I>>(object: I): QueryAllEscrowResponse {
    const message = createBaseQueryAllEscrowResponse();
    message.Escrow = object.Escrow?.map((e) => Escrow.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryEscrowsByAddressRequest(): QueryEscrowsByAddressRequest {
  return { address: "", pagination: undefined };
}

export const QueryEscrowsByAddressRequest = {
  encode(message: QueryEscrowsByAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryEscrowsByAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryEscrowsByAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEscrowsByAddressRequest {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryEscrowsByAddressRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryEscrowsByAddressRequest>, I>>(object: I): QueryEscrowsByAddressRequest {
    const message = createBaseQueryEscrowsByAddressRequest();
    message.address = object.address ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryEscrowsByAddressResponse(): QueryEscrowsByAddressResponse {
  return { Escrow: [], pagination: undefined };
}

export const QueryEscrowsByAddressResponse = {
  encode(message: QueryEscrowsByAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Escrow) {
      Escrow.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryEscrowsByAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryEscrowsByAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Escrow.push(Escrow.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEscrowsByAddressResponse {
    return {
      Escrow: Array.isArray(object?.Escrow) ? object.Escrow.map((e: any) => Escrow.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryEscrowsByAddressResponse): unknown {
    const obj: any = {};
    if (message.Escrow) {
      obj.Escrow = message.Escrow.map((e) => e ? Escrow.toJSON(e) : undefined);
    } else {
      obj.Escrow = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryEscrowsByAddressResponse>, I>>(
    object: I,
  ): QueryEscrowsByAddressResponse {
    const message = createBaseQueryEscrowsByAddressResponse();
    message.Escrow = object.Escrow?.map((e) => Escrow.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Escrow items. */
  Escrow(request: QueryGetEscrowRequest): Promise<QueryGetEscrowResponse>;
  EscrowAll(request: QueryAllEscrowRequest): Promise<QueryAllEscrowResponse>;
  /** Queries a list of EscrowsByAddress items. */
  EscrowsByAddress(request: QueryEscrowsByAddressRequest): Promise<QueryEscrowsByAddressResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Escrow = this.Escrow.bind(this);
    this.EscrowAll = this.EscrowAll.bind(this);
    this.EscrowsByAddress = this.EscrowsByAddress.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Escrow(request: QueryGetEscrowRequest): Promise<QueryGetEscrowResponse> {
    const data = QueryGetEscrowRequest.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Query", "Escrow", data);
    return promise.then((data) => QueryGetEscrowResponse.decode(new _m0.Reader(data)));
  }

  EscrowAll(request: QueryAllEscrowRequest): Promise<QueryAllEscrowResponse> {
    const data = QueryAllEscrowRequest.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Query", "EscrowAll", data);
    return promise.then((data) => QueryAllEscrowResponse.decode(new _m0.Reader(data)));
  }

  EscrowsByAddress(request: QueryEscrowsByAddressRequest): Promise<QueryEscrowsByAddressResponse> {
    const data = QueryEscrowsByAddressRequest.encode(request).finish();
    const promise = this.rpc.request("dreddsecure.escrow.Query", "EscrowsByAddress", data);
    return promise.then((data) => QueryEscrowsByAddressResponse.decode(new _m0.Reader(data)));
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
