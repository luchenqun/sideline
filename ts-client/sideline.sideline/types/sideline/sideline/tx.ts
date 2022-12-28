/* eslint-disable */
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

/** Msg defines the Msg service. */
export interface Msg {
  RegistEmployer(request: MsgRegistEmployer): Promise<MsgRegistEmployerResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RegistDeveloper(request: MsgRegistDeveloper): Promise<MsgRegistDeveloperResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RegistEmployer = this.RegistEmployer.bind(this);
    this.RegistDeveloper = this.RegistDeveloper.bind(this);
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
