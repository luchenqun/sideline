/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "sideline.sideline";

export interface MsgRegistEmployer {
  creator: string;
  name: string;
  introduce: string;
  email: string;
  avatar: string;
}

export interface MsgRegistEmployerResponse {
}

export interface MsgRegistDeveloper {
  creator: string;
  name: string;
  introduce: string;
  email: string;
  avatar: string;
  education: string;
  major: string;
  skills: string[];
}

export interface MsgRegistDeveloperResponse {
}

export interface MsgCreateTask {
  creator: string;
  title: string;
  description: string;
  remuneration: string;
  deposit: string;
  collateral: string;
  deadline: number;
}

export interface MsgCreateTaskResponse {
  id: number;
}

export interface MsgDoTask {
  creator: string;
  id: number;
}

export interface MsgDoTaskResponse {
}

function createBaseMsgRegistEmployer(): MsgRegistEmployer {
  return { creator: "", name: "", introduce: "", email: "", avatar: "" };
}

export const MsgRegistEmployer = {
  encode(message: MsgRegistEmployer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.introduce !== "") {
      writer.uint32(26).string(message.introduce);
    }
    if (message.email !== "") {
      writer.uint32(34).string(message.email);
    }
    if (message.avatar !== "") {
      writer.uint32(42).string(message.avatar);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegistEmployer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegistEmployer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.introduce = reader.string();
          break;
        case 4:
          message.email = reader.string();
          break;
        case 5:
          message.avatar = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegistEmployer {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      introduce: isSet(object.introduce) ? String(object.introduce) : "",
      email: isSet(object.email) ? String(object.email) : "",
      avatar: isSet(object.avatar) ? String(object.avatar) : "",
    };
  },

  toJSON(message: MsgRegistEmployer): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.introduce !== undefined && (obj.introduce = message.introduce);
    message.email !== undefined && (obj.email = message.email);
    message.avatar !== undefined && (obj.avatar = message.avatar);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegistEmployer>, I>>(object: I): MsgRegistEmployer {
    const message = createBaseMsgRegistEmployer();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.introduce = object.introduce ?? "";
    message.email = object.email ?? "";
    message.avatar = object.avatar ?? "";
    return message;
  },
};

function createBaseMsgRegistEmployerResponse(): MsgRegistEmployerResponse {
  return {};
}

export const MsgRegistEmployerResponse = {
  encode(_: MsgRegistEmployerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegistEmployerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegistEmployerResponse();
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

  fromJSON(_: any): MsgRegistEmployerResponse {
    return {};
  },

  toJSON(_: MsgRegistEmployerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegistEmployerResponse>, I>>(_: I): MsgRegistEmployerResponse {
    const message = createBaseMsgRegistEmployerResponse();
    return message;
  },
};

function createBaseMsgRegistDeveloper(): MsgRegistDeveloper {
  return { creator: "", name: "", introduce: "", email: "", avatar: "", education: "", major: "", skills: [] };
}

export const MsgRegistDeveloper = {
  encode(message: MsgRegistDeveloper, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.introduce !== "") {
      writer.uint32(26).string(message.introduce);
    }
    if (message.email !== "") {
      writer.uint32(34).string(message.email);
    }
    if (message.avatar !== "") {
      writer.uint32(42).string(message.avatar);
    }
    if (message.education !== "") {
      writer.uint32(50).string(message.education);
    }
    if (message.major !== "") {
      writer.uint32(58).string(message.major);
    }
    for (const v of message.skills) {
      writer.uint32(66).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegistDeveloper {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegistDeveloper();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.introduce = reader.string();
          break;
        case 4:
          message.email = reader.string();
          break;
        case 5:
          message.avatar = reader.string();
          break;
        case 6:
          message.education = reader.string();
          break;
        case 7:
          message.major = reader.string();
          break;
        case 8:
          message.skills.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegistDeveloper {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      introduce: isSet(object.introduce) ? String(object.introduce) : "",
      email: isSet(object.email) ? String(object.email) : "",
      avatar: isSet(object.avatar) ? String(object.avatar) : "",
      education: isSet(object.education) ? String(object.education) : "",
      major: isSet(object.major) ? String(object.major) : "",
      skills: Array.isArray(object?.skills) ? object.skills.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: MsgRegistDeveloper): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.introduce !== undefined && (obj.introduce = message.introduce);
    message.email !== undefined && (obj.email = message.email);
    message.avatar !== undefined && (obj.avatar = message.avatar);
    message.education !== undefined && (obj.education = message.education);
    message.major !== undefined && (obj.major = message.major);
    if (message.skills) {
      obj.skills = message.skills.map((e) => e);
    } else {
      obj.skills = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegistDeveloper>, I>>(object: I): MsgRegistDeveloper {
    const message = createBaseMsgRegistDeveloper();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.introduce = object.introduce ?? "";
    message.email = object.email ?? "";
    message.avatar = object.avatar ?? "";
    message.education = object.education ?? "";
    message.major = object.major ?? "";
    message.skills = object.skills?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgRegistDeveloperResponse(): MsgRegistDeveloperResponse {
  return {};
}

export const MsgRegistDeveloperResponse = {
  encode(_: MsgRegistDeveloperResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegistDeveloperResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegistDeveloperResponse();
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

  fromJSON(_: any): MsgRegistDeveloperResponse {
    return {};
  },

  toJSON(_: MsgRegistDeveloperResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegistDeveloperResponse>, I>>(_: I): MsgRegistDeveloperResponse {
    const message = createBaseMsgRegistDeveloperResponse();
    return message;
  },
};

function createBaseMsgCreateTask(): MsgCreateTask {
  return { creator: "", title: "", description: "", remuneration: "", deposit: "", collateral: "", deadline: 0 };
}

export const MsgCreateTask = {
  encode(message: MsgCreateTask, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.remuneration !== "") {
      writer.uint32(34).string(message.remuneration);
    }
    if (message.deposit !== "") {
      writer.uint32(42).string(message.deposit);
    }
    if (message.collateral !== "") {
      writer.uint32(50).string(message.collateral);
    }
    if (message.deadline !== 0) {
      writer.uint32(56).uint64(message.deadline);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateTask {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateTask();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.remuneration = reader.string();
          break;
        case 5:
          message.deposit = reader.string();
          break;
        case 6:
          message.collateral = reader.string();
          break;
        case 7:
          message.deadline = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTask {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
      remuneration: isSet(object.remuneration) ? String(object.remuneration) : "",
      deposit: isSet(object.deposit) ? String(object.deposit) : "",
      collateral: isSet(object.collateral) ? String(object.collateral) : "",
      deadline: isSet(object.deadline) ? Number(object.deadline) : 0,
    };
  },

  toJSON(message: MsgCreateTask): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined && (obj.description = message.description);
    message.remuneration !== undefined && (obj.remuneration = message.remuneration);
    message.deposit !== undefined && (obj.deposit = message.deposit);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    message.deadline !== undefined && (obj.deadline = Math.round(message.deadline));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateTask>, I>>(object: I): MsgCreateTask {
    const message = createBaseMsgCreateTask();
    message.creator = object.creator ?? "";
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.remuneration = object.remuneration ?? "";
    message.deposit = object.deposit ?? "";
    message.collateral = object.collateral ?? "";
    message.deadline = object.deadline ?? 0;
    return message;
  },
};

function createBaseMsgCreateTaskResponse(): MsgCreateTaskResponse {
  return { id: 0 };
}

export const MsgCreateTaskResponse = {
  encode(message: MsgCreateTaskResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateTaskResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateTaskResponse();
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

  fromJSON(object: any): MsgCreateTaskResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreateTaskResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateTaskResponse>, I>>(object: I): MsgCreateTaskResponse {
    const message = createBaseMsgCreateTaskResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgDoTask(): MsgDoTask {
  return { creator: "", id: 0 };
}

export const MsgDoTask = {
  encode(message: MsgDoTask, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDoTask {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDoTask();
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

  fromJSON(object: any): MsgDoTask {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgDoTask): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDoTask>, I>>(object: I): MsgDoTask {
    const message = createBaseMsgDoTask();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgDoTaskResponse(): MsgDoTaskResponse {
  return {};
}

export const MsgDoTaskResponse = {
  encode(_: MsgDoTaskResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDoTaskResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDoTaskResponse();
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

  fromJSON(_: any): MsgDoTaskResponse {
    return {};
  },

  toJSON(_: MsgDoTaskResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDoTaskResponse>, I>>(_: I): MsgDoTaskResponse {
    const message = createBaseMsgDoTaskResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RegistEmployer(request: MsgRegistEmployer): Promise<MsgRegistEmployerResponse>;
  RegistDeveloper(request: MsgRegistDeveloper): Promise<MsgRegistDeveloperResponse>;
  CreateTask(request: MsgCreateTask): Promise<MsgCreateTaskResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DoTask(request: MsgDoTask): Promise<MsgDoTaskResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RegistEmployer = this.RegistEmployer.bind(this);
    this.RegistDeveloper = this.RegistDeveloper.bind(this);
    this.CreateTask = this.CreateTask.bind(this);
    this.DoTask = this.DoTask.bind(this);
  }
  RegistEmployer(request: MsgRegistEmployer): Promise<MsgRegistEmployerResponse> {
    const data = MsgRegistEmployer.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Msg", "RegistEmployer", data);
    return promise.then((data) => MsgRegistEmployerResponse.decode(new _m0.Reader(data)));
  }

  RegistDeveloper(request: MsgRegistDeveloper): Promise<MsgRegistDeveloperResponse> {
    const data = MsgRegistDeveloper.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Msg", "RegistDeveloper", data);
    return promise.then((data) => MsgRegistDeveloperResponse.decode(new _m0.Reader(data)));
  }

  CreateTask(request: MsgCreateTask): Promise<MsgCreateTaskResponse> {
    const data = MsgCreateTask.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Msg", "CreateTask", data);
    return promise.then((data) => MsgCreateTaskResponse.decode(new _m0.Reader(data)));
  }

  DoTask(request: MsgDoTask): Promise<MsgDoTaskResponse> {
    const data = MsgDoTask.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Msg", "DoTask", data);
    return promise.then((data) => MsgDoTaskResponse.decode(new _m0.Reader(data)));
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
