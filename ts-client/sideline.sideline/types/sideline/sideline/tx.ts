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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RegistEmployer(request: MsgRegistEmployer): Promise<MsgRegistEmployerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RegistEmployer = this.RegistEmployer.bind(this);
  }
  RegistEmployer(request: MsgRegistEmployer): Promise<MsgRegistEmployerResponse> {
    const data = MsgRegistEmployer.encode(request).finish();
    const promise = this.rpc.request("sideline.sideline.Msg", "RegistEmployer", data);
    return promise.then((data) => MsgRegistEmployerResponse.decode(new _m0.Reader(data)));
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
