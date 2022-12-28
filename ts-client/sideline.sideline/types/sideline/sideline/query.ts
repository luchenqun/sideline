/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Developer } from "./developer";
import { Employer } from "./employer";
import { Params } from "./params";
import { Task } from "./task";

export const protobufPackage = "sideline.sideline";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetEmployerRequest {
  index: string;
}

export interface QueryGetEmployerResponse {
  employer: Employer | undefined;
}

export interface QueryAllEmployerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEmployerResponse {
  employer: Employer[];
  pagination: PageResponse | undefined;
}

export interface QueryGetDeveloperRequest {
  index: string;
}

export interface QueryGetDeveloperResponse {
  developer: Developer | undefined;
}

export interface QueryAllDeveloperRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDeveloperResponse {
  developer: Developer[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTaskRequest {
  id: number;
}

export interface QueryGetTaskResponse {
  Task: Task | undefined;
}

export interface QueryAllTaskRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTaskResponse {
  Task: Task[];
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

function createBaseQueryGetEmployerRequest(): QueryGetEmployerRequest {
  return { index: "" };
}

export const QueryGetEmployerRequest = {
  encode(message: QueryGetEmployerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEmployerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEmployerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEmployerRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetEmployerRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEmployerRequest>, I>>(object: I): QueryGetEmployerRequest {
    const message = createBaseQueryGetEmployerRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetEmployerResponse(): QueryGetEmployerResponse {
  return { employer: undefined };
}

export const QueryGetEmployerResponse = {
  encode(message: QueryGetEmployerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.employer !== undefined) {
      Employer.encode(message.employer, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEmployerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEmployerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.employer = Employer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEmployerResponse {
    return { employer: isSet(object.employer) ? Employer.fromJSON(object.employer) : undefined };
  },

  toJSON(message: QueryGetEmployerResponse): unknown {
    const obj: any = {};
    message.employer !== undefined && (obj.employer = message.employer ? Employer.toJSON(message.employer) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEmployerResponse>, I>>(object: I): QueryGetEmployerResponse {
    const message = createBaseQueryGetEmployerResponse();
    message.employer = (object.employer !== undefined && object.employer !== null)
      ? Employer.fromPartial(object.employer)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEmployerRequest(): QueryAllEmployerRequest {
  return { pagination: undefined };
}

export const QueryAllEmployerRequest = {
  encode(message: QueryAllEmployerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEmployerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEmployerRequest();
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

  fromJSON(object: any): QueryAllEmployerRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllEmployerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEmployerRequest>, I>>(object: I): QueryAllEmployerRequest {
    const message = createBaseQueryAllEmployerRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEmployerResponse(): QueryAllEmployerResponse {
  return { employer: [], pagination: undefined };
}

export const QueryAllEmployerResponse = {
  encode(message: QueryAllEmployerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.employer) {
      Employer.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEmployerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEmployerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.employer.push(Employer.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllEmployerResponse {
    return {
      employer: Array.isArray(object?.employer) ? object.employer.map((e: any) => Employer.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllEmployerResponse): unknown {
    const obj: any = {};
    if (message.employer) {
      obj.employer = message.employer.map((e) => e ? Employer.toJSON(e) : undefined);
    } else {
      obj.employer = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEmployerResponse>, I>>(object: I): QueryAllEmployerResponse {
    const message = createBaseQueryAllEmployerResponse();
    message.employer = object.employer?.map((e) => Employer.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDeveloperRequest(): QueryGetDeveloperRequest {
  return { index: "" };
}

export const QueryGetDeveloperRequest = {
  encode(message: QueryGetDeveloperRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDeveloperRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDeveloperRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDeveloperRequest {
    return { index: isSet(object.index) ? String(object.index) : "" };
  },

  toJSON(message: QueryGetDeveloperRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDeveloperRequest>, I>>(object: I): QueryGetDeveloperRequest {
    const message = createBaseQueryGetDeveloperRequest();
    message.index = object.index ?? "";
    return message;
  },
};

function createBaseQueryGetDeveloperResponse(): QueryGetDeveloperResponse {
  return { developer: undefined };
}

export const QueryGetDeveloperResponse = {
  encode(message: QueryGetDeveloperResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.developer !== undefined) {
      Developer.encode(message.developer, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDeveloperResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDeveloperResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.developer = Developer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDeveloperResponse {
    return { developer: isSet(object.developer) ? Developer.fromJSON(object.developer) : undefined };
  },

  toJSON(message: QueryGetDeveloperResponse): unknown {
    const obj: any = {};
    message.developer !== undefined
      && (obj.developer = message.developer ? Developer.toJSON(message.developer) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDeveloperResponse>, I>>(object: I): QueryGetDeveloperResponse {
    const message = createBaseQueryGetDeveloperResponse();
    message.developer = (object.developer !== undefined && object.developer !== null)
      ? Developer.fromPartial(object.developer)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDeveloperRequest(): QueryAllDeveloperRequest {
  return { pagination: undefined };
}

export const QueryAllDeveloperRequest = {
  encode(message: QueryAllDeveloperRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDeveloperRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDeveloperRequest();
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

  fromJSON(object: any): QueryAllDeveloperRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDeveloperRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDeveloperRequest>, I>>(object: I): QueryAllDeveloperRequest {
    const message = createBaseQueryAllDeveloperRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDeveloperResponse(): QueryAllDeveloperResponse {
  return { developer: [], pagination: undefined };
}

export const QueryAllDeveloperResponse = {
  encode(message: QueryAllDeveloperResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.developer) {
      Developer.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDeveloperResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDeveloperResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.developer.push(Developer.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllDeveloperResponse {
    return {
      developer: Array.isArray(object?.developer) ? object.developer.map((e: any) => Developer.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDeveloperResponse): unknown {
    const obj: any = {};
    if (message.developer) {
      obj.developer = message.developer.map((e) => e ? Developer.toJSON(e) : undefined);
    } else {
      obj.developer = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDeveloperResponse>, I>>(object: I): QueryAllDeveloperResponse {
    const message = createBaseQueryAllDeveloperResponse();
    message.developer = object.developer?.map((e) => Developer.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetTaskRequest(): QueryGetTaskRequest {
  return { id: 0 };
}

export const QueryGetTaskRequest = {
  encode(message: QueryGetTaskRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTaskRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTaskRequest();
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

  fromJSON(object: any): QueryGetTaskRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetTaskRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTaskRequest>, I>>(object: I): QueryGetTaskRequest {
    const message = createBaseQueryGetTaskRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryGetTaskResponse(): QueryGetTaskResponse {
  return { Task: undefined };
}

export const QueryGetTaskResponse = {
  encode(message: QueryGetTaskResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.Task !== undefined) {
      Task.encode(message.Task, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetTaskResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetTaskResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Task = Task.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTaskResponse {
    return { Task: isSet(object.Task) ? Task.fromJSON(object.Task) : undefined };
  },

  toJSON(message: QueryGetTaskResponse): unknown {
    const obj: any = {};
    message.Task !== undefined && (obj.Task = message.Task ? Task.toJSON(message.Task) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetTaskResponse>, I>>(object: I): QueryGetTaskResponse {
    const message = createBaseQueryGetTaskResponse();
    message.Task = (object.Task !== undefined && object.Task !== null) ? Task.fromPartial(object.Task) : undefined;
    return message;
  },
};

function createBaseQueryAllTaskRequest(): QueryAllTaskRequest {
  return { pagination: undefined };
}

export const QueryAllTaskRequest = {
  encode(message: QueryAllTaskRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTaskRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTaskRequest();
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

  fromJSON(object: any): QueryAllTaskRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllTaskRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllTaskRequest>, I>>(object: I): QueryAllTaskRequest {
    const message = createBaseQueryAllTaskRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllTaskResponse(): QueryAllTaskResponse {
  return { Task: [], pagination: undefined };
}

export const QueryAllTaskResponse = {
  encode(message: QueryAllTaskResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.Task) {
      Task.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllTaskResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllTaskResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Task.push(Task.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllTaskResponse {
    return {
      Task: Array.isArray(object?.Task) ? object.Task.map((e: any) => Task.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllTaskResponse): unknown {
    const obj: any = {};
    if (message.Task) {
      obj.Task = message.Task.map((e) => e ? Task.toJSON(e) : undefined);
    } else {
      obj.Task = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllTaskResponse>, I>>(object: I): QueryAllTaskResponse {
    const message = createBaseQueryAllTaskResponse();
    message.Task = object.Task?.map((e) => Task.fromPartial(e)) || [];
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
  /** Queries a Employer by index. */
  Employer(request: QueryGetEmployerRequest): Promise<QueryGetEmployerResponse>;
  /** Queries a list of Employer items. */
  EmployerAll(request: QueryAllEmployerRequest): Promise<QueryAllEmployerResponse>;
  /** Queries a Developer by index. */
  Developer(request: QueryGetDeveloperRequest): Promise<QueryGetDeveloperResponse>;
  /** Queries a list of Developer items. */
  DeveloperAll(request: QueryAllDeveloperRequest): Promise<QueryAllDeveloperResponse>;
  /** Queries a Task by id. */
  Task(request: QueryGetTaskRequest): Promise<QueryGetTaskResponse>;
  /** Queries a list of Task items. */
  TaskAll(request: QueryAllTaskRequest): Promise<QueryAllTaskResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Employer = this.Employer.bind(this);
    this.EmployerAll = this.EmployerAll.bind(this);
    this.Developer = this.Developer.bind(this);
    this.DeveloperAll = this.DeveloperAll.bind(this);
    this.Task = this.Task.bind(this);
    this.TaskAll = this.TaskAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Employer(request: QueryGetEmployerRequest): Promise<QueryGetEmployerResponse> {
    const data = QueryGetEmployerRequest.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Query", "Employer", data);
    return promise.then((data) => QueryGetEmployerResponse.decode(new _m0.Reader(data)));
  }

  EmployerAll(request: QueryAllEmployerRequest): Promise<QueryAllEmployerResponse> {
    const data = QueryAllEmployerRequest.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Query", "EmployerAll", data);
    return promise.then((data) => QueryAllEmployerResponse.decode(new _m0.Reader(data)));
  }

  Developer(request: QueryGetDeveloperRequest): Promise<QueryGetDeveloperResponse> {
    const data = QueryGetDeveloperRequest.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Query", "Developer", data);
    return promise.then((data) => QueryGetDeveloperResponse.decode(new _m0.Reader(data)));
  }

  DeveloperAll(request: QueryAllDeveloperRequest): Promise<QueryAllDeveloperResponse> {
    const data = QueryAllDeveloperRequest.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Query", "DeveloperAll", data);
    return promise.then((data) => QueryAllDeveloperResponse.decode(new _m0.Reader(data)));
  }

  Task(request: QueryGetTaskRequest): Promise<QueryGetTaskResponse> {
    const data = QueryGetTaskRequest.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Query", "Task", data);
    return promise.then((data) => QueryGetTaskResponse.decode(new _m0.Reader(data)));
  }

  TaskAll(request: QueryAllTaskRequest): Promise<QueryAllTaskResponse> {
    const data = QueryAllTaskRequest.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Query", "TaskAll", data);
    return promise.then((data) => QueryAllTaskResponse.decode(new _m0.Reader(data)));
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
